package lengthconv

// FToM converts feet to meters
func FToM(f float64) float64 { return f * .3048 }

// MToF converts meters to feet
func MToF(m float64) float64 { return m * 3.28084 }
