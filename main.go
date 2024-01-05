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

func writePixel3x3(c Canvas, x int64, y int64, col Color) {
	writePixel(c, x, y, col)
	writePixel(c, x-1, y, col)
	writePixel(c, x+1, y, col)
	writePixel(c, x, y-1, col)
	writePixel(c, x-1, y-1, col)
	writePixel(c, x+1, y-1, col)
	writePixel(c, x, y+1, col)
	writePixel(c, x-1, y+1, col)
	writePixel(c, x+1, y+1, col)
}

func writePixel9x9(c Canvas, x int64, y int64, col Color) {
	writePixel3x3(c, x, y, col)
	writePixel3x3(c, x-3, y, col)
	writePixel3x3(c, x+3, y, col)
	writePixel3x3(c, x, y-3, col)
	writePixel3x3(c, x-3, y-3, col)
	writePixel3x3(c, x+3, y-3, col)
	writePixel3x3(c, x, y+3, col)
	writePixel3x3(c, x-3, y+3, col)
	writePixel3x3(c, x+3, y+3, col)
}

func main() {
	color := Color{1, 0, 0}
	c := canvas(1000, 1000)
	radius := 400.

	origin := point(0, 0, 0)
	moveOriginToTwelve := translation(0, -radius, 0)
	moveOriginToCenter := translation(499, 499, 0)

	for i := 0; i < 12; i++ {
		t, err := transformation(
			moveOriginToTwelve,
			rotation_z(float64(i)*math.Pi/6.),
			moveOriginToCenter)
		if err != nil {
			os.Exit(1)
		}
		p, err := matrix4x4TupleMultiply(t, origin)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%d: %v\n", i, p)
		writePixel9x9(c, int64(p.X), int64(p.Y), color)
	}

	ppm := canvasToPPM(c)
	os.WriteFile("clock.ppm", []byte(ppm), 0644)
}
