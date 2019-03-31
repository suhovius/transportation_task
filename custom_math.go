package main

import (
	"math"
)

const epsilon float64 = 1e-8

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
