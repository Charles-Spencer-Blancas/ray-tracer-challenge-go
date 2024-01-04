package main

import (
	"fmt"
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
	start := point(0, 1, 0)
	vel := tupleScale(vectorNormalize(vector(1, 1.8, 0)), 11.25)
	p := Projectile{start, vel}
	e := Environment{vector(0, -0.1, 0), vector(-0.01, 0, 0)}
	t := 0

	color := Color{1, 0, 0}
	c := canvas(900, 550)
	for p.Position.Y > 0.0 {
		fmt.Printf("tick: %d, pos: %v\n", t, p.Position)
		p = tick(e, p)

		writePixel(c, int64(math.Round(p.Position.X)), int64(550-math.Round(p.Position.Y)), color)
	}

	ppm := canvasToPPM(c)
	os.WriteFile("example.ppm", []byte(ppm), 0644)
}
