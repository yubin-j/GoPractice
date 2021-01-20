package main

import (
	"fmt"
	"log"

	"github.com/yubin-j/GoPractice/accounts"
)

func main() {
	/*
	   struct에 대해서 연습.
	   accounts폴더에 accounts.go파일을 생성하여 accounts package를 구성.

	*/
	account := accounts.NewAccount("yubin")
	//account.balance = 1000  balance가 undefined라고 에러가 난다.
	//이는 balance가 private한 값이기 때문이며 이를 수정하기 위해서는 메서드를 만들어 사용해야한다.

	//메서드는 아래와 같이 .으로 전개하여 부르면 된다.
	account.Deposit(1000)

	//아래의 경우 에러가 출력되지만 Go에서는 이를 가지고 아무런 동작을 취하지 않는다.
	//그래서 다음 코드들 처럼 직접 에러를 처리하는 코드르 작성해야 한다.
	account.Withdraw(1200)

	err := account.Withdraw(1200)
	if err != nil {
		//log.Fatalln은 err를 출력해주고 프로그램 실행을 즉시 종료시키는 함수이다.
		log.Fatalln(err)
	}

	fmt.Println(account.Balance()) //받아온 값이 참조값을 바라보는 값을 받아오게 되어 출력형식에 이를 알려주는 &가 붙는다.
}
