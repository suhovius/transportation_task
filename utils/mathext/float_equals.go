package mathext

import (
	"math"
)

const epsilon float64 = 1e-8

// FloatEquals checks equality of float64 values with degree of accuracy
// defined at epsilon constant
func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
