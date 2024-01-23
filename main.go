package main

import (
	"math"
	"os"
)

type Projectile struct {
	Position Tuple // point
	Velocity Tuple // vector
}

type Environment struct {
	gravity Tuple // vector
	wind    Tuple // vector
}

func tick(e Environment, p Projectile) Projectile {
	pos := tupleAdd(p.Position, p.Velocity)
	vel := tupleAdd(tupleAdd(p.Velocity, e.gravity), e.wind)

	return Projectile{pos, vel}
}

func main() {
	floor := sphere()
	floor.Transform = scaling(10, 0.01, 10)
	floor.Material = material()
	floor.Material.Color = Color{1, 0.9, 0.9}
	floor.Material.Specular = 0
	floor.Material.Diffuse = 0.1

	leftWall := sphere()

	lt, err := transformation(
		scaling(10, 0.01, 10),
		rotationX(math.Pi/2.),
		rotationY(-math.Pi/4.),
		translation(0, 0, 5),
	)
	if err != nil {
		os.Exit(-1)
	}
	leftWall.Transform = lt
	leftWall.Material = floor.Material

	rightWall := sphere()
	rt, err := transformation(
		scaling(10, 0.01, 10),
		rotationX(math.Pi/2.),
		rotationY(math.Pi/4.),
		translation(0, 0, 5),
	)
	if err != nil {
		os.Exit(-1)
	}
	rightWall.Transform = rt
	rightWall.Material = floor.Material

	middle := sphere()
	t, err := transformation(translation(-0.5, 1, 0.5), scaling(1, 1.7, 1))
	if err != nil {
		os.Exit(-1)
	}
	middle.Transform = t
	middle.Material = material()
	middle.Material.Color = Color{0.1, 1, 0.5}
	middle.Material.Diffuse = 0.7
	middle.Material.Specular = 0.3

	right := sphere()
	t, err = transformation(scaling(0.5, 0.5, 0.5), translation(1.5, 0.5, -0.5))
	if err != nil {
		os.Exit(-1)
	}
	right.Transform = t
	right.Material.Color = Color{0.5, 1, 0.1}
	right.Material.Diffuse = 0.7
	right.Material.Specular = 0.3

	left := sphere()
	t, err = transformation(scaling(0.33, 0.33, 0.33), translation(-1.5, 0.33, -0.75), shearing(-0.3, 0.2, 0, 0, 0.5, 0))
	if err != nil {
		os.Exit(-1)
	}
	left.Transform = t
	left.Material.Color = Color{1, 0.8, 0.1}
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	float := sphere()
	t, err = transformation(translation(2, 2.1, -0.5), shearing(0.5, 0, 0, 0, 0, 0))
	if err != nil {
		os.Exit(-1)
	}
	float.Transform = t
	left.Material.Color = Color{0.8, 0.6, 0.6}
	left.Material.Diffuse = 0.7
	left.Material.Specular = 0.3

	world := World{}
	world.Objects = []Sphere{leftWall, rightWall, floor, left, right, middle, float}
	l1, err := pointLight(point(-10, 10, -10), Color{1, 0.2, 0.3})
	if err != nil {
		os.Exit(-1)
	}
	l2, err := pointLight(point(10, 10, 10), Color{0.2, 0.8, 1})
	if err != nil {
		os.Exit(-1)
	}
	l3, err := pointLight(point(0, 0, 0), Color{0.5, 0.5, 0.5})
	if err != nil {
		os.Exit(-1)
	}
	world.Lights = []PointLight{l1, l2, l3}
	camera := camera(500, 500, math.Pi/3.)
	vt, err := viewTransform(point(0, 1.5, -5), point(0, 1, 0), vector(0, 1, 0))
	camera.Transform = vt
	if err != nil {
		os.Exit(-1)
	}
	canvas, err := render(camera, world)
	if err != nil {
		os.Exit(-1)
	}
	ppm := canvasToPPM(canvas)

	os.WriteFile("camera.ppm", []byte(ppm), 0644)
}
