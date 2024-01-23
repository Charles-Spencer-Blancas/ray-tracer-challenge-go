package main

import (
	"math"
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

type Camera struct {
	Transform   Matrix
	HSize       int64
	VSize       int64
	FieldOfView float64
	PixelSize   float64
	HalfWidth   float64
	HalfHeight  float64
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

func viewTransform(from Tuple, to Tuple, up Tuple) (Matrix, error) {
	forward := vectorNormalize(tupleSubtract(to, from))
	upN := vectorNormalize(up)
	left := vectorCross(forward, upN)
	trueUp := vectorCross(left, forward)
	orientation := matrixConstruct([][]float64{
		{left.X, left.Y, left.Z, 0.},
		{trueUp.X, trueUp.Y, trueUp.Z, 0},
		{-forward.X, -forward.Y, -forward.Z, 0},
		{0, 0, 0, 1},
	})
	tr := translation(-from.X, -from.Y, -from.Z)
	return matrix4x4Multiply(orientation, tr)
}

func camera(hSize int64, vSize int64, fov float64) Camera {
	halfView := math.Tan(fov / 2.)
	aspect := float64(hSize) / float64(vSize)

	var hw, hh float64
	if aspect >= 1 {
		hw = halfView
		hh = halfView / aspect
	} else {
		hw = halfView * aspect
		hh = halfView
	}

	pixelSize := (hw * 2) / float64(hSize)

	return Camera{matrixConstructIdentity(4), hSize, vSize, fov, pixelSize, hw, hh}
}

func rayForPixel(camera Camera, px int64, py int64) (Ray, error) {
	// Offset from edge of canvas to center of pixel
	xOffset := (float64(px) + 0.5) * camera.PixelSize
	yOffset := (float64(py) + 0.5) * camera.PixelSize

	// Untransformed coords of pixel in world space
	// Camera looks toward -z, so +x is left
	worldX := camera.HalfWidth - xOffset
	worldY := camera.HalfHeight - yOffset

	// Transform canvas point and origin with camera matrix
	// Compute ray's direction vector
	// Note that canvas at z=-1
	inv, err := matrixInverse(camera.Transform)
	if err != nil {
		return Ray{}, err
	}
	pixel, err := matrix4x4TupleMultiply(inv, point(worldX, worldY, -1))
	if err != nil {
		return Ray{}, err
	}
	origin, err := matrix4x4TupleMultiply(inv, point(0, 0, 0))
	if err != nil {
		return Ray{}, err
	}
	direction := vectorNormalize(tupleSubtract(pixel, origin))
	return Ray{origin, direction}, err
}

func render(camera Camera, world World) (Canvas, error) {
	image := canvas(camera.HSize, camera.VSize)

	for y := int64(0); y < camera.VSize; y++ {
		for x := int64(0); x < camera.HSize; x++ {
			ray, err := rayForPixel(camera, x, y)
			if err != nil {
				return Canvas{}, err
			}
			color, err := colorAt(world, ray)
			if err != nil {
				return Canvas{}, err
			}
			writePixel(image, x, y, color)
		}
	}

	return image, nil
}
