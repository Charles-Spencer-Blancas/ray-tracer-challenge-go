package main

import "testing"

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
