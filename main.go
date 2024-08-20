package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFilename = "balance.txt"

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFilename)

	if err != nil {
		return 0, errors.New("failed to find or read balance file")
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, errors.New("failed to parse balance file")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	os.WriteFile(balanceFilename, []byte(fmt.Sprintf("%.2f", balance)), 0644)
}

func main() {
	accountBalance, err := readBalanceFromFile()

	if err != nil {
		fmt.Println(err)
		//panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayBalance(accountBalance)
		case 2:
			depositOrWithdraw(accountBalance, false)
		case 3:
			depositOrWithdraw(accountBalance, true)
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

func depositOrWithdraw(accountBalance float64, withdraw bool) {
	choice := "deposit"

	if withdraw {
		choice = "withdraw"
	}

	fmt.Print("How much do you want to ", choice, "? ")

	var amount float64
	fmt.Scan(&amount)

	if amount <= 0 || amount > accountBalance {
		fmt.Println("Invalid amount entered.")
		return
	}

	if withdraw {
		accountBalance -= amount
	} else {
		accountBalance += amount
	}

	displayBalance(accountBalance)
	writeBalanceToFile(accountBalance)
}

func displayBalance(accountBalance float64) {
	fmt.Printf("Your balance is %.2f\n", accountBalance)
}
