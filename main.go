package main

import (
	"fmt"
	"os"
	"strings"

	"io/ioutil"

	"./log"

	"github.com/schollz/progressbar"

	"./checkp"
)

func main() {
	generateFiglet()
	fmt.Println("Written by [REDACTED]")

	if getAction() == 0 {
		downloadProxies()
	} else {
		path := getFilePath()
		checkProxiesLoop(path)
	}

}

func checkProxiesLoop(fileName string) {
	var validProxies []string

	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Error("File not found. Exiting")
		os.Exit(0)
	}

	contentString := string(contents)

	proxyList := strings.Split(contentString, "\n")

	bar := progressbar.Default(int64(len(proxyList)), "Checking...")

	for _, proxy := range proxyList {
		if checkp.CheckProxy(proxy) {
			validProxies = append(validProxies, proxy)
			bar.Add(1)
		}
	}

	log.Info("Dumping valid proxies to validated_proxies.txt")

	dumpSliceToFile(validProxies, "validated_proxies.txt")

}

func downloadProxies() {
	var proxies []string

	proxies = append(proxies, getFromProxyScrape()...)

	log.Info("Dumping proxies to output.txt")
	dumpSliceToFile(proxies, "output.txt")
}
