package main

import (
	"strings"

	"io/ioutil"
)

func dumpSliceToFile(s []string, fileName string) {
	chunkString := strings.Join(s, "\n")
	data := []byte(chunkString)
	ioutil.WriteFile(fileName, data, 0777)
}
