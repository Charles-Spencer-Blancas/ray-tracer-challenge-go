package main

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func colorEqual(a Color, b Color) bool {
	return floatEqual(a.Red, b.Red) && floatEqual(a.Green, b.Green) && floatEqual(a.Blue, b.Blue)
}

func colorAdd(a Color, b Color) Color {
	return Color{a.Red + b.Red, a.Green + b.Green, a.Blue + b.Blue}
}

func colorSubtract(a Color, b Color) Color {
	return Color{a.Red - b.Red, a.Green - b.Green, a.Blue - b.Blue}
}

func colorScale(a Color, k float64) Color {
	return Color{a.Red * k, a.Green * k, a.Blue * k}
}

func colorBlend(a Color, b Color) Color {
	return Color{a.Red * b.Red, a.Green * b.Green, a.Blue * b.Blue}
}
