package canvas

import (
	"errors"
	"image"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/util"
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

func (c *Canvas) ToImage() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, c.W, c.H))

	for x := 0; x < c.W; x++ {
		for y := 0; y < c.H; y++ {
			r := uint8(util.Clamp(c.Pix[x+y*c.W].R, 0, 1) * 255)
			g := uint8(util.Clamp(c.Pix[x+y*c.W].G, 0, 1) * 255)
			b := uint8(util.Clamp(c.Pix[x+y*c.W].B, 0, 1) * 255)
			img.Pix[(x+y*c.W)*4] = r
			img.Pix[(x+y*c.W)*4+1] = g
			img.Pix[(x+y*c.W)*4+2] = b
			img.Pix[(x+y*c.W)*4+3] = 255
		}
	}

	return img
}
