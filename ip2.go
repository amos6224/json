package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupIP("https://ipinfo.io/json")
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		fmt.Println("IP address: ", addr)
	}
}
