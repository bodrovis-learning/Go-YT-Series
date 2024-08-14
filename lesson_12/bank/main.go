package main

import (
	"bank/bank"
	"fmt"
	"time"
)

func main() {
	go func() { // A
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
	}()

	go bank.Deposit(100) // B

	time.Sleep(time.Second * 1)

	fmt.Println("=", bank.Balance())

	bank.Withdraw(100)

	fmt.Println("=", bank.Balance())
}
