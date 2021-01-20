package main

import (
	"fmt"

	"github.com/yubin-j/GoPractice/accounts"
)

func main() {
	account := accounts.NewAccount("yubin")

	account.Deposit(1000)

	fmt.Println(account)
}
