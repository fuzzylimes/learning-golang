// Taking a look at how to write functions in Go

package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	fmt.Printf("res is type %T\n", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
