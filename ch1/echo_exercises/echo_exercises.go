// Echo_exercises prints its command line arguements.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// e1_1 prints all command line arguements
func e1_1() {
	s := time.Now()
	fmt.Println("e1_1:")
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Printf("%.8fs elapsed\n", time.Since(s).Seconds())
}

// e1_2 prints all command line arguements and their indexes
func e1_2() {
	s := time.Now()
	fmt.Println("e1_2")
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
	fmt.Printf("%.8fs elapsed\n", time.Since(s).Seconds())
}

func main() {
	e1_1()
	fmt.Println()
	e1_2()
}
