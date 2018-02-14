// Structs are basically Go's classes

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 32})
	fmt.Println(person{name: "Fread"})
	fmt.Println(&person{name: "Ann", age: 34})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.name, s.age)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
