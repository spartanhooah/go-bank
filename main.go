package main

import (
	"bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const balanceFilename = "balance.txt"

func main() {
	accountBalance, err := fileops.ReadFloatFromFile(balanceFilename)

	if err != nil {
		fmt.Println(err)
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us at ", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayBalance(accountBalance)
		case 2:
			accountBalance = depositOrWithdraw(accountBalance, false)
		case 3:
			accountBalance = depositOrWithdraw(accountBalance, true)
		case 4:
			fmt.Println("Have a good day!")
			return
		default:
			fmt.Printf("Invalid choice.\n\n")
			continue
		}

		fmt.Println()
	}
}

func depositOrWithdraw(accountBalance float64, withdraw bool) float64 {
	choice := "deposit"

	if withdraw {
		choice = "withdraw"
	}

	fmt.Print("How much do you want to ", choice, "? ")

	var amount float64
	fmt.Scan(&amount)

	if amount <= 0 || (withdraw && amount > accountBalance) {
		fmt.Println("Invalid amount entered.")
		return accountBalance
	}

	if withdraw {
		accountBalance -= amount
	} else {
		accountBalance += amount
	}

	displayBalance(accountBalance)
	fileops.WriteFloatToFile(accountBalance, balanceFilename)

	return accountBalance
}

func displayBalance(accountBalance float64) {
	fmt.Printf("Your balance is %.2f\n", accountBalance)
}
