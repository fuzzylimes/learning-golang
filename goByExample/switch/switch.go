// Switch statements seem pretty normal in go, though there are a few neat things you can do with them.

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("three")
	}

	// can sepparate multiple expressions in the same case statement with a comma
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Using switch statements to handle if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
