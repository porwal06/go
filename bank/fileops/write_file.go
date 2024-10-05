package fileops

import (
	"fmt"
	"os"
)

func WriteBalanceToFile(balance float64) {
	//Convert float to string
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
