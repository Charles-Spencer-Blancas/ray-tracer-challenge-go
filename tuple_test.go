package main

import "testing"

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
		t.Errorf("Expected a1 + a2 to be %v but got %v", expected, subbed)
	}
}
