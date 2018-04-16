// geometry.go
package main

import (
	"fmt"
	"log"

	"github.com/fuzzylimes/geometry/rectangle"
)

var rectLen, rectWidth float64 = 6, 7

func init() {
	fmt.Println("main package initialized...")
	if rectLen <= 0 {
		log.Fatal("rectLen cannot be less than or equal to 0")
	}
	if rectWidth <= 0 {
		log.Fatal("rectWidth cannot be less than or equal to 0")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")

	fmt.Printf("Area of rectangle: %.2f\n", rectangle.Area(rectLen, rectWidth))

	fmt.Printf("Diagonal of the rectangle: %.2f\n", rectangle.Diagonal(rectLen, rectWidth))
}
