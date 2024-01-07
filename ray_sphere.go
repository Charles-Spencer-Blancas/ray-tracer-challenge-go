package main

import "fmt"

type Ray struct {
	Origin    Tuple
	Direction Tuple
}

type Sphere struct {
	Origin Tuple
	Radius float64
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
