package main

import "fmt"

/*
structure에 기반이 되는 struct
javascript의 객체와 비슷한 기능을 한다.
각 요소별로 타입을 각각 명시하여 다양한 데이터 타입의 변수들을 담을 수 있다.
struct는 javascript의 constructor와 같은 함수가 없어서 원하면 직접 생성하여 사용해야 한다.
*/
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchijjigae", "gukbob"}

	//사용할때는 아래와 같이 각 데이터 타입에 맞는 데이터들을 전달해준다.
	yubin := person{"yubin", 28, favFood}
	fmt.Println(yubin)
	fmt.Println(yubin.name)
	fmt.Println(yubin.age)
	fmt.Println(yubin.favFood)

	//키값을 연결하여 즉각적으로 값들이 어떤 키값에 들어가는지 알 수 있게 사용할 수도 있다.
	yubin1 := person{name: "yubin", age: 28, favFood: favFood}
	fmt.Println(yubin1)

	//단 넘겨주는 형식이 인자들끼리 동일해야 한다.
	//아래 코드처럼 키 값을 연결한 것과 안한 것을 혼합해서 사용할 수가 없다.
	//그렇지 않으면 에러가 발생한다.
	//yubin2 := person{name: "yubin", 28, favFood}
}
