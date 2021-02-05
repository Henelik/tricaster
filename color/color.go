package color

import "github.com/Henelik/tricaster/util"

var (
	Red     = &Color{1, 0, 0}
	Green   = &Color{0, 1, 0}
	Blue    = &Color{0, 0, 1}
	Cyan    = &Color{0, 1, 1}
	Magenta = &Color{1, 0, 1}
	Yellow  = &Color{1, 1, 0}
	White   = &Color{1, 1, 1}
	Black   = &Color{0, 0, 0}
)

type Color struct {
	R float32
	G float32
	B float32
}

func NewColor(r, g, b float32) *Color {
	return &Color{r, g, b}
}

func Grey(v float32) *Color {
	return &Color{v, v, v}
}

func (c *Color) Add(o *Color) *Color {
	return &Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

func (c *Color) Sub(o *Color) *Color {
	return &Color{c.R - o.R, c.G - o.G, c.B - o.B}
}

func (c *Color) MultF(n float32) *Color {
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

func Avg(cs []*Color) *Color {
	avg := cs[0]
	for i := 1; i < len(cs); i++ {
		avg = avg.Add(cs[i])
	}
	return avg.MultF(1.0 / float32(len(cs)))
}

func (c *Color) Lerp(o *Color, factor float32) *Color {
	return &Color{
		util.Lerp(c.R, o.R, factor),
		util.Lerp(c.G, o.G, factor),
		util.Lerp(c.B, o.B, factor),
	}
}
