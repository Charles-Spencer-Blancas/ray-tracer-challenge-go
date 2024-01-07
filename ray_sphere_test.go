package main

import "testing"

func TestCreateRay(t *testing.T) {
	o := point(1, 2, 3)
	d := vector(4, 5, 6)
	r, err := ray(o, d)
	if err != nil {
		t.Fatal(err)
	}

	if !tupleEqual(r.Origin, o) {
		t.Errorf("Expected %v to equal %v", r.Origin, o)
	}
	if !tupleEqual(r.Direction, d) {
		t.Errorf("Expected %v to equal %v", r.Direction, d)
	}
}

func TestComputePointFromDistance(t *testing.T) {
	r, err := ray(point(2, 3, 4), vector(1, 0, 0))
	if err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		t        float64
		expected Tuple
	}

	cases := []testCase{
		{0, point(2, 3, 4)},
		{1, point(3, 3, 4)},
		{-1, point(1, 3, 4)},
		{2.5, point(4.5, 3, 4)},
	}

	for _, c := range cases {
		out := position(r, c.t)

		if !tupleEqual(out, c.expected) {
			t.Errorf("Expected position(%v, %f) to be %v but got %v", r, c.t, c.expected, out)
		}
	}
}
