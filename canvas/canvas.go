package canvas

import (
	"errors"

	"github.com/Henelik/tricaster/color"
)

type Canvas struct {
	W   int
	H   int
	Pix []color.Color
}

func NewCanvas(w, h int) (*Canvas, error) {
	if w <= 0 || h <= 0 {
		return nil, errors.New("canvas width and heigh must be 1 or more")
	}
	return &Canvas{
		W:   w,
		H:   h,
		Pix: make([]color.Color, w*h),
	}, nil
}

func (c *Canvas) Get(x, y int) *color.Color {
	return &c.Pix[x+y*c.W]
}

func (c *Canvas) Set(x, y int, col *color.Color) {
	c.Pix[x+y*c.W].R = col.R
	c.Pix[x+y*c.W].G = col.G
	c.Pix[x+y*c.W].B = col.B
}

func (c *Canvas) ToPPM() string {
	s := "PPM\n"

	return s
}
