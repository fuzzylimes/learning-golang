// Package rectangle defines calculations for rectangle
package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized...")
}

// Area calulates the area of a rectangle
func Area(length, width float64) float64 {
	return length * width
}

// Diagonal calculates the diagonal of a rectangle
func Diagonal(length, width float64) float64 {
	return math.Sqrt((length * length) + (width * width))
}
