package main

import (
	"math"
	"testing"
)

func TestTupleW1IsPoint(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 1.0}

	if !floatEqual(a.X, 4.3) {
		t.Errorf("Expected x to be 4.3 but got %f", a.X)
	}
	if !floatEqual(a.Y, -4.2) {
		t.Errorf("Expected y to be -4.2 but got %f", a.Y)
	}
	if !floatEqual(a.Z, 3.1) {
		t.Errorf("Expected z to be 3.1 but got %f", a.Z)
	}
	if !floatEqual(a.W, 1.0) {
		t.Errorf("Expected w to be 1.0 but got %f", a.W)
	}
	if !isPoint(a) {
		t.Errorf("Expected a to be a point but it is not")
	}
	if isVector(a) {
		t.Errorf("Expected a not to be a vector but it is")
	}
}

func TestTupleW0IsVector(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 0.0}

	if !floatEqual(a.X, 4.3) {
		t.Errorf("Expected x to be 4.3 but got %f", a.X)
	}
	if !floatEqual(a.Y, -4.2) {
		t.Errorf("Expected y to be -4.2 but got %f", a.Y)
	}
	if !floatEqual(a.Z, 3.1) {
		t.Errorf("Expected z to be 3.1 but got %f", a.Z)
	}
	if !floatEqual(a.W, 0.0) {
		t.Errorf("Expected w to be 0.0 but got %f", a.W)
	}
	if isPoint(a) {
		t.Errorf("Expected a not to be a point but it is")
	}
	if !isVector(a) {
		t.Errorf("Expected a to be a vector but it is not")
	}
}

func TestPointCreatesTupleW1(t *testing.T) {
	p := point(4, -4, 3)
	pp := Tuple{4, -4, 3, 1.0}

	if !tupleEqual(p, pp) {
		t.Errorf("expected point to be %v but got %v", pp, p)
	}
}

func TestVectorCreatesTupleW0(t *testing.T) {
	v := vector(4, -4, 3)
	vv := Tuple{4, -4, 3, 0.0}

	if !tupleEqual(v, vv) {
		t.Errorf("expected point to be %v but got %v", vv, v)
	}
}

func TestAddTwoTuples(t *testing.T) {
	a1 := Tuple{3, -2, 5, 1}
	a2 := Tuple{-2, 3, 1, 0}
	added := tupleAdd(a1, a2)
	expected := Tuple{1, 1, 6, 1}

	if !tupleEqual(added, expected) {
		t.Errorf("Expected a1 + a2 to be %v but got %v", expected, added)
	}
}

func TestSubtractTwoPoints(t *testing.T) {
	p1 := point(3, 2, 1)
	p2 := point(5, 6, 7)
	subbed := tupleSubtract(p1, p2)
	expected := vector(-2, -4, -6)

	if !tupleEqual(subbed, expected) {
		t.Errorf("Expected p1 - p2 to be %v but got %v", expected, subbed)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := point(3, 2, 1)
	v := vector(5, 6, 7)
	subbed := tupleSubtract(p, v)
	expected := point(-2, -4, -6)

	if !tupleEqual(subbed, expected) {
		t.Errorf("Expected p - v to be %v but got %v", expected, subbed)
	}
}

func TestSubtractTwoVectors(t *testing.T) {
	v1 := vector(3, 2, 1)
	v2 := vector(5, 6, 7)
	subbed := tupleSubtract(v1, v2)
	expected := vector(-2, -4, -6)

	if !tupleEqual(subbed, expected) {
		t.Errorf("Expected v1 - v2 to be %v but got %v", expected, subbed)
	}
}

func TestSubtractVectorFromZeroVector(t *testing.T) {
	zero := vector(0, 0, 0)
	v := vector(1, -2, 3)
	subbed := tupleSubtract(zero, v)
	expected := vector(-1, 2, -3)

	if !tupleEqual(subbed, expected) {
		t.Errorf("Expected zero - v to be %v but got %v", expected, subbed)
	}
}

func TestNegateTuple(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	aNeg := tupleNegate(a)
	expected := Tuple{-1, 2, -3, 4}

	if !tupleEqual(aNeg, expected) {
		t.Errorf("Expected -a to be %v but got %v", expected, aNeg)
	}
}

func TestMultiplyTupleByScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	ak := tupleMultiply(a, 3.5)
	expected := Tuple{3.5, -7, 10.5, -14}

	if !tupleEqual(ak, expected) {
		t.Errorf("Expected a * 3.5 to be %v but got %v", expected, ak)
	}
}

func TestMultiplyTupleByFraction(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	ak := tupleMultiply(a, 0.5)
	expected := Tuple{0.5, -1, 1.5, -2}

	if !tupleEqual(ak, expected) {
		t.Errorf("Expected a * 0.5 to be %v but got %v", expected, ak)
	}
}

func TestDivideTupleByScalar(t *testing.T) {
	a := Tuple{1, -2, 3, -4}
	ak := tupleDivide(a, 2)
	expected := Tuple{0.5, -1, 1.5, -2}

	if !tupleEqual(ak, expected) {
		t.Errorf("Expected a / 2 to be %v but got %v", expected, ak)
	}
}

func TestMagnitude(t *testing.T) {
	vs := []Tuple{vector(1, 0, 0), vector(0, 1, 0), vector(0, 0, 1), vector(1, 2, 3), vector(-1, -2, -3)}
	expecteds := []float64{1, 1, 1, math.Sqrt(14), math.Sqrt(14)}

	if len(vs) != len(expecteds) {
		t.Fatalf("Do not have the same number of vectors and expected values. Cannot continue test")
	}

	for i := 0; i < len(vs); i++ {
		if !floatEqual(vectorMagnitude(vs[i]), expecteds[i]) {
			t.Errorf("Expected %v to have magnitude %f but got %f", vs[i], expecteds[i], vectorMagnitude(vs[i]))
		}
	}
}
