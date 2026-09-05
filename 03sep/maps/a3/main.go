package main

import "fmt"

func isValidKey(typeName string) bool {
	// TODO: Return false for invalid map key types ("slice", "map", "func")
	// TODO: Return true for valid map key types
	if typeName == "slice" || typeName == "map" || typeName == "func" {
		return false
	}
	return true
}

// don't touch below this line

func main() {
	tests := []string{"string", "int", "slice", "map", "func"}

	for _, t := range tests {
		fmt.Println(t, "-> valid key?", isValidKey(t))
	}
}
