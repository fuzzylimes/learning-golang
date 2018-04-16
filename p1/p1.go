package main

import (
	"fmt"
)

func multiples(x int) ([]int, int) {
	var l []int
	var s int
	for i := 1; i < x; i++ {
		if i%3 == 0 || i%5 == 0 {
			l = append(l, i)
			s += i
		}
	}
	return l, s

}

func main() {
	l, s := multiples(1000)
	fmt.Println("Matching numbers:\n", l)
	fmt.Println("Sum of numbers:", s)
}
