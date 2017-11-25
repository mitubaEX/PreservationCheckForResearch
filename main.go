package main

import (
	"./compare"
	"./search"
	"./settings"
	"fmt"
)

func main() {
	var argument settings.Argument = settings.SearchInit()
	for _, v := range argument.Modes {
		if v == "search" {
			search.Search(argument)
		} else if v == "compare" {
			compare.Compare(argument)
		}
	}
	fmt.Println(argument.Filename)
}
