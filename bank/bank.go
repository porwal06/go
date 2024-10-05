package main

import (
	"fmt"

	"example.com/bank/fileops"           // Importing the custom package 'fileops' from the 'example.com/bank' directory
	"github.com/Pallinder/go-randomdata" //Importing the third party package 'github.com/Pallinder/go-randomdata' for generating random phone numbers
)

func main() {
	fmt.Println("Welcome to Go Bank !!")
	fmt.Println("Reach 24/7", randomdata.PhoneNumber())

	// var accountBalance float64 = 1000.00

	var accountBalance = fileops.ReadBalanceFromFile()

	for {
		displayMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			checkBalance(accountBalance)
		case 2:
			accountBalance = deposit(accountBalance)
		case 3:
			accountBalance = withdraw(accountBalance)
		case 4:
			fmt.Println("Thank you for using our bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			return
		}
	}

}

func displayMenu() {
	fmt.Println("1. Account Balance")
	fmt.Println("2. Deposit")
	fmt.Println("3. Withdraw")
	fmt.Println("4. Exit")
}

func checkBalance(accountBalance float64) {
	fmt.Printf("Your current balance is: %.2f\n", accountBalance)
}
func deposit(accountBalance float64) (newBalance float64) {
	var amount float64
	newBalance = accountBalance
	fmt.Print("Enter the amount to deposit: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Invalid amount. Please try again.")
		return
	}
	newBalance += amount
	fmt.Printf("Deposited $%.2f successfully. Now total balance is $%.2f\n", amount, newBalance)
	fileops.WriteBalanceToFile(newBalance)
	return
}
func withdraw(accountBalance float64) (newBalance float64) {
	var amount float64
	newBalance = accountBalance
	fmt.Print("Enter the amount to withdraw: ")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Invalid amount. Please try again.")
		return
	}
	if amount > newBalance {
		fmt.Println("Insufficient balance. Please try again.")
		return
	}
	newBalance -= amount
	fmt.Printf("Withdrew $%.2f successfully. Now total balance $%.2f\n", amount, newBalance)
	fileops.WriteBalanceToFile(newBalance)
	return
}
