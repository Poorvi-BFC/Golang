package main

import "fmt"

func sum(nums ...int) int { //for int type only
	//func sum(nums ...interface) int { // for any type
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	result := sum(2, 3, 4, 5)
	fmt.Println(result)

}
