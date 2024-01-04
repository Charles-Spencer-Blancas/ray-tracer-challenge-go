package main

type Canvas struct {
	Pixels [][]Color
	Width  int64
	Height int64
}

func canvas(w int64, h int64) Canvas {
	p := make([][]Color, h)
	for i := range p {
		p[i] = make([]Color, w)
		for j := range p[i] {
			p[i][j] = Color{0, 0, 0}
		}
	}

	return Canvas{p, w, h}
}

func pixelAt(c Canvas, x int64, y int64) Color {
	return c.Pixels[y][x]
}

func writePixel(c Canvas, x int64, y int64, color Color) {
	c.Pixels[y][x] = color
}
