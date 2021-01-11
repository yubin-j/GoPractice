package main

import "fmt"

func superAdd(numbers ...int) int {
	/*
		Go 에서 반복문은 오직 for만 존재한다.
		foreach나 map, for of, for in 은 존재하지 않는다.
		단 range를 사용하여 for of와 비슷하게 동작한다.
		range는 for 에서만 동작한다.
	*/
	for number := range numbers {
		fmt.Println(number) // 0, 1, 2, 3, 4 ,5
		//위와 같이 작성하면 배열의 인덱스를 가져오게 된다. 값을 가져오려면 아래와 같이 하면 된다.
	}

	fmt.Println("----------------------------")

	for index, value := range numbers {
		fmt.Println(index, value)
		/*.
		0 1
		1 2
		2 3
		3 4
		4 5
		5 6
		*/
		//위와 같이 작성하면 배열의 인덱스를 가져오게 된다. 값을 가져오려면 아래와 같이 하면 된다.
	}

	fmt.Println("----------------------------")

	/*
		superAdd 함수의 기능에 맞게 for문을 사용해보자.
		range의 index값은 필요없기에 _로 무시를 해준다.
	*/

	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	fmt.Println(superAdd(1, 2, 3, 4, 5, 6))
}
