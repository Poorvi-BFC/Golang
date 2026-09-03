package main

import "fmt"

func deleteUser(userMap map[string]user, name string) bool {
	// TODO: check if user exists
	_, ok := userMap[name]
	if !ok {
		return false
	}
	// TODO: delete user if found
	delete(userMap, name)
	// TODO: return true if deleted, false if not found
	return true
}

// don't touch below this line

type user struct {
	name  string
	phone int
}

func main() {
	users := map[string]user{
		"Alice": {name: "Alice", phone: 123},
		"Bob":   {name: "Bob", phone: 456},
	}

	deleted := deleteUser(users, "Alice")
	fmt.Println("Deleted:", deleted)
	fmt.Println("Remaining users:", users)
}

// Example:
// map before: {"Alice": user{...}}
// deleteUser(userMap, "Alice")
// map after:  {}
// result: true
