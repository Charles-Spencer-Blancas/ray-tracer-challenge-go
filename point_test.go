package main

import "testing"

func TestTupleW1IsPoint(t *testing.T) {
	a := Tuple{4.3, -4.2, 3.1, 1.0}

	if a.X != 4.3 {
		t.Errorf("Expected x to be 4.3 but got %f", a.X)
	}
	if a.Y != -4.2 {
		t.Errorf("Expected y to be -4.2 but got %f", a.Y)
	}
	if a.Z != 3.1 {
		t.Errorf("Expected z to be 3.1 but got %f", a.Z)
	}
	if a.W != 1.0 {
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

	if a.X != 4.3 {
		t.Errorf("Expected x to be 4.3 but got %f", a.X)
	}
	if a.Y != -4.2 {
		t.Errorf("Expected y to be -4.2 but got %f", a.Y)
	}
	if a.Z != 3.1 {
		t.Errorf("Expected z to be 3.1 but got %f", a.Z)
	}
	if a.W != 0.0 {
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

	if p != pp {
		t.Errorf("expected point to be %v but got %v", pp, p)
	}
}

func TestVectorCreatesTupleW0(t *testing.T) {
	v := vector(4, -4, 3)
	vv := Tuple{4, -4, 3, 0.0}

	if v != vv {
		t.Errorf("expected point to be %v but got %v", vv, v)
	}
}
