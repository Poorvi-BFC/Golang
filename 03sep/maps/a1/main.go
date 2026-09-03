package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{name: name, phoneNumber: phoneNumber}
	}

	return userMap, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int
}

func main() {
	names := []string{"Alice", "Bob"}
	phones := []int{123, 456}

	m, err := getUserMap(names, phones)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(m)
}
