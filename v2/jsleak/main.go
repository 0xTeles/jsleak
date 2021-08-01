package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/gijsbers/go-pcre"
)

func getLeak(url string, data string, pattern string){
	re := pcre.MustCompile(pattern,0)
	matches := re.MatcherString(data,0).Matches()
	if matches{
		fmt.Println("[+] Match:", matches, "-",pattern,"-", url)
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

	var url, path string 

	if len(os.Args) > 1 {

		url = os.Args[1]
		path = os.Args[2]

	} else {

		fmt.Println("[+] It is necessary to set the URL and the file with the regexs.")
	    os.Exit(0)

	}

	file, err := os.Open(path)
	if err != nil{
		log.Fatal(err)
	}
	data := req(url)
	defer file.Close()
	pattern := bufio.NewScanner(file)
	for pattern.Scan(){
		getLeak(url,data,pattern.Text())
	}
}
