package main

import (
	"reflect"
	"testing"
)

func TestCreateWorld(t *testing.T) {
	w := World{}

	if w.Light != (PointLight{}) {
		t.Errorf("Expected there to be no light source but there is %v", w.Light)
	}

	if len(w.Objects) != 0 {
		t.Errorf("Expected there to be no objects but there are %v", w.Objects)
	}
}

func TestDefaultWorld(t *testing.T) {
	l, err := pointLight(point(-10, 10, -10), Color{1, 1, 1})
	if err != nil {
		t.Fatal(err)
	}
	s1 := sphere()
	m := material()
	m.Color = Color{0.8, 1.0, 0.6}
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.Material = m

	s2 := sphere()
	s2.Transform = scaling(0.5, 0.5, 0.5)

	w, err := defaultWorld()
	if err != nil {
		t.Fatal(err)
	}

	expected := World{[]Sphere{s1, s2}, l}

	if !reflect.DeepEqual(w.Light, l) ||
		!reflect.DeepEqual(w.Objects[0], s1) ||
		!reflect.DeepEqual(w.Objects[1], s2) {
		t.Errorf("Error with default world %v is supposed to be %v", w, expected)
	}
}

func TestIntersectWorldWithRay(t *testing.T) {
	w, err := defaultWorld()
	if err != nil {
		t.Fatal(err)
	}
	r, err := ray(point(0, 0, -5), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}

	xs, err := worldRayIntersect(w, r)
	if err != nil {
		t.Fatal(err)
	}
	if len(xs) != 4 ||
		!floatEqual(xs[0].t, 4) ||
		!floatEqual(xs[1].t, 4.5) ||
		!floatEqual(xs[2].t, 5.5) ||
		!floatEqual(xs[3].t, 6) {
		t.Errorf("Expected ts to be [4,4.5,5.5,6] but intersection array is %v", xs)
	}
}

func TestPrecomputeStateOfIntersection(t *testing.T) {
	r, err := ray(point(0, 0, -5), vector(0, 0, 1))
	if err != nil {
		t.Fatal(err)
	}

	shape := sphere()
	i := Intersection{shape, 4}
	comps, err := prepareComputations(i, r)
	if err != nil {
		t.Fatal(err)
	}

	expected := Computation{i.Object, i.t, point(0, 0, -1), vector(0, 0, -1), vector(0, 0, -1)}

	if !floatEqual(comps.t, expected.t) ||
		!reflect.DeepEqual(comps.Object, expected.Object) ||
		!tupleEqual(comps.Point, expected.Point) ||
		!tupleEqual(comps.EyeV, expected.EyeV) ||
		!tupleEqual(comps.NormalV, expected.NormalV) {
		t.Errorf("Expected %v to equal %v", comps, expected)
	}
}
