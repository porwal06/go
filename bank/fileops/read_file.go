package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func ReadBalanceFromFile() float64 {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println("Error reading balance from file:", err)
		return 0
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		fmt.Println("Error parsing balance from file:", err)
		return 0
	}
	return balance
}
