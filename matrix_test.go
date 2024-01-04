package main

import "testing"

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
		{9, 10, 11, 12},
		{13, 14, 15, 16},
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
		t.Errorf("Expected %v * %v to be %v but got %v", a, b, expectedVals, c)
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
