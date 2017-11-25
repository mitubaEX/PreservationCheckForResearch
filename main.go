package main

import (
	"./settings"
	"fmt"
	"./search"
)

func main() {
	var argument settings.Argument = settings.SearchInit()
	if argument.Mode == "search" {
		search.Search(argument)
	}
	fmt.Println(argument.Filename)
}
