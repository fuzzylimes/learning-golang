// A look at arrays in Go

package main

import "fmt"

func main() {

	// Arrays are like arrays in python, only they are of a set length. Arrays that aren't of a specific length are called a slice
	// Arrays are initialized with 0 values
	var a [5]int
	fmt.Println("empty int array:", a)

	a[4] = 100
	fmt.Println("Set value:", a)
	fmt.Println("Get value:", a[4])

	// Can use len() just like in python to get the length of an array
	fmt.Println("len:", len(a))

	// len() works the same way for strings too
	fmt.Println("length of string:", len("Testing"))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Initialize with values:", b)

	// Creating 2d arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	var c [6]string
	fmt.Println("Empty string array:", c)
}
