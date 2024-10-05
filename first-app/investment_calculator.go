package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6
	var investmentAmount float64
	var investmentRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Investment Rate: ")
	fmt.Scan(&investmentRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	// futureValue = investmentAmount * (1 + investmentRate / 100) ^ years
	futureValue := investmentAmount * math.Pow(1+(investmentRate/100), years)
	// futureRealValue = futureValue / (1 + inflationRate / 100) ^ years
	futureRealValue := futureValue / math.Pow(1+(inflationRate/100), years)

	fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
}
