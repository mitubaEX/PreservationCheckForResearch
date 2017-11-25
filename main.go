package main

import (
	"./settings"
	"fmt"
	"./search"
)

func main() {
	var argument settings.Argument = settings.SearchInit()
	for _, v := range argument.Modes{
		if v == "search" {
			search.Search(argument)
		}
	}
	fmt.Println(argument.Filename)
}
