package main

import (
	"strings"
	"testing"
)

func TestCreateCanvas(t *testing.T) {
	c := canvas(10, 20)

	if c.Width != 10 {
		t.Errorf("Expected width to be 10 but got %d", c.Width)
	}

	if c.Height != 20 {
		t.Errorf("Expected height to be 20 but got %d", c.Height)
	}

	for i, v := range c.Pixels {
		for j, u := range v {
			if !colorEqual(u, Color{0, 0, 0}) {
				t.Errorf("Expected pixel at %d,%d to be %v but got %v", j, i, Color{0, 0, 0}, u)
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := canvas(10, 20)
	red := Color{1, 0, 0}
	writePixel(c, 2, 3, red)

	if !colorEqual(pixelAt(c, 2, 3), red) {
		t.Errorf("Expected pixel at %d,%d to be %v but got %v", 2, 3, red, c.Pixels[3][2])
	}
}

func TestPPMHeader(t *testing.T) {
	c := canvas(5, 3)
	p := canvasToPPM(c)
	lines := strings.Split(p, "\n")

	if lines[0] != "P3" {
		t.Errorf("Expected line 0 to be %s but got %s", "P3", lines[0])
	}
	if lines[1] != "5 3" {
		t.Errorf("Expected line 1 to be %s but got %s", "5 3", lines[1])
	}
	if lines[2] != "255" {
		t.Errorf("Expected line 2 to be %s but got %s", "255", lines[2])
	}
}

func TestPPMData(t *testing.T) {
	c := canvas(5, 3)
	c1 := Color{1.5, 0, 0}
	c2 := Color{0, 0.5, 0}
	c3 := Color{-0.5, 0, 1}

	writePixel(c, 0, 0, c1)
	writePixel(c, 2, 1, c2)
	writePixel(c, 4, 2, c3)

	p := canvasToPPM(c)
	lines := strings.Split(p, "\n")

	if lines[3] != "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0" {
		t.Errorf("Expected line 3 to be %s but got %s", "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0", lines[3])
	}
	if lines[4] != "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0" {
		t.Errorf("Expected line 4 to be %s but got %s", "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0", lines[4])
	}
	if lines[5] != "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255" {
		t.Errorf("Expected line 5 to be %s but got %s", "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255", lines[5])
	}
}

func Test70MaxLineLength(t *testing.T) {
	c := canvas(10, 2)
	color := Color{1, 0.8, 0.6}

	for i := range c.Pixels {
		for j := range c.Pixels[i] {
			writePixel(c, int64(j), int64(i), color)
		}
	}

	p := canvasToPPM(c)

	lines := strings.Split(p, "\n")

	if lines[3] != "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204" {
		t.Errorf("Expected line 3 to be %s but got %s", "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", lines[3])
	}
	if lines[4] != "153 255 204 153 255 204 153 255 204 153 255 204 153" {
		t.Errorf("Expected line 4 to be %s but got %s", "153 255 204 153 255 204 153 255 204 153 255 204 153", lines[4])
	}
	if lines[5] != "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204" {
		t.Errorf("Expected line 5 to be %s but got %s", "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204", lines[5])
	}
	if lines[6] != "153 255 204 153 255 204 153 255 204 153 255 204 153" {
		t.Errorf("Expected line 6 to be %s but got %s", "153 255 204 153 255 204 153 255 204 153 255 204 153", lines[6])
	}
}

func TestTerminateWithNewline(t *testing.T) {
	c := canvas(5, 3)
	ppm := canvasToPPM(c)
	if ppm[len(ppm)-1] != '\n' {
		t.Errorf("Expected ppm to end with \n but got %b", ppm[len(ppm)-1])
	}
}
