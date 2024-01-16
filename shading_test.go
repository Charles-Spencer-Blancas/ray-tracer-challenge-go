package main

import (
	"math"
	"testing"
)

func TestNormalSphere(t *testing.T) {
	type testCase struct {
		point  Tuple
		normal Tuple
	}

	cases := []testCase{
		{point(1, 0, 0), vector(1, 0, 0)},
		{point(0, 1, 0), vector(0, 1, 0)},
		{point(0, 0, 1), vector(0, 0, 1)},
		{point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3), vector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)},
	}

	for _, v := range cases {
		s := sphere()
		n, err := sphereNormalAt(s, v.point)
		if err != nil {
			t.Fatal(err)
		}
		if !tupleEqual(v.normal, n) {
			t.Errorf("Expected normal at %v to be %v but got %v", v.point, v.normal, n)
		}
	}
}

func TestNormalIsNormalized(t *testing.T) {
	s := sphere()
	n, err := sphereNormalAt(s, point(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if err != nil {
		t.Fatal(err)
	}
	expected := vectorNormalize(n)

	if !tupleEqual(n, expected) {
		t.Errorf("Expected %v to be %v but it is not", n, expected)
	}
}

func TestNormalOnTranslatedSphere(t *testing.T) {
	s := sphere()
	s.Transform = translation(0, 1, 0)
	n, err := sphereNormalAt(s, point(0, 1.70711, -0.70711))
	if err != nil {
		t.Fatal(err)
	}
	expected := vector(0, 0.70711, -0.70711)
	if !tupleEqual(n, expected) {
		t.Errorf("Expected %v to be %v but it is not", n, expected)
	}
}

func TestNormalOnTransformedSphere(t *testing.T) {
	s := sphere()
	ts, err := matrix4x4Multiply(scaling(1, 0.5, 1), rotationZ(math.Pi/5.))
	if err != nil {
		t.Fatal(err)
	}
	s.Transform = ts
	n, err := sphereNormalAt(s, point(0, math.Sqrt2/2, -math.Sqrt2/2))
	if err != nil {
		t.Fatal(err)
	}
	expected := vector(0, 0.97014, -0.24254)
	if !tupleEqual(n, expected) {
		t.Errorf("Expected %v to be %v but it is not", n, expected)
	}
}

func TestReflectVector45Deg(t *testing.T) {
	v := vector(1, -1, 0)
	n := vector(0, 1, 0)
	r := vectorNormalReflect(v, n)
	e := vector(1, 1, 0)

	if !tupleEqual(r, e) {
		t.Errorf("Expected %v to be %v", r, e)
	}
}

func TestReflectVectorSlantedSurface(t *testing.T) {
	v := vector(0, -1, 0)
	n := vector(math.Sqrt2/2, math.Sqrt2/2, 0)
	r := vectorNormalReflect(v, n)
	e := vector(1, 0, 0)

	if !tupleEqual(r, e) {
		t.Errorf("Expected %v to be %v", r, e)
	}
}

func TestPointLightHasPositionAndIntensity(t *testing.T) {
	i := Color{1, 1, 1}
	p := point(0, 0, 0)
	l, err := pointLight(p, i)
	if err != nil {
		t.Fatal(err)
	}
	if !colorEqual(l.Intensity, i) || !tupleEqual(l.Position, p) {
		t.Errorf("PointLight was not set, got %v", l)
	}
}

func TestDefaultMaterial(t *testing.T) {
	m := material()
	if !colorEqual(m.Color, Color{1, 1, 1}) ||
		!floatEqual(m.Ambient, 0.1) ||
		!floatEqual(m.Diffuse, 0.9) ||
		!floatEqual(m.Specular, 0.9) ||
		!floatEqual(m.Shininess, 200.) {
		t.Errorf("Default material is incorrect %v", m)
	}
}
