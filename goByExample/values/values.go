package main

import "fmt"

func main() {

	// Can combine strings just like in python
	fmt.Println("go" + "lang")

	// Also can combine different types when printing, just like python.
	// Go automagically adds a space when doing a comma in a println
	fmt.Println("1+2 =", 1+2)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println("hello", "world")

	// Basic boolean stuff
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
