package main

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func point(X float64, Y float64, Z float64) Tuple {
	return Tuple{X, Y, Z, 1.0}
}

func vector(X float64, Y float64, Z float64) Tuple {
	return Tuple{X, Y, Z, 0.0}
}

func isPoint(t Tuple) bool {
	return t.W == 1.0
}

func isVector(t Tuple) bool {
	return t.W == 0.0
}
