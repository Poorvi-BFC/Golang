package main

import "fmt"

func logActivity(activity string) {
	fmt.Println("Starting activity:", activity)
	defer fmt.Println("Finished activity: ", activity)

	// TODO: Defer printing "Finished activity: " + activity

	fmt.Println("Performing step 1")
	fmt.Println("Performing step 2")
}

// don't touch below this line

func main() {
	logActivity("Data Backup")
}
