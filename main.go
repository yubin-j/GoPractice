package main

import "fmt"

func main() {
	/*
		Go의 map
		javascript의 map과 같은 기능은 아니다.
		오히려 object의 자료형과 흡사한 자료형이다.
		map은 키값과 값이 있으며 키 값과 값의 자료형 또한 명시적으로 지정하여 사용해야 한다.
		그래서 javascript의 object와는 비슷하지만 한정적인 자료형이다.
		만드는 방법은 아래 코드와 같이 작성되며 의미는 다음줄의 주석과 같다.
		map[키값의 자료형]값의 자료형{키:값, 키:값}
	*/
	yubin := map[string]string{"name": "yubin", "age": "28"}

	//map은 range로 전달하여 사용할 수 있다.
	for key, value := range yubin {
		fmt.Println(key, value)
	}

	//map에 있는 값을 가져오거나 추가하는 방법과 관련 함수들은 추후에 연습을 할 것이다.
}
