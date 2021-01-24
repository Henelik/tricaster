package color

type Color struct {
	R float64
	G float64
	B float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}
