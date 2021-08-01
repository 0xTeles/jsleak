package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

func getLeak(url string, data string, pattern string){
	re := regexp.MustCompile(pattern)
	all := re.FindAllStringSubmatch(data, -1)
	for _, element := range all {
		fmt.Println("[+] Match:", pattern, "-", element[0], "-",url)
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
	url := os.Args[1]
	path := os.Args[2]
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