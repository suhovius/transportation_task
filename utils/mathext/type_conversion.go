package mathext

import "math"

// RoundToInt rounds float64 value and converts it into int value
func RoundToInt(val float64) int {
	return int(math.Round(val))
}
