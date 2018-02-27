// range can be used to iterate over elements

package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	// first value is index, second is the number
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over the keys
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points. First value is the byte index of the rune, second is the rune value itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
