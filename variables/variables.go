package main

import "fmt"

func main() {

	// Must use double quotes for values, singles are not supported
	var a = "initial"
	fmt.Println(a)

	// Can assign multiple variables of the same type in the same line
	var b, c int = 1, 2
	fmt.Println(b, c)

	// This is equivlant to:
	bb, cc := 1, 2
	fmt.Println(bb, cc)

	// Also this:
	var (
		bbb = 1
		ccc = 2
	)
	fmt.Println(bbb, ccc)

	var d = true
	fmt.Println(d)

	// An int is set to 0 by default
	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	// Empty strings are set to ""
	var g string
	fmt.Println(g)
	fmt.Println(g, "Hello?")
}
