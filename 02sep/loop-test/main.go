package main

import (
	"fmt"
)

// COMPLETE THE calculateInvestment() FUNCTION
func calculateInvestment(initialDeposit float64, annualContribution float64, interestRate float64, years int) float64 {
	balance := initialDeposit
	rateAsDecimal := interestRate / 100.0

	for i := 0; i < years; i++ {
		balance += balance * rateAsDecimal
		balance += annualContribution
	}

	// Write your loop here

	return balance
}

func test(initialDeposit, annualContribution, interestRate float64, years int) {
	total := calculateInvestment(initialDeposit, annualContribution, interestRate, years)
	fmt.Printf("After %d years, balance is: $%.2f\n", years, total)
}

func main() {
	// Test Case: $1,000 initial, $500 added yearly, 5% interest, over 3 years
	// Expected Output: $2,831.63
	test(1000.0, 500.0, 5.0, 3)
}
