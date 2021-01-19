package main

import (
	"fmt"

	"github.com/yubin-j/GoPractice/accounts"
)

func main() {
	/*
	   struct에 대해서 연습.
	   accounts폴더에 accounts.go파일을 생성하여 accounts package를 구성.

	*/
	account := accounts.NewAccount("yubin")
	fmt.Println(account)
}
