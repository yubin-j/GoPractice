package accounts

import "errors"

/*
Account struct
Account struct의 경우 지금은 Account로 하여 이 파일을 임포트 하는 곳에서 접근하도록 해둔다.
Account struct의 owner와 balance는 현재 소문자로 시작하여 private한 변수이고
만일 이를 Owner, Balance로 작성 시 외부에서도 접근이 가능한 public 변수가 된다.
즉 이 public과 private한 성질을 결정짓는 요인은
변수명이나 함수명 들이 대문자로 시작하면 public한 것이고 소문자로 시작하면 private한 것이다.
*/
type Account struct {
	owner   string
	balance int
}

//err관련 변수의 경우 변수 이름시작이 err로 되야 한다고 golint에서 알려준다.
//변수명으로 어떤 종류의 변수인지 한번에 알 수 있으므로 보기 좋다.
var errNoMoney = errors.New("Can't withdraw you are poor")

/*
NewAccount create Account
Account struct의 owner와 balance가 private변수 이므로 seter와 같은 함수를 작성해줘야 한다.
여기서는 Account struct의 참조값을 출력함으로서 생성자(constructor)라고 볼 수 있다.

반환해주는 account struct의 주소값을 *Account형식 즉 account struct의 값을 출력해준다.
메모리에 복사가 일어나지 않으므로 효율적으로 동작한다고 볼 수 있다.
*/
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

/*
Deposit x amount on your account
Go에서 메서드는 func와는 다른 것이다.
형식은 func와 비슷하지만 func와 함수 이름 사이에 리시버라는 것이 들어간다.
아래 메서드에서 리시버는 (a Account)이다.
이는 이 함수를 Account의 메서드로 사용하겠다는 것이다. 리시버를 넣게 되면 메서드로 동작하게 된다고 생각하면 된다.
리시버의 구성은 다음과 같다 => (struct의 정보를 담을 이름 struct)
아래 메서드에서 struct의 정보를 담을 이름은 a이고 stuct는 Account이다.
참고로 이름은 struct의 첫글자를 소문자로 한 것으로 시작해야 한다.
Account의 첫글자의 소문자는 a이므로 아래와 같이 작성한 것이다.
메서드 안에서는 a가 Account의 정보를 담고 있으므로
a.balance와 같은 구문으로 값에 접근 할 수 있다.
마치 자바스크립트의 this처럼 말이다.

자 추가적으로 리시버를 (a Account)로 해서 값을 넣어봤더니 값이 들어가지 않았다.
왜냐하면 포인터 때문인데 위에 생성자를 보면 반환하는 값이 account의 주소값이다.
즉 우리가 이를 받을때도 주소값을 그대로 받는 것이 아닌 그 값이 바라보고 있는 struct를 가져와야 한다.
그래서 아래에 적혀있는 것 처럼 리시버를 (a *Account)로 하여 값을 가져와야 한다.

이러한 리시버를 가리켜 포인터 리시버라고 한다.
*/
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

/*
Balance show balance
아래에서 사용한 리시버는 포인터 리시버가 아닌 일반 리시버인데
main.go에서 같은 방식으로 메서드를 호출하는데 왜 일반 리시버도 동작을 할까?
Go는 기본적으로 값을 복사하여 전달하는데 아래와 같이 타입이 포인터 타입이 아닌 경우에 값을 복사하여 전달하고
포인터 타입인 경우에만 참조값을 전달한다.
아래의 경우 값을 조작할 필요 없이 불러오기만 하면 되기에 일반 타입을 사용한 리시버를 이용한 것이다.
*/
func (a Account) Balance() int {
	return a.balance
}

/*
Withdraw from your account
Go에서는 Exception관련 기능(try - except, try - catch)이 없어서 직접 관리를 해줘야 한다.
*/
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		//return errors.New("Can't withdraw you are poor")
		//위와 같이 errors.New함수를 사용하여 에러를 출력해준다.
		//추가적으로 리턴하는 곳에 새로운 에러를 만들어서 사용해도 되지만
		//전역에 error를 선언해두고 값을 불러오는 것이 더 깔끔하다고 한다.
		return errNoMoney
	}
	a.balance -= amount
	return nil //nil은 null과 같은 것이라고 보면 된다.
	//error가 반환타입으로 되어 있기에 정상적인 경우 nil값을 출력함으로서 함수의 반환타입을 만족시킨다.
}
