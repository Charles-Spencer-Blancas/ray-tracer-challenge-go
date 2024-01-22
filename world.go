package main

import (
	"reflect"
)

type World struct {
	Objects []Sphere
	Lights  []PointLight
}

type Computation struct {
	Object   Sphere
	t        float64
	Point    Tuple
	EyeV     Tuple
	NormalV  Tuple
	IsInside bool
}

func defaultWorld() (World, error) {
	l, err := pointLight(point(-10, 10, -10), Color{1, 1, 1})
	if err != nil {
		return World{}, err
	}
	ls := []PointLight{l}
	s1 := sphere()
	m := material()
	m.Color = Color{0.8, 1.0, 0.6}
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.Material = m

	s2 := sphere()
	s2.Transform = scaling(0.5, 0.5, 0.5)

	return World{[]Sphere{s1, s2}, ls}, nil
}

func worldRayIntersect(w World, r Ray) ([]Intersection, error) {
	intersections := []Intersection{}
	for _, s := range w.Objects {
		is, err := sphereRayIntersect(s, r)
		if err != nil {
			return []Intersection{}, err
		}
		intersections = append(intersections, is...)
	}

	return sortIntersections(intersections), nil
}

func prepareComputations(i Intersection, r Ray) (Computation, error) {
	p := rayPosition(r, i.t)
	n, err := sphereNormalAt(i.Object, p)
	if err != nil {
		return Computation{}, nil
	}
	isInside := false
	eye := tupleNegate(r.Direction)
	if vectorDot(n, eye) < 0 {
		isInside = true
		n = tupleNegate(n)
	}

	return Computation{i.Object, i.t, p, eye, n, isInside}, nil
}

func shadeHit(world World, comps Computation) Color {
	color := Color{0, 0, 0}
	for _, l := range world.Lights {
		color = colorAdd(color,
			lighting(comps.Object.Material,
				l,
				comps.Point,
				comps.EyeV,
				comps.NormalV))
	}
	return color
}

func colorAt(w World, r Ray) (Color, error) {
	is, err := worldRayIntersect(w, r)
	if err != nil {
		return Color{}, err
	}
	h := hit(is)
	if reflect.ValueOf(h).IsZero() {
		return Color{0, 0, 0}, nil
	}
	comps, err := prepareComputations(h, r)
	if err != nil {
		return Color{}, err
	}
	return shadeHit(w, comps), nil
}
