package tempconv

// CToF converts celsius to fahrenheit
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts fahrenheit to celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZero)
}

// FToK converts fahrenheit to kelvin
func FToK(f Fahrenheit) Kelvin {
	return Kelvin(CToK(FToC(f)))
}

// KToC converts kelvin to celsius
func KToC(k Kelvin) Celsius {
	return Celsius(k - KelvinC)
}

// KToF converts kelvin to fahrenheit
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(CToF(KToC(k)))
}
