// Maps are go's dictionary. Objects made up of key-value pairs

package main

import "fmt"

func main() {
	// Creating an empty map, you have to again declare the types of both the key and value
	m := make(map[string]int)

	// Set the values just like you would in python. Remember to use double quotes
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Can also get the length of a map object
	fmt.Println("len:", len(m))

	// Can rip out values from the map using delete
	delete(m, "k2")
	fmt.Println("map:", m)

	// You can check to see if a value exists in a map by checking for two values. The second value will be whether or not the key exists
	_, prs := m["k2"] // ignores the first value returned
	fmt.Println("prs:", prs)

	// declare and init new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
