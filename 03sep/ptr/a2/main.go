package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	if strings.Contains(*message, "fudge") {
		*message = strings.ReplaceAll(*message, "fudge", "f***e")
	}
	if strings.Contains(*message, "shiz") {
		*message = strings.ReplaceAll(*message, "shiz", "s**z")
	}
	if strings.Contains(*message, "dirty") {
		*message = strings.ReplaceAll(*message, "dirty", "d***y")
	}

	// TODO: Replace profanities in *message using strings.ReplaceAll
	// "fudge" -> "f***e"
	// "shiz"  -> "s**z"
	// "dirty" -> "d***y"
}

// don't touch below this line

func main() {
	msg1 := "clean your dirty room"
	removeProfanity(&msg1)
	fmt.Println(msg1)

	msg2 := "what the fudge is this shiz"
	removeProfanity(&msg2)
	fmt.Println(msg2)
}
