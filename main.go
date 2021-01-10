package main

import (
	"fmt"
	"strings"
)

/*
인자들의 변수명 뒤에도 타입을 정해주어야 하고 함수의 return 값이 있는 경우 역시 함수명 뒤에 타입을 정해줘야 한다.
입력 인자가 동일한 타입이면 다음과 같이 작성할 수도 있다.
a int, b int === a, b
*/
func multiply(a int, b int) int {
	return a * b
}

/*
Go언어의 특별한점.
여러개의 값을 함수가 리턴할 수 있다.
이런 기능은 완전 처음이다. 보통 배열같은 것을 넘겨서 하는 경우는 있어도
함수 자체가 여러개의 값을 출력해주는건 아마 Go가 유일한 듯 하다.
*/
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

/*
인자로 몇개의 값이 들어오는지 모를때는 타입앞에 ...을 붙여준다.
자바스크립트의 spread구문과 흡사하다.
*/
func repeatMe(words ...string) {
	fmt.Println(words)
}

/*
naked return
리턴하는 타입에 변수명을 붙여서 해당 값을 변경 시 자동으로 그 값으로 출력이 된다.
이는 함수 호출 시 리턴되는 값에도 미리 변수를 선언해둔다고 보면 된다.
return구문은 작성을 해줘야 한다.
이런 기능또한 처음 본다.
*/
func lenAndUpperNakedReturn(name string) (length int, upperCase string) {
	/*
		defer
		이 기능은 무엇인가... 바로 함수가 종료된 후 실행하고 싶은 코드를 실행할 수 있는 기능이다.
		이 코드들을 실행해보면 defer에 있는 코드가 먼저 실행되고 나머지 연산이 진행되는 것을 볼 수 있다.
		이는 사진을 편집하는 함수였다면 편집이 완료된후 사진 파일을 닫는 식으로 사용할 수 있다.
		난 아직 이 기능이 얼마나 좋은지를 잘 모르겠으나..
		앞으로 배워나가면서 얼마나 좋은지 경험해 보자
	*/
	defer fmt.Println("I'm done")
	length = len(name)
	upperCase = strings.ToUpper(name)
	//아래 코드는 실험삼아 넣어본 것인데 defer의 경우 거꾸로 올라가면서 실행되고 함수 안의 변수도 접근이 가능하다.
	defer fmt.Println(length)
	defer fmt.Println("I'm done2")
	return
}

func main() {
	fmt.Println(multiply(2, 4))

	totalLen, upperName := lenAndUpper("Yubin")
	fmt.Println(totalLen, upperName)

	//만약 여러개의 값을 출력해주는 함수에서 특정한 값만 받고 싶으면
	//다음과 같이 _구문을 사용하여 해당 값을 무시해 줄 수 있다.
	//_는 변수명이 아닌 값을 무시하는 기능을 하는 구문이다.
	totalLen2, _ := lenAndUpper("test")
	fmt.Println(totalLen2)

	repeatMe("a", "b", "c", "d", "e")

	fmt.Println(lenAndUpperNakedReturn("test11"))
}
