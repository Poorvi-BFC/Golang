// Maps are used to store data values in key:value pairs;
// are an unordered and changeable collection that does not allow duplicates.
// syntax: map[KeyType]ValueType{key1: value1, key2: value2, key3: value3}
// can be created using the make() function or map literal.

package main

import "fmt"

func main() {
	var a = make(map[string]string) // The map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964" // a is no longer empty

	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

// overwriting is possible inside scope.
// to remove :{ delete(map_name, key) }
// to check for any event :{ val, ok :=map_name[key] } ; gives o/p in bool.
