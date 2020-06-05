package log

import (
	"github.com/TwinProduction/go-color"

	

	"fmt"
)

type logElem struct {
	content string
	kind    string
}

var logged []logElem

func Error(err string) {
	output := color.Ize(color.Red, fmt.Sprintf("Error : %s", err))
	fmt.Println(output)
	elem := logElem{
		output,
		"error",
	}
	logged = append(logged, elem)
}

func Warn(warn string) {
	output := color.Ize(color.Yellow, fmt.Sprintf("Warn : %s", warn))
	fmt.Println(output)
	elem := logElem{
		output,
		"warn",
	}
	logged = append(logged, elem)
}

func Info(info string) {
	output := color.Ize(color.Yellow, fmt.Sprintf("Info : %s", info))
	fmt.Println(output)
	elem := logElem{
		output,
		"info",
	}
	logged = append(logged, elem)
}

func DumpToFile(){
	for _, elem := range logged {
		fmt.Println(elem.content)
	}
}
