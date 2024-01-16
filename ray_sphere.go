package main

import (
	"fmt"
	"math"
	"sort"
)

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

type Sphere struct {
	Transform Matrix
	Origin    Tuple
	Radius    float64
	Material  Material
}

type Intersection struct {
	Object Sphere
	t      float64
}

func ray(origin Tuple, direction Tuple) (Ray, error) {
	if !isPoint(origin) {
		return Ray{}, fmt.Errorf("origin %v must be a point but it is not", origin)
	}
	if !isVector(direction) {
		return Ray{}, fmt.Errorf("direction %v must be a vector but it is not", direction)
	}

	return Ray{origin, direction}, nil
}

func rayPosition(ray Ray, t float64) Tuple {
	return tupleAdd(ray.Origin, tupleScale(ray.Direction, t))
}

func sphere() Sphere {
	return Sphere{matrixConstructIdentity(4), point(0, 0, 0), 1., material()}
}

func sphereRayIntersect(s Sphere, r Ray) ([]Intersection, error) {
	t, err := matrixInverse(s.Transform)
	if err != nil {
		return []Intersection{}, err
	}
	r, err = rayMatrixTransform(r, t)
	if err != nil {
		return []Intersection{}, err
	}
	sphereToRay := tupleSubtract(r.Origin, s.Origin)

	a := vectorDot(r.Direction, r.Direction)
	b := vectorDot(r.Direction, sphereToRay) * 2
	c := vectorDot(sphereToRay, sphereToRay) - 1

	disc := b*b - 4*a*c

	if disc < 0 {
		return []Intersection{}, nil
	}

	t1 := (-b - math.Sqrt(disc)) / (2 * a)
	t2 := (-b + math.Sqrt(disc)) / (2 * a)

	return []Intersection{{s, t1}, {s, t2}}, nil
}

// Returns sorted intersections
func intersections(ts ...Intersection) []Intersection {
	arr := make([]Intersection, len(ts))
	copy(arr, ts)
	return sortIntersections(arr)
}

func sortIntersections(ts []Intersection) []Intersection {
	sort.Slice(ts, func(i, j int) bool { return ts[i].t < ts[j].t })

	return ts
}

// Relies on intersections being sorted
func hit(is []Intersection) Intersection {
	for _, i := range is {
		if i.t < 0.0 {
			continue
		}

		return i
	}

	return Intersection{}
}

func rayMatrixTransform(r Ray, m Matrix) (Ray, error) {
	origin, err := matrix4x4TupleMultiply(m, r.Origin)
	if err != nil {
		return Ray{}, err
	}
	dir, err := matrix4x4TupleMultiply(m, r.Direction)
	if err != nil {
		return Ray{}, err
	}

	return Ray{origin, dir}, err
}
