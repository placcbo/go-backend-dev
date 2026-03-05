package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (d *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit money cant be zero or less")
	}
	d.Balance += amount

}

func main() {
	kevin := BankAccount{
		Owner:   "Kevin",
		Balance: 250.0,
	}

	kevin.Deposit(600)
	fmt.Println(kevin)

}
