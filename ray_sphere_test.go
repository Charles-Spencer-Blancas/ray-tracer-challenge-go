package main

import (
	"reflect"
	"testing"
)

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
		out := rayPosition(r, c.t)

		if !tupleEqual(out, c.expected) {
			t.Errorf("Expected position(%v, %f) to be %v but got %v", r, c.t, c.expected, out)
		}
	}
}

func TestRayIntersectsSphereAtTwoPoints(t *testing.T) {
	r, err := ray(point(0, 0, -5), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}
	s := sphere()
	xs := sphereRayIntersect(s, r)

	if len(xs) != 2 || xs[0].t != 4.0 || xs[1].t != 6.0 {
		t.Errorf("Expected %v to be [4.0, 6.0] but it is not", xs)
	}
}

func TestRayIntersectsSphereAtTangent(t *testing.T) {
	r, err := ray(point(0, 1, -5), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}
	s := sphere()
	xs := sphereRayIntersect(s, r)

	if len(xs) != 2 || xs[0].t != 5.0 || xs[1].t != 5.0 {
		t.Errorf("Expected %v to be [5.0, 5.0] but it is not", xs)
	}
}

func TestRayMissesSphere(t *testing.T) {
	r, err := ray(point(0, 2, -5), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}
	s := sphere()
	xs := sphereRayIntersect(s, r)

	if len(xs) != 0 {
		t.Errorf("Expected %v to be [] but it is not", xs)
	}
}

func TestRayOriginatesInsideSphere(t *testing.T) {
	r, err := ray(point(0, 0, 0), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}
	s := sphere()
	xs := sphereRayIntersect(s, r)

	if len(xs) != 2 || xs[0].t != -1.0 || xs[1].t != 1.0 {
		t.Errorf("Expected %v to be [-1.0, 1.0] but it is not", xs)
	}
}

func TestRayIntersectsIsInFrontOfSphere(t *testing.T) {
	r, err := ray(point(0, 0, 5), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}
	s := sphere()
	xs := sphereRayIntersect(s, r)

	if len(xs) != 2 || xs[0].t != -6.0 || xs[1].t != -4.0 {
		t.Errorf("Expected %v to be [-6.0, -4.0] but it is not", xs)
	}
}

func TestIntersectionEncapsulatesTAndObject(t *testing.T) {
	s := sphere()
	i := Intersection{s, 3.5}
	if !floatEqual(i.t, 3.5) || !reflect.DeepEqual(i.Object, s) {
		t.Errorf("Expected %v but got %v", Intersection{s, 3.5}, i)
	}
}

func TestAggregateIntersections(t *testing.T) {
	s := sphere()
	i1 := Intersection{s, 1}
	i2 := Intersection{s, 2}

	xs := intersections(i1, i2)

	if len(xs) != 2 || !floatEqual(xs[0].t, 1) || !floatEqual(xs[1].t, 2) {
		t.Errorf("Expected %v to be %v but it is not", xs, []Intersection{i1, i2})
	}
}

func TestIntersectSetsObjectOnIntersection(t *testing.T) {
	r, err := ray(point(0, 0, -5), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}
	s := sphere()
	xs := sphereRayIntersect(s, r)
	if len(xs) != 2 || !reflect.DeepEqual(xs[0].Object, s) || !reflect.DeepEqual(xs[1].Object, s) {
		t.Errorf("Object is not set: %v", xs)
	}
}

func TestHitPositiveT(t *testing.T) {
	s := sphere()
	i1 := Intersection{s, 1}
	i2 := Intersection{s, 2}
	xs := intersections(i2, i1)
	i := hit(xs)
	if !reflect.DeepEqual(i1, i) {
		t.Errorf("Expected hit to be %v but it is %v", i1, i)
	}
}

func TestHitPositiveAndNegative(t *testing.T) {
	s := sphere()
	i1 := Intersection{s, -1}
	i2 := Intersection{s, 1}
	xs := intersections(i2, i1)
	i := hit(xs)
	if !reflect.DeepEqual(i2, i) {
		t.Errorf("Expected hit to be %v but it is %v", i2, i)
	}
}

func TestHitNegativeT(t *testing.T) {
	s := sphere()
	i1 := Intersection{s, -1}
	i2 := Intersection{s, -2}
	xs := intersections(i2, i1)
	i := hit(xs)
	if !reflect.DeepEqual(i, (Intersection{})) {
		t.Errorf("Expected %v to be blank", i)
	}
}

func TestHitLowestNonNegative(t *testing.T) {
	s := sphere()
	i1 := Intersection{s, 5}
	i2 := Intersection{s, 7}
	i3 := Intersection{s, -3}
	i4 := Intersection{s, 2}
	xs := intersections(i1, i2, i3, i4)
	i := hit(xs)
	if !reflect.DeepEqual(i, i4) {
		t.Errorf("Expected hit to be %v but it is %v", i4, i)
	}
}

func TestIntersectionsIsSorted(t *testing.T) {
	s := sphere()
	i1 := Intersection{s, 5}
	i2 := Intersection{s, 7}
	i3 := Intersection{s, -3}
	i4 := Intersection{s, 2}
	xs := intersections(i1, i2, i3, i4)
	if !reflect.DeepEqual(xs[0], i3) || !reflect.DeepEqual(xs[1], i4) || !reflect.DeepEqual(xs[2], i1) || !reflect.DeepEqual(xs[3], i2) {
		t.Errorf("Expected %v to be sorted", xs)
	}
}

func TestTranslateRay(t *testing.T) {
	r, err := ray(point(1, 2, 3), vector(0, 1, 0))
	if err != nil {
		t.Fatal(err)
	}
	m := translation(3, 4, 5)
	r2, err := rayMatrixTransform(r, m)
	if err != nil {
		t.Fatal(err)
	}
	expected := Ray{point(4, 6, 8), vector(0, 1, 0)}

	if r2 != expected {
		t.Errorf("Expected %v to be %v", r2, expected)
	}
}

func TestScaleRay(t *testing.T) {
	r, err := ray(point(1, 2, 3), vector(0, 1, 0))
	if err != nil {
		t.Fatal(err)
	}
	m := scaling(2, 3, 4)
	r2, err := rayMatrixTransform(r, m)
	if err != nil {
		t.Fatal(err)
	}
	expected := Ray{point(2, 6, 12), vector(0, 3, 0)}

	if r2 != expected {
		t.Errorf("Expected %v to be %v", r2, expected)
	}
}
