package main

import (
	"fmt"

	"strings"

	"github.com/imroc/req"
)


func getFromProxyScrape() []string {
	resp, err := req.Get("https://api.proxyscrape.com/?request=getproxies&proxytype=http&timeout=10000&country=all&ssl=all&anonymity=all")

	if err != nil {
		fmt.Println(err)
	}

	proxiesString, err := resp.ToString()

	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(proxiesString, "\n") 
}