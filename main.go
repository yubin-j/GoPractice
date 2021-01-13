package main

import "fmt"

func main() {
	/*
		Go에서의 배열
		배열을 만드는 법
		[]안에 배열의 길이를 지정해주고 다음으로는
		배열안의 요소가 어떤 타입인지 지정해준다.
		그리고 값을 {}를 이용하여 넣어준다.
	*/
	names := [5]string{"yubin", "gildong"}
	fmt.Println(names)

	//특정 인덱스에 값을 넣을 시에는 아래처럼 넣으면 되며, 인덱스가 배열의 길이를 벗어나면 에러가 발생한다.
	names[2] = "test"
	fmt.Println(names)

	/*
		사실... 무조건 배열을 선언할 때 길이를 지정해야 하는 것은 아니다.
		배열의 길이를 지정하는 []안에 아무런 값도 안 넣으면 된다.
		이러한 데이터 타입은 배열이라 하지 않고 slice라고 한다.
		slice는 기본적으로 배열이긴 하나 길이가 지정되지 않은 데이터 형식이다.
	*/
	citys := []string{"Seoul", "Busan"}
	fmt.Println(citys)

	/*
		slice를 선언하고 나서 값을 더 추가하고 싶으면 append라는 함수를 사용해야 하는데
		append함수는 인자로 slice와 추가하고 싶은 값을 받아들여서 값을 추가한 slice를 반환해준다.
		그래서 반환값을 받는 변수도 필요하다 아래 코드를 보자.
	*/
	citys = append(citys, "Daegu")
	fmt.Println(citys)
}
