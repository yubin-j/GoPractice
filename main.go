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
	fmt.Println(account) //받아온 값이 참조값을 바라보는 값을 받아오게 되어 출력형식에 이를 알려주는 &가 붙는다.
}
