package main

import (
	"fmt"

	"github.com/yubin-j/GoPractice/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "First")

	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

	err = dictionary.Update("baseWord", "Second")
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	word, err = dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
