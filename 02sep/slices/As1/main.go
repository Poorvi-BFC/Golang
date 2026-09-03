/*Assignment: Array Messages
Objective: Write a function that returns a fixed-size array containing three reminder messages, and print them out using a for loop.

Instructions:
Create a function named getMessages() that returns an array of 3 strings ([3]string).

Inside getMessages(), return an array initialized with these 3 strings:

"Your bill is due"

"Payment reminder"

"Final notice"

In main(), call getMessages() and store the returned array.

Iterate over the array using a for loop (using len()) and print each message in the format:

Sending: "message"*/

package main

import "fmt"

func getMessages() [3]string {
	// TODO: Return an array of 3 strings
	return [3]string{"Your bill is due", "Payment reminder", "Final notice"}

}

func main() {
	messages := getMessages()

	// TODO: Loop through the messages array and print each item
	for i := 0; i < len(messages); i++ {
		fmt.Printf("Sending: \"%s\"\n", messages[i])
	}

}
