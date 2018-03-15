package main

import (
	"fmt"
)

func main() {
	a, b, sum := 1, 2, 0
	for b < 4000000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	fmt.Println(sum)
}
