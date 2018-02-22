package main

import (
	"fmt"
	"net"
)

func main() {

	// Resolve by IP
	addrs, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}

	//Resolve by address
	ips, err := net.LookupIP("localhost")
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
