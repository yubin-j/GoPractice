package main

import "fmt"

global := "I am global" //func밖에서는 무조건 모든 선언식을 작성해줘야 한다. 아래에 내용이 나온다.

func main() {
	fmt.Println("Hello Word")

	/*
	   상수, 변수 선언 법
	   const와 var로 상수, 변수 구분을 해주고
	   타입을 이름뒤에 지정해주어야 한다.
	*/
	const name string = "Yubin"
	var temp string = "AAA"

	name = "test" //상수이기에 변경 불가
	temp = "test" //변수이기에 변경 가능

	//변수 선언 시 다음과 같이도 선언 할 수 있으며 이는 func안에서만 가능하다.
	//func 밖에서는 불가능한 선언방식이며 타입은 처음에 대입한 값의 타입으로 선언되며 타입 변경이 불가능하다.
	//상수는 이와 같이 선언하는 것이 불가능하다. 오직 변수용이다.
	temp2 := true

	temp2 = "str" //Bollean 타입이 아닌 값을 대입하므로 에러가 난다.
}
