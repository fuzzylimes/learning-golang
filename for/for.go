// Look at for loops in Go. For loops are the only looping construct in Go.

package main

import "fmt"

func main() {

	// Basic
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Classic looping
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Without any conditions, will loop forever until a break is hit. So, think while loop.
	for {
		fmt.Println("loop")
		break
	}

	// Can continue to the next iterition of the loop
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
