package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		os.Stderr.WriteString("usage: go-dns-test [hostname]")
		os.Exit(1)
	}
	hostname := os.Args[1]
	ips, err := net.LookupIP(hostname)
	if err != nil {
		panic(err)
	}
	if len(ips) == 0 {
		os.Stderr.WriteString("error: no A or AAAA records resolved")
		os.Exit(1)
	}
	fmt.Println(ips)
}
