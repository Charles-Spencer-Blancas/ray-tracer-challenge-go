package main

import (
	"fmt"
)

type Matrix struct {
	Values [][]float64
	Height int64
	Width  int64
}

func matrixConstruct(values [][]float64) Matrix {
	return Matrix{values, int64(len(values)), int64(len(values[0]))}
}

func matrixEqual(a Matrix, b Matrix) bool {
	if a.Height != b.Height || a.Width != b.Width {
		return false
	}

	for i := range a.Values {
		for j := range a.Values[i] {
			if !floatEqual(a.Values[i][j], b.Values[i][j]) {
				return false
			}
		}
	}
	return true
}

func matrix4x4Multiply(aM Matrix, bM Matrix) (Matrix, error) {
	if aM.Height != 4 || aM.Width != 4 || bM.Height != 4 || bM.Width != 4 {
		return Matrix{}, fmt.Errorf("can only multiply 4 x 4 matrices but got %d x %d * %d x %d", aM.Height, aM.Width, bM.Height, bM.Width)
	}
	a := aM.Values
	b := bM.Values
	out := make([][]float64, 4)
	for i := range a {
		out[i] = make([]float64, 4)
		for j := range a {
			out[i][j] = a[i][0]*b[0][j] +
				a[i][1]*b[1][j] +
				a[i][2]*b[2][j] +
				a[i][3]*b[3][j]
		}
	}

	return Matrix{out, 4, 4}, nil
}

func matrix4x4TupleMultiply(a Matrix, t Tuple) (Tuple, error) {
	if a.Height != 4 || a.Width != 4 {
		return Tuple{}, fmt.Errorf("can only multiply 4 x 4 matrix but got %d x %d", a.Height, a.Width)
	}
	tup := [4]float64{t.X, t.Y, t.Z, t.W}
	out := [4]float64{}
	for i := range a.Values {
		out[i] = a.Values[i][0]*tup[0] + a.Values[i][1]*tup[1] + a.Values[i][2]*tup[2] + a.Values[i][3]*tup[3]
	}

	return Tuple{out[0], out[1], out[2], out[3]}, nil
}

func matrixConstructIdentity(n int64) Matrix {
	a := make([][]float64, n)

	for i := int64(0); i < n; i++ {
		a[i] = make([]float64, n)
		for j := int64(0); j < n; j++ {
			if i == j {
				a[i][j] = 1
			} else {
				a[i][j] = 0
			}
		}
	}
	return Matrix{a, n, n}
}
