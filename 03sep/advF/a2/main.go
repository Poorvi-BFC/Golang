package main

import "fmt"

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

// don't touch below this line

func addSignature(message string) string {
	return message + " Kind regards."
}

func main() {
	messages := []string{
		"Thanks for getting back to me.",
		"Great to see you again.",
	}

	formatted := getFormattedMessages(messages, addSignature)
	fmt.Println(formatted)
}
