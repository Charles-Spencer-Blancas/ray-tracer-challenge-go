package main

import "math"

const EPSILON = float64(0.00001)

func floatEqual(a float64, b float64) bool {
	return math.Abs(a-b) < EPSILON
}
