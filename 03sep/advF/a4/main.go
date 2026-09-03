package main

import "fmt"

func adder() func(int) int {
	// TODO: Define a variable to keep track of running total sum
	// TODO: Return an anonymous function that accepts an int, adds it to sum, and returns sum
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// don't touch below this line

type emailBill struct {
	costInPennies int
}

func test(bills []emailBill) {
	fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("Bill #%d: %d pennies\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func main() {
	test([]emailBill{
		{costInPennies: 100},
		{costInPennies: 200},
		{costInPennies: 300},
	})
}
