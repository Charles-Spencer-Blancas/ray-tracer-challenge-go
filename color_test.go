package main

import "testing"

func TestColorRGB(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}

	if !floatEqual(c.Red, -0.5) {
		t.Errorf("Expected red to be -0.5 but got %f", c.Red)
	}
	if !floatEqual(c.Green, 0.4) {
		t.Errorf("Expected green to be 0.4 but got %f", c.Green)
	}
	if !floatEqual(c.Blue, 1.7) {
		t.Errorf("Expected blue to be 1.7 but got %f", c.Blue)
	}
}

func TestAddColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	expected := Color{1.6, 0.7, 1.0}

	if !colorEqual(colorAdd(c1, c2), expected) {
		t.Errorf("Expected c1 + c2 to be %v but got %v", expected, colorAdd(c1, c2))
	}
}

func TestSubtractColors(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	expected := Color{0.2, 0.5, 0.5}

	if !colorEqual(colorSubtract(c1, c2), expected) {
		t.Errorf("Expected c1 - c2 to be %v but got %v", expected, colorSubtract(c1, c2))
	}
}

func TestScaleColor(t *testing.T) {
	c := Color{0.2, 0.3, 0.4}
	expected := Color{0.4, 0.6, 0.8}
	if !colorEqual(colorScale(c, 2), expected) {
		t.Errorf("Expected c * 2 to be %v but got %v", expected, colorScale(c, 2))
	}
}

func TestBlendColors(t *testing.T) {
	c1 := Color{1, 0.2, 0.4}
	c2 := Color{0.9, 1, 0.1}
	expected := Color{0.9, 0.2, 0.04}

	if !colorEqual(colorBlend(c1, c2), expected) {
		t.Errorf("Expected c1 * c2 to be %v but got %v", expected, colorBlend(c1, c2))
	}
}
