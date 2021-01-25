package main

import (
	"fmt"

	"github.com/yubin-j/GoPractice/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	err := dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
	definition, _ := dictionary.Search("hello")
	fmt.Println("found: hello, definition:", definition)
	err = dictionary.Add("hello", "Greeting")
	if err != nil {
		fmt.Println(err)
	}
}
