package main

import "fmt"

type PointLight struct {
	Position  Tuple
	Intensity Color
}

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func material() Material {
	return Material{Color{1, 1, 1}, 0.1, 0.9, 0.9, 200.}
}

func pointLight(p Tuple, i Color) (PointLight, error) {
	if !isPoint(p) {
		return PointLight{}, fmt.Errorf("can only make PointLight with point but got %v", p)
	}
	return PointLight{p, i}, nil
}

func sphereNormalAt(s Sphere, p Tuple) (Tuple, error) {
	inv, err := matrixInverse(s.Transform)
	if err != nil {
		return Tuple{}, err
	}
	objectPoint, err := matrix4x4TupleMultiply(inv, p)
	if err != nil {
		return Tuple{}, err
	}
	// Get the normal in object space
	objectNormal := tupleSubtract(objectPoint, point(0, 0, 0))
	// Convert the normal from object to world space
	worldNormal, err := matrix4x4TupleMultiply(matrixTranspose(inv), objectNormal)
	if err != nil {
		return Tuple{}, err
	}
	worldNormal.W = 0
	return vectorNormalize(worldNormal), nil
}

func vectorNormalReflect(in Tuple, normal Tuple) Tuple {
	d := vectorDot(in, normal) * 2
	return tupleSubtract(in, tupleScale(normal, d))
}
