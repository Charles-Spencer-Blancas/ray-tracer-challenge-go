package main

import (
	"testing"
)

func TestConstructAndInspect4x4(t *testing.T) {
	values := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	a := Matrix{values, 4, 4}
	coords := [7][2]int64{
		{0, 0},
		{0, 3},
		{1, 0},
		{1, 2},
		{2, 2},
		{3, 0},
		{3, 2},
	}
	expected := [7]float64{1, 4, 5.5, 7.5, 11, 13.5, 15.5}

	for i := range coords {
		if !floatEqual(a.Values[coords[i][0]][coords[i][1]], expected[i]) {
			t.Errorf("Expected a[%d][%d] to be %f but got %f", coords[i][0], coords[i][1], expected[i], a.Values[coords[i][0]][coords[i][1]])
		}
	}
}

func TestConstructAndInspect2x2(t *testing.T) {
	values := [][]float64{
		{-3, 5},
		{1, -2},
	}
	a := Matrix{values, 2, 2}
	coords := [4][2]int64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	expected := [4]float64{-3, 5, 1, -2}

	for i := range coords {
		if !floatEqual(a.Values[coords[i][0]][coords[i][1]], expected[i]) {
			t.Errorf("Expected a[%d][%d] to be %f but got %f", coords[i][0], coords[i][1], expected[i], a.Values[coords[i][0]][coords[i][1]])
		}
	}
}

func TestConstructAndInspect3x3(t *testing.T) {
	values := [][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}
	a := Matrix{values, 3, 3}
	coords := [][]int64{
		{0, 0},
		{1, 1},
		{2, 2},
	}
	expected := [3]float64{-3, -2, 1}

	for i := range coords {
		if !floatEqual(a.Values[coords[i][0]][coords[i][1]], expected[i]) {
			t.Errorf("Expected a[%d][%d] to be %f but got %f", coords[i][0], coords[i][1], expected[i], a.Values[coords[i][0]][coords[i][1]])
		}
	}
}

func TestEqualityIdenticalMatrices(t *testing.T) {
	aValues := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	a := Matrix{aValues, 4, 4}
	bValues := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	b := Matrix{bValues, 4, 4}
	if !matrixEqual(a, b) {
		t.Errorf("Expected a and be to be equal but got that they are not")
	}
}

func TestEqualityDifferentMatrices(t *testing.T) {
	aValues := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	a := Matrix{aValues, 4, 4}
	bValues := [][]float64{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{10, 10, 11, 12},
		{15, 14, 15, 12},
	}
	b := Matrix{bValues, 4, 4}
	if matrixEqual(a, b) {
		t.Errorf("Expected a and be to not be equal but got that they are")
	}
}

func TestMultiplyTwo4x4(t *testing.T) {
	aValues := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	a := Matrix{aValues, 4, 4}
	bValues := [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	b := Matrix{bValues, 4, 4}
	c, err := matrix4x4Multiply(a, b)
	if err != nil {
		t.Fatal(err)
	}
	expectedVals := [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}
	if !matrixEqual(c, matrixConstruct(expectedVals)) {
		t.Errorf("Expected %v * %v to be %v but got %v", a, b, matrixConstruct(expectedVals), c)
	}
}

func TestMultiply4x4MatrixByTuple(t *testing.T) {
	vals := [][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	}
	a := matrixConstruct(vals)
	b := Tuple{1, 2, 3, 1}
	got, err := matrix4x4TupleMultiply(a, b)
	if err != nil {
		t.Fatal(err)
	}
	expected := Tuple{18, 24, 33, 1}

	if !tupleEqual(got, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", a.Values, b, expected, got)
	}
}

func TestMultiply4x4MatrixByIdentity(t *testing.T) {
	vals := [][]float64{
		{0, 1, 2, 4},
		{1, 2, 4, 8},
		{2, 4, 8, 16},
		{4, 8, 16, 32},
	}
	a := matrixConstruct(vals)
	i := matrixConstructIdentity(4)
	got, err := matrix4x4Multiply(a, i)
	if err != nil {
		t.Fatal(err)
	}
	if !matrixEqual(a, got) {
		t.Errorf("Expected %v * identity to be %v but got %v", a, a, got)
	}
}

func TestMultiply4x4IdentityWithTuple(t *testing.T) {
	a := Tuple{1, 2, 3, 4}
	i := matrixConstructIdentity(4)
	got, err := matrix4x4TupleMultiply(i, a)
	if err != nil {
		t.Fatal(err)
	}
	if !tupleEqual(got, a) {
		t.Errorf("Expected %v * %v to be %v but got %v", a, i, a, got)
	}
}

func TestTransposeMatrix(t *testing.T) {
	vals := [][]float64{
		{0, 9, 3, 0},
		{9, 8, 0, 8},
		{1, 8, 5, 3},
		{0, 0, 5, 8},
	}
	a := matrixConstruct(vals)
	transpose := matrixTranspose(a)
	expected := matrixConstruct([][]float64{
		{0, 9, 1, 0},
		{9, 8, 8, 0},
		{3, 0, 5, 5},
		{0, 8, 3, 8},
	})

	if !matrixEqual(transpose, expected) {
		t.Errorf("Expected transpose(%v) to be %v but got %v", a, expected, transpose)
	}
}

func TestTransposeIdentityMatrix(t *testing.T) {
	i := matrixConstructIdentity(4)
	transpose := matrixTranspose(i)

	if !matrixEqual(transpose, i) {
		t.Errorf("Expected transpose(%v) to be %v but got %v", i, i, transpose)
	}
}

func TestDeterminant2x2Matrix(t *testing.T) {
	a := matrixConstruct([][]float64{
		{1, 5},
		{-3, 2},
	})
	det, err := matrix2x2Determinant(a)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(det, 17) {
		t.Errorf("Expected det(%v) to be %f but got %f", a, 17., det)
	}
}

func TestMinorOf3x3(t *testing.T) {
	a := matrixConstruct([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	b, err := matrixSubmatrix(a, 1, 0)
	if err != nil {
		t.Fatal(err)
	}
	det, err := matrix2x2Determinant(b)
	if err != nil {
		t.Fatal(err)
	}
	minor, err := matrix3x3Minor(a, 1, 0)
	if err != nil {
		t.Fatal(err)
	}

	if !floatEqual(det, 25) {
		t.Errorf("Expected det(%v) to be %f but got %f", b, 25.0, det)
	}
	if !floatEqual(minor, 25) {
		t.Errorf("Expected minor(%v, 1, 0) to be %f but got %f", a, 25.0, minor)
	}
}

func TestCofactorOf3x3(t *testing.T) {
	a := matrixConstruct([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})

	minorA00, err := matrix3x3Minor(a, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(minorA00, -12) {
		t.Errorf("Expected minor(%v, 0, 0) to be %f but got %f", a, -12., minorA00)
	}

	cofactorA00, err := matrixCofactor(a, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(cofactorA00, -12) {
		t.Errorf("Expected cofactor(%v, 0, 0) to be %f but got %f", a, -12., cofactorA00)
	}

	minorA10, err := matrix3x3Minor(a, 1, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(minorA10, 25) {
		t.Errorf("Expected minor(%v, 1, 0) to be %f but got %f", a, 25., minorA00)
	}

	cofactorA10, err := matrixCofactor(a, 1, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(cofactorA10, -25) {
		t.Errorf("Expected cofactor(%v, 1, 0) to be %f but got %f", a, -25., cofactorA10)
	}
}

func TestDeterminant3x3Matrix(t *testing.T) {
	a := matrixConstruct([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})

	a00, err := matrixCofactor(a, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(a00, 56) {
		t.Errorf("Expected cofactor(%v, 0, 0) to be 56 but got %f", a, a00)
	}

	a01, err := matrixCofactor(a, 0, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(a01, 12) {
		t.Errorf("Expected cofactor(%v, 0, 1) to be 12 but got %f", a, a01)
	}

	a02, err := matrixCofactor(a, 0, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(a02, -46) {
		t.Errorf("Expected cofactor(%v, 0, 2) to be -46 but got %f", a, a02)
	}

	det, err := matrixDeterminant(a)
	if err != nil {
		t.Fatal(err)
	}

	if !floatEqual(det, -196) {
		t.Errorf("Expected det(%v) to be -196 but got %f", a, det)
	}
}

func TestDeterminant4x4Matrix(t *testing.T) {
	a := matrixConstruct([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})

	a00, err := matrixCofactor(a, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(a00, 56) {
		t.Errorf("Expected cofactor(%v, 0, 0) to be 56 but got %f", a, a00)
	}

	a01, err := matrixCofactor(a, 0, 1)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(a01, 12) {
		t.Errorf("Expected cofactor(%v, 0, 1) to be 12 but got %f", a, a01)
	}

	a02, err := matrixCofactor(a, 0, 2)
	if err != nil {
		t.Fatal(err)
	}
	if !floatEqual(a02, -46) {
		t.Errorf("Expected cofactor(%v, 0, 2) to be -46 but got %f", a, a02)
	}

	det, err := matrixDeterminant(a)
	if err != nil {
		t.Fatal(err)
	}

	if !floatEqual(det, -196) {
		t.Errorf("Expected det(%v) to be -196 but got %f", a, det)
	}
}
