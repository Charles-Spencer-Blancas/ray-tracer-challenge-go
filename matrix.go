package main

type Matrix struct {
	Values [][]float64
	Height int64
	Width  int64
}

func matrixZeroes(n int64, m int64) Matrix {
	a := [][]float64{}
	for i := int64(0); i < n; i++ {
		for j := int64(0); j < m; j++ {
			a[i][j] = 0.0
		}
	}

	return Matrix{a, n, m}
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

func matrix4x4Multiply(aM Matrix, bM Matrix) Matrix {
	a := aM.Values
	b := bM.Values
	out := [][]float64{}
	for i := range out {
		for j := range out {
			out[i][j] = a[i][0]*b[0][j] +
				a[i][1]*b[1][j] +
				a[i][2]*b[2][j] +
				a[i][3]*b[3][j]
		}
	}

	return Matrix{out, 4, 4}
}
