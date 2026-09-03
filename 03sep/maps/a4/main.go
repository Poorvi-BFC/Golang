package main

import (
	"fmt"
)

func getCounts(userIDs []string) map[string]int {
	counts := make(map[string]int)

	for _, id := range userIDs {
		// TODO: Increment the count for this user ID in the map
		counts[id]++
	}

	return counts
}

// don't touch below this line

func main() {
	ids := []string{"cersei", "jaime", "cersei", "tyrion", "jaime", "cersei"}
	res := getCounts(ids)
	fmt.Println(res)
}
