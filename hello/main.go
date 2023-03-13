package main

import "fmt"

type BackAccount struct {
	owner   string
	balance int
}

func (b *BackAccount) Deposit(amount int) {
	b.balance += amount
}

func (b *BackAccount) Withdraw(amount int) {
	b.balance -= amount
}

func (b *BackAccount) ChangeName(name string) {
	b.owner = name
}

func main() {
	account := &BackAccount{"Shuto", 1000}
	account.Deposit(100)
	// account.balance = 200
	fmt.Println(account.balance)
	account.Withdraw(200)
	fmt.Println(account.balance)
	account.ChangeName("Shuto2")
	fmt.Println(account.owner)
}
