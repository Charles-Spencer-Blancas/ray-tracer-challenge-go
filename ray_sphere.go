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
	Origin Tuple
	Radius float64
}

type Intersection struct {
	t      float64
	Object Sphere
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

func position(ray Ray, t float64) Tuple {
	return tupleAdd(ray.Origin, tupleScale(ray.Direction, t))
}

func sphere() Sphere {
	return Sphere{point(0, 0, 0), 1.}
}

func intersect(s Sphere, r Ray) []Intersection {
	sphereToRay := tupleSubtract(r.Origin, s.Origin)

	a := vectorDot(r.Direction, r.Direction)
	b := vectorDot(r.Direction, sphereToRay) * 2
	c := vectorDot(sphereToRay, sphereToRay) - 1

	disc := b*b - 4*a*c

	if disc < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(disc)) / (2 * a)
	t2 := (-b + math.Sqrt(disc)) / (2 * a)

	return []Intersection{{t1, s}, {t2, s}}
}

// Returns sorted intersections
func intersections(ts ...Intersection) []Intersection {
	arr := make([]Intersection, len(ts))
	copy(arr, ts)
	sort.Slice(arr, func(i, j int) bool { return arr[i].t < arr[j].t })
	return arr
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
