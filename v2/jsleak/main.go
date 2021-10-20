package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"strings"

	"github.com/gijsbers/go-pcre"
)

type JsonReturn struct {
	Url     string
	Pattern string
	Match   string
}

func getLeak(url string, data string, c_patterns []pcre.Regexp, raw_patterns []string, jsonArray *[]JsonReturn) {
	
	for i,pattern := range c_patterns {

		// Match function
		regex_ := func(i int, pattern pcre.Regexp, data_i *string) (bool){
			data_b := *data_i
			matches := pattern.MatcherString(data_b, 0)
			matches_index := matches.Index()
			matches_str := matches.GroupString(0)

			if len(matches_str) != 0 {
				fmt.Printf("[+] Url: %v\n[+] Pattern: %v\n[+] Match: %v\n", url, raw_patterns[i], matches_str)
				
				//JSON Output
				jsn := JsonReturn{url, raw_patterns[i], matches_str}
				*jsonArray = append(*jsonArray, jsn)
				
				//Remove match value from data
				data_b = (data_b[:matches_index[0]]+data_b[matches_index[1]:]) 
				*data_i = data_b
				return true

			}else {
				return false
			}
		}

		// Loop same pattern until find no more 
		for regex_(i,pattern, &data) { fmt.Printf("") }
		
	}
}

func get_inputs() []string {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return strings.Fields(string(output))
}

func req(url string, timeout int) string {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{
		Transport: transCfg,
		Timeout: time.Duration(timeout) * time.Second,
	}
	res, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(data)
}

func main() {
	path := flag.String("pattern", "", "[+] File contains patterns to test")
	verbose := flag.Bool("verbose", false, "[+] Verbose Mode")
	jsonOutput := flag.String("json", "", "[+] Json output file")
	timeout := flag.Int("timeout", 5, "[+] Timeout for request in seconds")
	flag.Parse()

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("[+] Use in Pipeline")
		os.Exit(1)
	}

	file, err := os.Open(*path)
	defer file.Close()
	lines := make([]string, 0)

	patterns := bufio.NewScanner(file)
	jsonArray := make([]JsonReturn, 1)

	for patterns.Scan() {
		lines = append(lines, patterns.Text())
	}

	// Compile all patterns
	c_patterns := []pcre.Regexp{}
	for _, pattern := range lines {
		c_patterns = append(c_patterns, pcre.MustCompile(pattern, 0))
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, url := range get_inputs() {
		if *verbose {
			fmt.Println("[-] Looking: " + url)
		}
		data := req(url,*timeout)
		getLeak(url, data, c_patterns, lines, &jsonArray)

	}

	if *jsonOutput != "" {
		fo, err2 := os.Create(*jsonOutput)
		k, err1 := json.MarshalIndent(jsonArray, "", "\t")
		if _, err := fo.Write(k); err1 != nil || err2 != nil {
			panic(err)
		}
	}
}
