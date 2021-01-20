package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw you are poor")

//NewAccount create account (constructor)
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit on account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

//Balance show from account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw from account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

/*
Account struct를 사용하는 경우
println과 같은 곳에서 Account struct를 가리키는 변수 자체를 출력하는 형식을 지정할 수 있다.

자바스크립트에서를 예로 들면 함수 선언 후 해당 주소값을 가지고 있는 변수를 콘솔로그에 출력 시
함수 전체 코드가 문자열로 나오는 것을 볼 수 있다.
이는 함수 메서드의 toString()을 통해서 나오는 결과인데 이 toString 메서드를 다른 함수로 변경 시
변경된 함수에서 출력해주는 대로 나오는 것을 볼 수 있다.

Go에서도 위의 예와 같은 것을 할 수 있으며 아래 코드들처럼 작성을 하면 된다.
메서드명은 String으로 해야하며 리시버로는 포인터가 아닌 값을 복사해서 오게 하면 된다.
*/
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
