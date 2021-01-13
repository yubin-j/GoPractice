package main

import "fmt"

func canIDrink(age int) bool {
	/*
		if에서 작동하는 variable expression이 switch에서도 작동한다.
		switch에 인자를 주지 않고 case에서 직접 비교식을 사용하여 조건 검사를 수행 할 수도 있다.
	*/
	// switch koreanAge := age + 1; koreanAge {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// case 20:
	// 	return true
	// }
	switch koreanAge := age + 1; {
	case koreanAge < 18:
		return false
	case koreanAge > 50:
		return false
	case koreanAge >= 18:
		return true
	}

	return false
}

func main() {
	fmt.Println(canIDrink(18))
}
