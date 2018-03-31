package main

import (
	"testing"
	"MyGo/code/network/dns/file"
)

func TestAdd(t * testing.T) {
	fd := file.Read("subnames_full.txt");
	CheckDNS("google.com",fd, 0, 500, 0);
}
func BenchmarkCheckDNS(b *testing.B){
	fd := file.Read("subnames_full.txt");
	for i := 0;i < b.N; i++{	
		go CheckDNS("google.com",fd, 200*i, 200*i+200, 200*i);
	}
}
