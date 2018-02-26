// Echo1 prints its command-line arguements.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ", "
	}
	fmt.Printf("%d arguments supplied:\n", len(os.Args)-1)
	fmt.Println(s)
}
