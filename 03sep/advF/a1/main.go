package main

import "fmt"

func aggregate(a, b, c int, formatter func(int) string) string {
	// TODO: Calculate sum of a, b, and c
	// TODO: Format the sum using formatter function and return it
	sum := a + b + c
	return formatter(sum)
}

// don't touch below this line

func formatCoins(amount int) string {
	return fmt.Sprintf("You have %d coins", amount)
}

func formatPoints(amount int) string {
	return fmt.Sprintf("Score: %d pts", amount)
}

func main() {
	coins := aggregate(5, 10, 15, formatCoins)
	points := aggregate(100, 200, 50, formatPoints)

	fmt.Println(coins)
	fmt.Println(points)
}
