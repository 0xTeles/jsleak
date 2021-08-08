package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"flag"
	"github.com/gijsbers/go-pcre"
)

func getLeak(url string, data string, pattern string){
	re := pcre.MustCompile(pattern,0)
	matches := re.MatcherString(data,0).Group(0)
	fmt.Println(len(matches))
	if (len(matches) != 0){
		fmt.Printf("[+] Url: %v\n[+] Pattern: %v\n[+] Match: %v\n", url,pattern,string(matches))
	}
}

func req(url string) string {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Transport: transCfg}
	res, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return string(data)
}

func main(){
	url := flag.String("url", "", "JS endpoint to test")
	path := flag.String("pattern", "", "File contains patterns to test")
	flag.Parse()
	if *url == ""{
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(*path)
	if err != nil{
		log.Fatal(err)
	}
	data := req(*url)
	defer file.Close()
	pattern := bufio.NewScanner(file)
	for pattern.Scan(){
		getLeak(*url,data,pattern.Text())
	}
}
