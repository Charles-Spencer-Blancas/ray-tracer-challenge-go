package main

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func tupleEqual(a Tuple, b Tuple) bool {
	return floatEqual(a.X, b.X) && floatEqual(a.Y, b.Y) && floatEqual(a.Z, b.Z) && floatEqual(a.W, b.W)
}

func tupleAdd(a Tuple, b Tuple) Tuple {
	return Tuple{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

func tupleSubtract(a Tuple, b Tuple) Tuple {
	return Tuple{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
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
