package main

import (
	"github.com/TwinProduction/go-color"

	fig "github.com/mbndr/figlet4go"

	prompt "github.com/segmentio/go-prompt"

	"fmt"
)

const figletText = "Proxy-Dl"


func generateFiglet() {
	ascii := fig.NewAsciiRender()

	//load font
	ascii.LoadFont("fonts")

	options := fig.NewRenderOptions()

	options.FontName = "speed"

	renderStr, _ := ascii.RenderOpts(figletText, options)
	fmt.Print(color.Ize(color.Red, renderStr))
}


var actionOptions = []string{
	color.Ize(color.Yellow, "Download"),
	color.Ize(color.Yellow, "Check"),
}

func getAction() int {
	return prompt.Choose(color.Ize(color.Cyan, "Please select an action"), actionOptions)
}

func getFilePath() string {
	return prompt.StringRequired(color.Ize(color.Cyan, "Please enter the name of the proxy file you wish to check"))
}
