package main

import (
	"errors"
	"fmt"
)

// Create struct
// We can create it as seperate package & import here. For that check bank app
type Finance struct {
	revenue, expenses, taxRate float64
}

// attach function to struct. Only difference from normal function we define "receiver" type for struct function.
// "f Finance" called as receiver type
func (f Finance) CalculateFinance() (ebt, profit, ratio float64) {
	ebt = f.revenue - f.expenses
	profit = ebt * (1 - f.taxRate/100)
	ratio = ebt / profit
	// return ebt, profit, ratio OR simply
	return
}

// Create constructor/creator function to return struct with pointer & do validation
func NewFinance(revenue, expenses, taxRate float64) (*Finance, error) {
	if revenue == 0.0 || expenses == 0.0 || taxRate == 0.0 {
		return nil, errors.New("invalid input")
	}

	return &Finance{
		revenue:  revenue,
		expenses: expenses,
		taxRate:  taxRate,
	}, nil
}

// Main function

func main() {
	var calculateCustomFinance *Finance
	var err error

	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter taxRate: ")

	//Pass value to struct
	// calculateCustomFinance = Finance{
	// 	revenue:  revenue,
	// 	expenses: expenses,
	// 	taxRate:  taxRate,
	// }

	// Use constructor function to pass value to struct
	calculateCustomFinance, err = NewFinance(revenue, expenses, taxRate)
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateCustomFinance.CalculateFinance()

	fmt.Printf("Total ebt : %.0f\n", ebt)
	fmt.Printf("Total Profit : %.0f\n", profit)
	fmt.Printf("Total ratio : %.2f\n", ratio)
}

// func calculateFinance(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
// 	ebt = revenue - expenses
// 	profit = ebt * (1 - taxRate/100)
// 	ratio = ebt / profit

// 	return
// }

func getUserInput(input string) (value float64) {
	fmt.Print(input)
	fmt.Scanln(&value)
	return
}
