package main

import "fmt"

func canIDrink(age int) bool {
	/*
		variable expression
		if문안에서만 사용할 수 있는 변수를 생성할 수 있으며
		아래와 같이 변수를 생성 후 ;으로 구분을 해주면 된다.
		자바스크립트 for문의 초기화 구문이라고 보면 된다.
	*/
	if koreanAge := age + 1; koreanAge < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
}
