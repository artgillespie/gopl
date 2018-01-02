package weightconv

const c = 0.453592
const invC = 1.0 / c

// LToK converts pounds to kilograms
func LToK(l float64) float64 { return l * c }

// KToL converts kilograms to pounds
func KToL(k float64) float64 { return k * invC }
