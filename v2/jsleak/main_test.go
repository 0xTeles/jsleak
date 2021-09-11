package main

import (
	"testing"
)

func BenchmarkReq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Random URL
		req("http://testphp.vulnweb.com/", 5)
	}
}
