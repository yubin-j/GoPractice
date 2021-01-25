package main

import (
	"fmt"

	"github.com/yubin-j/GoPractice/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("second")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
