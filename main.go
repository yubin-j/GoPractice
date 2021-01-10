package main

import (
	"fmt"

	"github.com/yubin-j/GoPractice/something"
)

func main() {
	fmt.Println("Hello Word")

	something.SayHello() //대문자로 시작되는 something의 함수는 정상적으로 불러진다.
	something.sayBye()   //반면 소문자로 시작되는 함수는 불러지지 않는다.
}
