// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"fmt"
)

// Celsius is for celsius temps
type Celsius float64

// Fahrenheit is for fahrenheit temps
type Fahrenheit float64

// Kelvin is for kelvin temps
type Kelvin float64

// Constant temp values
const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
	KelvinC      Kelvin  = 273.15
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
