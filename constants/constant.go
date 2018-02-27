package main

// You can do multiple imports like so (similar to how you can do variables)
import (
	"fmt"
	"math"
)

// Can have two types of constants: typed and untyped. Below is an example of a typed constant
const s string = "constant"

// And this would be an untyped:
const t = "value"

// Basically the untyped is just a value that you are free to do what you want to with it. It isn't restricted to the types of rueles typed values are.
// Here's a good blog post with more info on all of this: https://blog.golang.org/constants

func main() {
	// const variables defined outside of main function are accesible
	fmt.Println(s)

	// Cannot set a new value to a consant
	// s = "another value"

	// But you CAN create a new variable with the same name???
	s := "A new value"
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
