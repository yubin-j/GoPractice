package accounts

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

/*
NewAccount create Account
Account struct의 owner와 balance가 private변수 이므로 seter와 같은 함수를 작성해줘야 한다.
여기서는 Account struct의 참조값을 출력함으로서 생성자라고 볼 수 있다.

반환해주는 account struct의 주소값을 *Account형식 즉 account struct의 값을 출력해준다.
메모리에 복사가 일어나지 않으므로 효율적으로 동작한다고 볼 수 있다.
*/
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
