package main

import (
	"math"
	"testing"
)

func TestMultiplyByTranslationMatrix(t *testing.T) {
	transform := translation(5, -3, 2)
	p := point(-3, 4, 5)

	translated, err := matrix4x4TupleMultiply(transform, p)
	if err != nil {
		t.Fatal(err)
	}

	if !tupleEqual(translated, point(2, 1, 7)) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, p, point(2, 1, 7), translated)
	}
}

func TestMultiplyByInverseOfTranslationMatrix(t *testing.T) {
	transform, err := matrixInverse(translation(5, -3, 2))
	if err != nil {
		t.Fatal(err)
	}
	p := point(-3, 4, 5)

	translated, err := matrix4x4TupleMultiply(transform, p)
	if err != nil {
		t.Fatal(err)
	}

	if !tupleEqual(translated, point(-8, 7, 3)) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, p, point(-8, 7, 3), translated)
	}
}

func TestTranslationDoesNotAffectVectors(t *testing.T) {
	transform := translation(5, -3, 2)
	v := vector(-3, 4, 5)

	translated, err := matrix4x4TupleMultiply(transform, v)
	if err != nil {
		t.Fatal(err)
	}

	if !tupleEqual(translated, v) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, v, v, translated)
	}
}

func TestScalingPoint(t *testing.T) {
	transform := scaling(2, 3, 4)
	p := point(-4, 6, 8)

	translated, err := matrix4x4TupleMultiply(transform, p)
	if err != nil {
		t.Fatal(err)
	}

	expected := point(-8, 18, 32)

	if !tupleEqual(translated, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, p, expected, translated)
	}
}

func TestScalingVector(t *testing.T) {
	transform := scaling(2, 3, 4)
	v := vector(-4, 6, 8)

	translated, err := matrix4x4TupleMultiply(transform, v)
	if err != nil {
		t.Fatal(err)
	}

	expected := vector(-8, 18, 32)

	if !tupleEqual(translated, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, v, expected, translated)
	}
}

func TestScalingInverse(t *testing.T) {
	init := scaling(2, 3, 4)
	transform, err := matrixInverse(init)
	if err != nil {
		t.Fatal(err)
	}
	v := vector(-4, 6, 8)

	translated, err := matrix4x4TupleMultiply(transform, v)
	if err != nil {
		t.Fatal(err)
	}

	expected := vector(-2, 2, 2)

	if !tupleEqual(translated, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, v, expected, translated)
	}
}

func TestReflectingIsScalingByNegative(t *testing.T) {
	transform := scaling(-1, 1, 1)
	p := point(2, 3, 4)

	translated, err := matrix4x4TupleMultiply(transform, p)
	if err != nil {
		t.Fatal(err)
	}

	expected := point(-2, 3, 4)

	if !tupleEqual(translated, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, p, expected, translated)
	}
}

func TestRotateAroundX(t *testing.T) {
	halfQuarter := rotationX(math.Pi / 4)
	fullQuarter := rotationX(math.Pi / 2)
	p := point(0, 1, 0)

	translatedHalf, err := matrix4x4TupleMultiply(halfQuarter, p)
	if err != nil {
		t.Fatal(err)
	}

	expectedHalf := point(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !tupleEqual(translatedHalf, expectedHalf) {
		t.Errorf("Expected %v * %v to be %v but got %v", halfQuarter, p, expectedHalf, translatedHalf)
	}

	translatedFull, err := matrix4x4TupleMultiply(fullQuarter, p)
	if err != nil {
		t.Fatal(err)
	}

	expectedFull := point(0, 0, 1)

	if !tupleEqual(translatedFull, expectedFull) {
		t.Errorf("Expected %v * %v to be %v but got %v", fullQuarter, p, expectedFull, translatedFull)
	}
}

func TestInverseXRotation(t *testing.T) {
	init := rotationX(math.Pi / 4)
	transform, err := matrixInverse(init)
	if err != nil {
		t.Fatal(err)
	}
	p := point(0, 1, 0)

	translated, err := matrix4x4TupleMultiply(transform, p)
	if err != nil {
		t.Fatal(err)
	}

	expected := point(0, math.Sqrt(2)/2, -math.Sqrt(2)/2)

	if !tupleEqual(translated, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", transform, p, expected, translated)
	}
}

func TestRotateAroundY(t *testing.T) {
	halfQuarter := rotationY(math.Pi / 4)
	fullQuarter := rotationY(math.Pi / 2)
	p := point(0, 0, 1)

	translatedHalf, err := matrix4x4TupleMultiply(halfQuarter, p)
	if err != nil {
		t.Fatal(err)
	}

	expectedHalf := point(math.Sqrt(2)/2, 0, math.Sqrt(2)/2)

	if !tupleEqual(translatedHalf, expectedHalf) {
		t.Errorf("Expected %v * %v to be %v but got %v", halfQuarter, p, expectedHalf, translatedHalf)
	}

	translatedFull, err := matrix4x4TupleMultiply(fullQuarter, p)
	if err != nil {
		t.Fatal(err)
	}

	expectedFull := point(1, 0, 0)

	if !tupleEqual(translatedFull, expectedFull) {
		t.Errorf("Expected %v * %v to be %v but got %v", fullQuarter, p, expectedFull, translatedFull)
	}
}

func TestRotateAroundZ(t *testing.T) {
	halfQuarter := rotationZ(math.Pi / 4)
	fullQuarter := rotationZ(math.Pi / 2)
	p := point(0, 1, 0)

	translatedHalf, err := matrix4x4TupleMultiply(halfQuarter, p)
	if err != nil {
		t.Fatal(err)
	}

	expectedHalf := point(-math.Sqrt(2)/2, math.Sqrt(2)/2, 0)

	if !tupleEqual(translatedHalf, expectedHalf) {
		t.Errorf("Expected %v * %v to be %v but got %v", halfQuarter, p, expectedHalf, translatedHalf)
	}

	translatedFull, err := matrix4x4TupleMultiply(fullQuarter, p)
	if err != nil {
		t.Fatal(err)
	}

	expectedFull := point(-1, 0, 0)

	if !tupleEqual(translatedFull, expectedFull) {
		t.Errorf("Expected %v * %v to be %v but got %v", fullQuarter, p, expectedFull, translatedFull)
	}
}

func TestShearing(t *testing.T) {
	type testCase struct {
		shearing Matrix
		p        Tuple
		expected Tuple
	}

	cases := []testCase{
		{shearing(1, 0, 0, 0, 0, 0), point(2, 3, 4), point(5, 3, 4)},
		{shearing(0, 1, 0, 0, 0, 0), point(2, 3, 4), point(6, 3, 4)},
		{shearing(0, 0, 1, 0, 0, 0), point(2, 3, 4), point(2, 5, 4)},
		{shearing(0, 0, 0, 1, 0, 0), point(2, 3, 4), point(2, 7, 4)},
		{shearing(0, 0, 0, 0, 1, 0), point(2, 3, 4), point(2, 3, 6)},
		{shearing(0, 0, 0, 0, 0, 1), point(2, 3, 4), point(2, 3, 7)},
	}

	for _, c := range cases {
		res, err := matrix4x4TupleMultiply(c.shearing, c.p)
		if err != nil {
			t.Fatal(err)
		}
		if !tupleEqual(res, c.expected) {
			t.Errorf("Expected %v * %v to be %v but got %v", c.shearing, c.p, c.expected, res)
		}
	}
}

func TestMultipleTransformationsInSequence(t *testing.T) {
	p := point(1, 0, 1)
	A := rotationX(math.Pi / 2)
	B := scaling(5, 5, 5)
	C := translation(10, 5, 7)

	p2, err := matrix4x4TupleMultiply(A, p)
	if err != nil {
		t.Fatal(err)
	}
	expected := point(1, -1, 0)

	if !tupleEqual(p2, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", A, p, expected, p2)
	}

	p3, err := matrix4x4TupleMultiply(B, p2)
	if err != nil {
		t.Fatal(err)
	}
	expected = point(5, -5, 0)

	if !tupleEqual(p3, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", B, p2, expected, p3)
	}

	p4, err := matrix4x4TupleMultiply(C, p3)
	if err != nil {
		t.Fatal(err)
	}
	expected = point(15, 0, 7)

	if !tupleEqual(p4, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", C, p3, expected, p4)
	}
}

func TestMultipleTransformationsChained(t *testing.T) {
	p := point(1, 0, 1)
	A := rotationX(math.Pi / 2)
	B := scaling(5, 5, 5)
	C := translation(10, 5, 7)
	expected := point(15, 0, 7)

	ABC1, err := matrix4x4Multiply(B, A)
	if err != nil {
		t.Fatal(err)
	}
	ABC1, err = matrix4x4Multiply(C, ABC1)
	if err != nil {
		t.Fatal(err)
	}
	out, err := matrix4x4TupleMultiply(ABC1, p)
	if err != nil {
		t.Fatal(err)
	}
	if !tupleEqual(out, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", ABC1, p, expected, out)
	}

	ABC2, err := transformation(A, B, C)
	if err != nil {
		t.Fatal(err)
	}
	out, err = matrix4x4TupleMultiply(ABC2, p)
	if err != nil {
		t.Fatal(err)
	}
	if !tupleEqual(out, expected) {
		t.Errorf("Expected %v * %v to be %v but got %v", ABC1, p, expected, out)
	}
}
