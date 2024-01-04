package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	COLOR_MIN = 0
	COLOR_MAX = 255
)

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

func scaleColorDimension(f float64) int64 {
	if f <= 0 {
		return COLOR_MIN
	}
	if f >= 1 {
		return COLOR_MAX
	}

	return int64(math.Round((COLOR_MAX - COLOR_MIN) * f))
}

func canvasToPPM(c Canvas) string {
	b := strings.Builder{}
	b.WriteString("P3\n")
	b.WriteString(fmt.Sprintf("%d %d\n", c.Width, c.Height))
	b.WriteString(fmt.Sprintf("%d\n", COLOR_MAX))

	for _, v := range c.Pixels {
		line := ""
		for _, u := range v {
			red := fmt.Sprintf("%d ", scaleColorDimension(u.Red))
			if len(line)+len(red) > 70 {
				line = strings.TrimRight(line, " ")
				line += "\n"
				b.WriteString(line)
				line = ""
			}
			line += red

			green := fmt.Sprintf("%d ", scaleColorDimension(u.Green))
			if len(line)+len(green) > 70 {
				line = strings.TrimRight(line, " ")
				line += "\n"
				b.WriteString(line)
				line = ""
			}
			line += green

			blue := fmt.Sprintf("%d ", scaleColorDimension(u.Blue))
			if len(line)+len(blue) > 70 {
				line = strings.TrimRight(line, " ")
				line += "\n"
				b.WriteString(line)
				line = ""
			}
			line += blue

		}
		b.WriteString(strings.TrimRight(line, " "))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	return b.String()
}
