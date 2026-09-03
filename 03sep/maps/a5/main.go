package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		if len(name) == 0 {
			continue
		}
		firstChar := rune(name[0])

		// TODO: Check if the inner map exists for firstChar; if not, initialize it
		// TODO: Increment the count for the name in
		// the inner map
		b := counts[firstChar]
		if b == nil {
			b = make(map[string]int)
			counts[firstChar] = b
		}
		b[name]++

	}

	return counts
}

// don't touch below this line.

func test(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %d names...\n", len(names))
	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c] [%s]: %d\n\n", initial, name, count)
}

func main() {
	names := []string{"billy", "billy", "bob", "joe", "Poorvi"}
	test(names, 'b', "billy")
	test(names, 'b', "bob")
	test(names, 'j', "joe")
	test(names, 'P', "Poorvi")
}
