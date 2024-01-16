package main

import (
	"fmt"
	"math"
)

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

func lighting(material Material, light PointLight, point Tuple, eyeV Tuple, normalV Tuple) Color {
	// Blend surface color with light's color
	effectiveColor := colorBlend(material.Color, light.Intensity)

	// Find direction to light source
	lightV := vectorNormalize(tupleSubtract(light.Position, point))

	// Compute ambient contribution
	ambient := colorScale(effectiveColor, material.Ambient)

	// LightDotNormal: Cos of angle between light vector and normal
	// Negative means light on other side of surface, so just ambient, no diffuse and specular
	lightDotNormal := vectorDot(lightV, normalV)
	if lightDotNormal < 0 {
		return ambient
	}

	// Compute diffuse contribution
	diffuse := colorScale(effectiveColor, material.Diffuse*lightDotNormal)

	// ReflectDotEye: Cos of angle between reflection vector and eye vector
	// Negative means light reflects away from eye
	// So no specular, just ambient and diffuse
	reflectV := vectorNormalReflect(tupleNegate(lightV), normalV)
	reflectDotEye := vectorDot(reflectV, eyeV)
	if reflectDotEye <= 0 {
		return colorAdd(ambient, diffuse)
	}

	// Here, compute specular contribution
	factor := math.Pow(reflectDotEye, material.Shininess)
	specular := colorScale(light.Intensity, material.Specular*factor)

	return colorAdd(ambient, colorAdd(diffuse, specular))
}
