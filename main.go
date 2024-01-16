package main

import (
	"os"
	"reflect"
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
	canvasSize := int64(300)
	source := point(0, 0, -5)
	wallZ := float64(10)
	wallSize := int64(7)
	pixelSize := float64(wallSize) / float64(canvasSize)
	half := float64(wallSize) / 2.

	c := canvas(canvasSize, canvasSize)

	s := sphere()
	s.Material.Color = Color{0.8, 0.6, 0.6}

	lightPos := point(10, 10, -10)
	lightColor := Color{1, 1, 1}
	light, err := pointLight(lightPos, lightColor)
	if err != nil {
		os.Exit(-1)
	}

	// Loop through canvas pixels, but transform to world coords
	for y := int64(0); y < canvasSize; y++ {
		worldY := half - pixelSize*float64(y)

		for x := int64(0); x < canvasSize; x++ {
			worldX := -half + pixelSize*float64(x)

			pos := point(float64(worldX), float64(worldY), wallZ)

			ray, err := ray(source, vectorNormalize(tupleSubtract(pos, source)))
			if err != nil {
				os.Exit(-1)
			}

			xs, err := sphereRayIntersect(s, ray)
			if err != nil {
				os.Exit(-1)
			}

			h := hit(xs)
			if !reflect.DeepEqual(h, Intersection{}) {
				point := rayPosition(ray, h.t)
				normal, err := sphereNormalAt(h.Object, point)
				if err != nil {
					os.Exit(-2)
				}
				eye := tupleNegate(ray.Direction)
				color := lighting(h.Object.Material, light, point, eye, normal)

				writePixel(c, x, y, color)
			}
		}
	}

	ppm := canvasToPPM(c)
	os.WriteFile("wall.ppm", []byte(ppm), 0644)
}
