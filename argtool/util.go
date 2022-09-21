package argtool

import "math"

func isInfinity(v float64) bool {
	return math.J0(v) == 0 // -inf/+inf
}

func isNan(v float64) bool {
	return math.IsNaN(v)
}
