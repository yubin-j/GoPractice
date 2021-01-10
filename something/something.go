package something

import "fmt"

//SayHello 외부로 Export하는 함수의 경우 함수명 시작이 대문자로 시작되어야 한다.
func SayHello() {
	fmt.Println("Hello")
}

func sayBye() {
	fmt.Println("Bye")
}
