package main

import (
	"crypto/tls"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func BenchmarkReq_Old(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reqhttp_old("http://testphp.vulnweb.com/", 5)
	}
}

func BenchmarkReq_FastHttp(b *testing.B) {
	c := &fasthttp.Client{
		Name: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36",
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxConnWaitTimeout: time.Duration(5) * time.Second,
	}

	for i := 0; i < b.N; i++ {
		reqhttp("http://testphp.vulnweb.com/", c)
	}
}
