package color

import "github.com/Henelik/tricaster/util"

type Color struct {
	R float64
	G float64
	B float64
}

func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

func (c *Color) Add(o *Color) *Color {
	return &Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

func (c *Color) Sub(o *Color) *Color {
	return &Color{c.R - o.R, c.G - o.G, c.B - o.B}
}

func (c *Color) MultF(n float64) *Color {
	return &Color{c.R * n, c.G * n, c.B * n}
}

func (c *Color) MultCol(o *Color) *Color {
	return &Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

func (c *Color) Equal(o *Color) bool {
	return util.Equal(c.R, o.R) &&
		util.Equal(c.G, o.G) &&
		util.Equal(c.B, o.B)
}
