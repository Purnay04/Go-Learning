package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	holderName string
	balance    float64
}

func (ba *BankAccount) depositeMoney(user string, deposit float64) {
	ba.balance += deposit
	fmt.Printf("%s added %.2f money into bank account", user, deposit)
	fmt.Println()
}

func main() {
	var wg sync.WaitGroup

	bankAccount := &BankAccount{
		holderName: "Purnay Barge",
		balance:    0.00,
	}

	numOfUser := 5
	bankBal := bankAccount.balance
	for i := range numOfUser {
		wg.Add(1)
		// depositeMoney := float64(rand.Intn(500))
		depositeMoney := 100.00
		go func() {
			defer wg.Done()
			bankAccount.depositeMoney(fmt.Sprintf("user%d", i), depositeMoney)
		}()
		bankBal += depositeMoney
	}

	wg.Wait()

	fmt.Println()
	fmt.Println()
	fmt.Println("Total Bank Account Balance Should be:", bankBal)
	fmt.Println("Actual Balance is:", bankAccount.balance)

}
