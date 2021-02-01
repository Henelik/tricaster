package canvas

import (
	"image"
	"testing"

	"github.com/Henelik/tricaster/color"
	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	w := 10
	h := 20

	c := NewCanvas(w, h)

	assert.Equal(t, w, c.W)
	assert.Equal(t, h, c.H)

	assert.Equal(t, w*h, len(c.Pix))
	assert.Equal(t, color.NewColor(0, 0, 0), &c.Pix[0])

	c2 := NewCanvas(0, 10)
	assert.Nil(t, c2)

	c3 := NewCanvas(10, 0)
	assert.Nil(t, c3)
}

func TestGet(t *testing.T) {
	w := 10
	h := 20
	black := color.NewColor(0, 0, 0)

	c := NewCanvas(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			assert.Equal(t, black, c.Get(x, y))
		}
	}
}

func TestSet(t *testing.T) {
	w := 10
	h := 20
	red := color.NewColor(1, 0, 0)

	c := NewCanvas(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c.Set(x, y, red)
		}
	}

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			assert.Equal(t, red, c.Get(x, y))
		}
	}
}

func TestToImage(t *testing.T) {
	canv := NewCanvas(32, 32)

	for x := 0; x < canv.W; x++ {
		for y := 0; y < canv.H; y++ {
			col := color.NewColor(
				float64(x)/float64(canv.W),
				float64(y)/float64(canv.H),
				0)
			canv.Set(x, y, col)
		}
	}

	img := canv.ToImage()

	wantPix := []byte{0x0, 0x0, 0x0, 0xff, 0x7, 0x0, 0x0, 0xff, 0xf, 0x0, 0x0, 0xff, 0x17, 0x0, 0x0, 0xff, 0x1f, 0x0, 0x0, 0xff, 0x27, 0x0, 0x0, 0xff, 0x2f, 0x0, 0x0, 0xff, 0x37, 0x0, 0x0, 0xff, 0x3f, 0x0, 0x0, 0xff, 0x47, 0x0, 0x0, 0xff, 0x4f, 0x0, 0x0, 0xff, 0x57, 0x0, 0x0, 0xff, 0x5f, 0x0, 0x0, 0xff, 0x67, 0x0, 0x0, 0xff, 0x6f, 0x0, 0x0, 0xff, 0x77, 0x0, 0x0, 0xff, 0x7f, 0x0, 0x0, 0xff, 0x87, 0x0, 0x0, 0xff, 0x8f, 0x0, 0x0, 0xff, 0x97, 0x0, 0x0, 0xff, 0x9f, 0x0, 0x0, 0xff, 0xa7, 0x0, 0x0, 0xff, 0xaf, 0x0, 0x0, 0xff, 0xb7, 0x0, 0x0, 0xff, 0xbf, 0x0, 0x0, 0xff, 0xc7, 0x0, 0x0, 0xff, 0xcf, 0x0, 0x0, 0xff, 0xd7, 0x0, 0x0, 0xff, 0xdf, 0x0, 0x0, 0xff, 0xe7, 0x0, 0x0, 0xff, 0xef, 0x0, 0x0, 0xff, 0xf7, 0x0, 0x0, 0xff, 0x0, 0x7, 0x0, 0xff, 0x7, 0x7, 0x0, 0xff, 0xf, 0x7, 0x0, 0xff, 0x17, 0x7, 0x0, 0xff, 0x1f, 0x7, 0x0, 0xff, 0x27, 0x7, 0x0, 0xff, 0x2f, 0x7, 0x0, 0xff, 0x37, 0x7, 0x0, 0xff, 0x3f, 0x7, 0x0, 0xff, 0x47, 0x7, 0x0, 0xff, 0x4f, 0x7, 0x0, 0xff, 0x57, 0x7, 0x0, 0xff, 0x5f, 0x7, 0x0, 0xff, 0x67, 0x7, 0x0, 0xff, 0x6f, 0x7, 0x0, 0xff, 0x77, 0x7, 0x0, 0xff, 0x7f, 0x7, 0x0, 0xff, 0x87, 0x7, 0x0, 0xff, 0x8f, 0x7, 0x0, 0xff, 0x97, 0x7, 0x0, 0xff, 0x9f, 0x7, 0x0, 0xff, 0xa7, 0x7, 0x0, 0xff, 0xaf, 0x7, 0x0, 0xff, 0xb7, 0x7, 0x0, 0xff, 0xbf, 0x7, 0x0, 0xff, 0xc7, 0x7, 0x0, 0xff, 0xcf, 0x7, 0x0, 0xff, 0xd7, 0x7, 0x0, 0xff, 0xdf, 0x7, 0x0, 0xff, 0xe7, 0x7, 0x0, 0xff, 0xef, 0x7, 0x0, 0xff, 0xf7, 0x7, 0x0, 0xff, 0x0, 0xf, 0x0, 0xff, 0x7, 0xf, 0x0, 0xff, 0xf, 0xf, 0x0, 0xff, 0x17, 0xf, 0x0, 0xff, 0x1f, 0xf, 0x0, 0xff, 0x27, 0xf, 0x0, 0xff, 0x2f, 0xf, 0x0, 0xff, 0x37, 0xf, 0x0, 0xff, 0x3f, 0xf, 0x0, 0xff, 0x47, 0xf, 0x0, 0xff, 0x4f, 0xf, 0x0, 0xff, 0x57, 0xf, 0x0, 0xff, 0x5f, 0xf, 0x0, 0xff, 0x67, 0xf, 0x0, 0xff, 0x6f, 0xf, 0x0, 0xff, 0x77, 0xf, 0x0, 0xff, 0x7f, 0xf, 0x0, 0xff, 0x87, 0xf, 0x0, 0xff, 0x8f, 0xf, 0x0, 0xff, 0x97, 0xf, 0x0, 0xff, 0x9f, 0xf, 0x0, 0xff, 0xa7, 0xf, 0x0, 0xff, 0xaf, 0xf, 0x0, 0xff, 0xb7, 0xf, 0x0, 0xff, 0xbf, 0xf, 0x0, 0xff, 0xc7, 0xf, 0x0, 0xff, 0xcf, 0xf, 0x0, 0xff, 0xd7, 0xf, 0x0, 0xff, 0xdf, 0xf, 0x0, 0xff, 0xe7, 0xf, 0x0, 0xff, 0xef, 0xf, 0x0, 0xff, 0xf7, 0xf, 0x0, 0xff, 0x0, 0x17, 0x0, 0xff, 0x7, 0x17, 0x0, 0xff, 0xf, 0x17, 0x0, 0xff, 0x17, 0x17, 0x0, 0xff, 0x1f, 0x17, 0x0, 0xff, 0x27, 0x17, 0x0, 0xff, 0x2f, 0x17, 0x0, 0xff, 0x37, 0x17, 0x0, 0xff, 0x3f, 0x17, 0x0, 0xff, 0x47, 0x17, 0x0, 0xff, 0x4f, 0x17, 0x0, 0xff, 0x57, 0x17, 0x0, 0xff, 0x5f, 0x17, 0x0, 0xff, 0x67, 0x17, 0x0, 0xff, 0x6f, 0x17, 0x0, 0xff, 0x77, 0x17, 0x0, 0xff, 0x7f, 0x17, 0x0, 0xff, 0x87, 0x17, 0x0, 0xff, 0x8f, 0x17, 0x0, 0xff, 0x97, 0x17, 0x0, 0xff, 0x9f, 0x17, 0x0, 0xff, 0xa7, 0x17, 0x0, 0xff, 0xaf, 0x17, 0x0, 0xff, 0xb7, 0x17, 0x0, 0xff, 0xbf, 0x17, 0x0, 0xff, 0xc7, 0x17, 0x0, 0xff, 0xcf, 0x17, 0x0, 0xff, 0xd7, 0x17, 0x0, 0xff, 0xdf, 0x17, 0x0, 0xff, 0xe7, 0x17, 0x0, 0xff, 0xef, 0x17, 0x0, 0xff, 0xf7, 0x17, 0x0, 0xff, 0x0, 0x1f, 0x0, 0xff, 0x7, 0x1f, 0x0, 0xff, 0xf, 0x1f, 0x0, 0xff, 0x17, 0x1f, 0x0, 0xff, 0x1f, 0x1f, 0x0, 0xff, 0x27, 0x1f, 0x0, 0xff, 0x2f, 0x1f, 0x0, 0xff, 0x37, 0x1f, 0x0, 0xff, 0x3f, 0x1f, 0x0, 0xff, 0x47, 0x1f, 0x0, 0xff, 0x4f, 0x1f, 0x0, 0xff, 0x57, 0x1f, 0x0, 0xff, 0x5f, 0x1f, 0x0, 0xff, 0x67, 0x1f, 0x0, 0xff, 0x6f, 0x1f, 0x0, 0xff, 0x77, 0x1f, 0x0, 0xff, 0x7f, 0x1f, 0x0, 0xff, 0x87, 0x1f, 0x0, 0xff, 0x8f, 0x1f, 0x0, 0xff, 0x97, 0x1f, 0x0, 0xff, 0x9f, 0x1f, 0x0, 0xff, 0xa7, 0x1f, 0x0, 0xff, 0xaf, 0x1f, 0x0, 0xff, 0xb7, 0x1f, 0x0, 0xff, 0xbf, 0x1f, 0x0, 0xff, 0xc7, 0x1f, 0x0, 0xff, 0xcf, 0x1f, 0x0, 0xff, 0xd7, 0x1f, 0x0, 0xff, 0xdf, 0x1f, 0x0, 0xff, 0xe7, 0x1f, 0x0, 0xff, 0xef, 0x1f, 0x0, 0xff, 0xf7, 0x1f, 0x0, 0xff, 0x0, 0x27, 0x0, 0xff, 0x7, 0x27, 0x0, 0xff, 0xf, 0x27, 0x0, 0xff, 0x17, 0x27, 0x0, 0xff, 0x1f, 0x27, 0x0, 0xff, 0x27, 0x27, 0x0, 0xff, 0x2f, 0x27, 0x0, 0xff, 0x37, 0x27, 0x0, 0xff, 0x3f, 0x27, 0x0, 0xff, 0x47, 0x27, 0x0, 0xff, 0x4f, 0x27, 0x0, 0xff, 0x57, 0x27, 0x0, 0xff, 0x5f, 0x27, 0x0, 0xff, 0x67, 0x27, 0x0, 0xff, 0x6f, 0x27, 0x0, 0xff, 0x77, 0x27, 0x0, 0xff, 0x7f, 0x27, 0x0, 0xff, 0x87, 0x27, 0x0, 0xff, 0x8f, 0x27, 0x0, 0xff, 0x97, 0x27, 0x0, 0xff, 0x9f, 0x27, 0x0, 0xff, 0xa7, 0x27, 0x0, 0xff, 0xaf, 0x27, 0x0, 0xff, 0xb7, 0x27, 0x0, 0xff, 0xbf, 0x27, 0x0, 0xff, 0xc7, 0x27, 0x0, 0xff, 0xcf, 0x27, 0x0, 0xff, 0xd7, 0x27, 0x0, 0xff, 0xdf, 0x27, 0x0, 0xff, 0xe7, 0x27, 0x0, 0xff, 0xef, 0x27, 0x0, 0xff, 0xf7, 0x27, 0x0, 0xff, 0x0, 0x2f, 0x0, 0xff, 0x7, 0x2f, 0x0, 0xff, 0xf, 0x2f, 0x0, 0xff, 0x17, 0x2f, 0x0, 0xff, 0x1f, 0x2f, 0x0, 0xff, 0x27, 0x2f, 0x0, 0xff, 0x2f, 0x2f, 0x0, 0xff, 0x37, 0x2f, 0x0, 0xff, 0x3f, 0x2f, 0x0, 0xff, 0x47, 0x2f, 0x0, 0xff, 0x4f, 0x2f, 0x0, 0xff, 0x57, 0x2f, 0x0, 0xff, 0x5f, 0x2f, 0x0, 0xff, 0x67, 0x2f, 0x0, 0xff, 0x6f, 0x2f, 0x0, 0xff, 0x77, 0x2f, 0x0, 0xff, 0x7f, 0x2f, 0x0, 0xff, 0x87, 0x2f, 0x0, 0xff, 0x8f, 0x2f, 0x0, 0xff, 0x97, 0x2f, 0x0, 0xff, 0x9f, 0x2f, 0x0, 0xff, 0xa7, 0x2f, 0x0, 0xff, 0xaf, 0x2f, 0x0, 0xff, 0xb7, 0x2f, 0x0, 0xff, 0xbf, 0x2f, 0x0, 0xff, 0xc7, 0x2f, 0x0, 0xff, 0xcf, 0x2f, 0x0, 0xff, 0xd7, 0x2f, 0x0, 0xff, 0xdf, 0x2f, 0x0, 0xff, 0xe7, 0x2f, 0x0, 0xff, 0xef, 0x2f, 0x0, 0xff, 0xf7, 0x2f, 0x0, 0xff, 0x0, 0x37, 0x0, 0xff, 0x7, 0x37, 0x0, 0xff, 0xf, 0x37, 0x0, 0xff, 0x17, 0x37, 0x0, 0xff, 0x1f, 0x37, 0x0, 0xff, 0x27, 0x37, 0x0, 0xff, 0x2f, 0x37, 0x0, 0xff, 0x37, 0x37, 0x0, 0xff, 0x3f, 0x37, 0x0, 0xff, 0x47, 0x37, 0x0, 0xff, 0x4f, 0x37, 0x0, 0xff, 0x57, 0x37, 0x0, 0xff, 0x5f, 0x37, 0x0, 0xff, 0x67, 0x37, 0x0, 0xff, 0x6f, 0x37, 0x0, 0xff, 0x77, 0x37, 0x0, 0xff, 0x7f, 0x37, 0x0, 0xff, 0x87, 0x37, 0x0, 0xff, 0x8f, 0x37, 0x0, 0xff, 0x97, 0x37, 0x0, 0xff, 0x9f, 0x37, 0x0, 0xff, 0xa7, 0x37, 0x0, 0xff, 0xaf, 0x37, 0x0, 0xff, 0xb7, 0x37, 0x0, 0xff, 0xbf, 0x37, 0x0, 0xff, 0xc7, 0x37, 0x0, 0xff, 0xcf, 0x37, 0x0, 0xff, 0xd7, 0x37, 0x0, 0xff, 0xdf, 0x37, 0x0, 0xff, 0xe7, 0x37, 0x0, 0xff, 0xef, 0x37, 0x0, 0xff, 0xf7, 0x37, 0x0, 0xff, 0x0, 0x3f, 0x0, 0xff, 0x7, 0x3f, 0x0, 0xff, 0xf, 0x3f, 0x0, 0xff, 0x17, 0x3f, 0x0, 0xff, 0x1f, 0x3f, 0x0, 0xff, 0x27, 0x3f, 0x0, 0xff, 0x2f, 0x3f, 0x0, 0xff, 0x37, 0x3f, 0x0, 0xff, 0x3f, 0x3f, 0x0, 0xff, 0x47, 0x3f, 0x0, 0xff, 0x4f, 0x3f, 0x0, 0xff, 0x57, 0x3f, 0x0, 0xff, 0x5f, 0x3f, 0x0, 0xff, 0x67, 0x3f, 0x0, 0xff, 0x6f, 0x3f, 0x0, 0xff, 0x77, 0x3f, 0x0, 0xff, 0x7f, 0x3f, 0x0, 0xff, 0x87, 0x3f, 0x0, 0xff, 0x8f, 0x3f, 0x0, 0xff, 0x97, 0x3f, 0x0, 0xff, 0x9f, 0x3f, 0x0, 0xff, 0xa7, 0x3f, 0x0, 0xff, 0xaf, 0x3f, 0x0, 0xff, 0xb7, 0x3f, 0x0, 0xff, 0xbf, 0x3f, 0x0, 0xff, 0xc7, 0x3f, 0x0, 0xff, 0xcf, 0x3f, 0x0, 0xff, 0xd7, 0x3f, 0x0, 0xff, 0xdf, 0x3f, 0x0, 0xff, 0xe7, 0x3f, 0x0, 0xff, 0xef, 0x3f, 0x0, 0xff, 0xf7, 0x3f, 0x0, 0xff, 0x0, 0x47, 0x0, 0xff, 0x7, 0x47, 0x0, 0xff, 0xf, 0x47, 0x0, 0xff, 0x17, 0x47, 0x0, 0xff, 0x1f, 0x47, 0x0, 0xff, 0x27, 0x47, 0x0, 0xff, 0x2f, 0x47, 0x0, 0xff, 0x37, 0x47, 0x0, 0xff, 0x3f, 0x47, 0x0, 0xff, 0x47, 0x47, 0x0, 0xff, 0x4f, 0x47, 0x0, 0xff, 0x57, 0x47, 0x0, 0xff, 0x5f, 0x47, 0x0, 0xff, 0x67, 0x47, 0x0, 0xff, 0x6f, 0x47, 0x0, 0xff, 0x77, 0x47, 0x0, 0xff, 0x7f, 0x47, 0x0, 0xff, 0x87, 0x47, 0x0, 0xff, 0x8f, 0x47, 0x0, 0xff, 0x97, 0x47, 0x0, 0xff, 0x9f, 0x47, 0x0, 0xff, 0xa7, 0x47, 0x0, 0xff, 0xaf, 0x47, 0x0, 0xff, 0xb7, 0x47, 0x0, 0xff, 0xbf, 0x47, 0x0, 0xff, 0xc7, 0x47, 0x0, 0xff, 0xcf, 0x47, 0x0, 0xff, 0xd7, 0x47, 0x0, 0xff, 0xdf, 0x47, 0x0, 0xff, 0xe7, 0x47, 0x0, 0xff, 0xef, 0x47, 0x0, 0xff, 0xf7, 0x47, 0x0, 0xff, 0x0, 0x4f, 0x0, 0xff, 0x7, 0x4f, 0x0, 0xff, 0xf, 0x4f, 0x0, 0xff, 0x17, 0x4f, 0x0, 0xff, 0x1f, 0x4f, 0x0, 0xff, 0x27, 0x4f, 0x0, 0xff, 0x2f, 0x4f, 0x0, 0xff, 0x37, 0x4f, 0x0, 0xff, 0x3f, 0x4f, 0x0, 0xff, 0x47, 0x4f, 0x0, 0xff, 0x4f, 0x4f, 0x0, 0xff, 0x57, 0x4f, 0x0, 0xff, 0x5f, 0x4f, 0x0, 0xff, 0x67, 0x4f, 0x0, 0xff, 0x6f, 0x4f, 0x0, 0xff, 0x77, 0x4f, 0x0, 0xff, 0x7f, 0x4f, 0x0, 0xff, 0x87, 0x4f, 0x0, 0xff, 0x8f, 0x4f, 0x0, 0xff, 0x97, 0x4f, 0x0, 0xff, 0x9f, 0x4f, 0x0, 0xff, 0xa7, 0x4f, 0x0, 0xff, 0xaf, 0x4f, 0x0, 0xff, 0xb7, 0x4f, 0x0, 0xff, 0xbf, 0x4f, 0x0, 0xff, 0xc7, 0x4f, 0x0, 0xff, 0xcf, 0x4f, 0x0, 0xff, 0xd7, 0x4f, 0x0, 0xff, 0xdf, 0x4f, 0x0, 0xff, 0xe7, 0x4f, 0x0, 0xff, 0xef, 0x4f, 0x0, 0xff, 0xf7, 0x4f, 0x0, 0xff, 0x0, 0x57, 0x0, 0xff, 0x7, 0x57, 0x0, 0xff, 0xf, 0x57, 0x0, 0xff, 0x17, 0x57, 0x0, 0xff, 0x1f, 0x57, 0x0, 0xff, 0x27, 0x57, 0x0, 0xff, 0x2f, 0x57, 0x0, 0xff, 0x37, 0x57, 0x0, 0xff, 0x3f, 0x57, 0x0, 0xff, 0x47, 0x57, 0x0, 0xff, 0x4f, 0x57, 0x0, 0xff, 0x57, 0x57, 0x0, 0xff, 0x5f, 0x57, 0x0, 0xff, 0x67, 0x57, 0x0, 0xff, 0x6f, 0x57, 0x0, 0xff, 0x77, 0x57, 0x0, 0xff, 0x7f, 0x57, 0x0, 0xff, 0x87, 0x57, 0x0, 0xff, 0x8f, 0x57, 0x0, 0xff, 0x97, 0x57, 0x0, 0xff, 0x9f, 0x57, 0x0, 0xff, 0xa7, 0x57, 0x0, 0xff, 0xaf, 0x57, 0x0, 0xff, 0xb7, 0x57, 0x0, 0xff, 0xbf, 0x57, 0x0, 0xff, 0xc7, 0x57, 0x0, 0xff, 0xcf, 0x57, 0x0, 0xff, 0xd7, 0x57, 0x0, 0xff, 0xdf, 0x57, 0x0, 0xff, 0xe7, 0x57, 0x0, 0xff, 0xef, 0x57, 0x0, 0xff, 0xf7, 0x57, 0x0, 0xff, 0x0, 0x5f, 0x0, 0xff, 0x7, 0x5f, 0x0, 0xff, 0xf, 0x5f, 0x0, 0xff, 0x17, 0x5f, 0x0, 0xff, 0x1f, 0x5f, 0x0, 0xff, 0x27, 0x5f, 0x0, 0xff, 0x2f, 0x5f, 0x0, 0xff, 0x37, 0x5f, 0x0, 0xff, 0x3f, 0x5f, 0x0, 0xff, 0x47, 0x5f, 0x0, 0xff, 0x4f, 0x5f, 0x0, 0xff, 0x57, 0x5f, 0x0, 0xff, 0x5f, 0x5f, 0x0, 0xff, 0x67, 0x5f, 0x0, 0xff, 0x6f, 0x5f, 0x0, 0xff, 0x77, 0x5f, 0x0, 0xff, 0x7f, 0x5f, 0x0, 0xff, 0x87, 0x5f, 0x0, 0xff, 0x8f, 0x5f, 0x0, 0xff, 0x97, 0x5f, 0x0, 0xff, 0x9f, 0x5f, 0x0, 0xff, 0xa7, 0x5f, 0x0, 0xff, 0xaf, 0x5f, 0x0, 0xff, 0xb7, 0x5f, 0x0, 0xff, 0xbf, 0x5f, 0x0, 0xff, 0xc7, 0x5f, 0x0, 0xff, 0xcf, 0x5f, 0x0, 0xff, 0xd7, 0x5f, 0x0, 0xff, 0xdf, 0x5f, 0x0, 0xff, 0xe7, 0x5f, 0x0, 0xff, 0xef, 0x5f, 0x0, 0xff, 0xf7, 0x5f, 0x0, 0xff, 0x0, 0x67, 0x0, 0xff, 0x7, 0x67, 0x0, 0xff, 0xf, 0x67, 0x0, 0xff, 0x17, 0x67, 0x0, 0xff, 0x1f, 0x67, 0x0, 0xff, 0x27, 0x67, 0x0, 0xff, 0x2f, 0x67, 0x0, 0xff, 0x37, 0x67, 0x0, 0xff, 0x3f, 0x67, 0x0, 0xff, 0x47, 0x67, 0x0, 0xff, 0x4f, 0x67, 0x0, 0xff, 0x57, 0x67, 0x0, 0xff, 0x5f, 0x67, 0x0, 0xff, 0x67, 0x67, 0x0, 0xff, 0x6f, 0x67, 0x0, 0xff, 0x77, 0x67, 0x0, 0xff, 0x7f, 0x67, 0x0, 0xff, 0x87, 0x67, 0x0, 0xff, 0x8f, 0x67, 0x0, 0xff, 0x97, 0x67, 0x0, 0xff, 0x9f, 0x67, 0x0, 0xff, 0xa7, 0x67, 0x0, 0xff, 0xaf, 0x67, 0x0, 0xff, 0xb7, 0x67, 0x0, 0xff, 0xbf, 0x67, 0x0, 0xff, 0xc7, 0x67, 0x0, 0xff, 0xcf, 0x67, 0x0, 0xff, 0xd7, 0x67, 0x0, 0xff, 0xdf, 0x67, 0x0, 0xff, 0xe7, 0x67, 0x0, 0xff, 0xef, 0x67, 0x0, 0xff, 0xf7, 0x67, 0x0, 0xff, 0x0, 0x6f, 0x0, 0xff, 0x7, 0x6f, 0x0, 0xff, 0xf, 0x6f, 0x0, 0xff, 0x17, 0x6f, 0x0, 0xff, 0x1f, 0x6f, 0x0, 0xff, 0x27, 0x6f, 0x0, 0xff, 0x2f, 0x6f, 0x0, 0xff, 0x37, 0x6f, 0x0, 0xff, 0x3f, 0x6f, 0x0, 0xff, 0x47, 0x6f, 0x0, 0xff, 0x4f, 0x6f, 0x0, 0xff, 0x57, 0x6f, 0x0, 0xff, 0x5f, 0x6f, 0x0, 0xff, 0x67, 0x6f, 0x0, 0xff, 0x6f, 0x6f, 0x0, 0xff, 0x77, 0x6f, 0x0, 0xff, 0x7f, 0x6f, 0x0, 0xff, 0x87, 0x6f, 0x0, 0xff, 0x8f, 0x6f, 0x0, 0xff, 0x97, 0x6f, 0x0, 0xff, 0x9f, 0x6f, 0x0, 0xff, 0xa7, 0x6f, 0x0, 0xff, 0xaf, 0x6f, 0x0, 0xff, 0xb7, 0x6f, 0x0, 0xff, 0xbf, 0x6f, 0x0, 0xff, 0xc7, 0x6f, 0x0, 0xff, 0xcf, 0x6f, 0x0, 0xff, 0xd7, 0x6f, 0x0, 0xff, 0xdf, 0x6f, 0x0, 0xff, 0xe7, 0x6f, 0x0, 0xff, 0xef, 0x6f, 0x0, 0xff, 0xf7, 0x6f, 0x0, 0xff, 0x0, 0x77, 0x0, 0xff, 0x7, 0x77, 0x0, 0xff, 0xf, 0x77, 0x0, 0xff, 0x17, 0x77, 0x0, 0xff, 0x1f, 0x77, 0x0, 0xff, 0x27, 0x77, 0x0, 0xff, 0x2f, 0x77, 0x0, 0xff, 0x37, 0x77, 0x0, 0xff, 0x3f, 0x77, 0x0, 0xff, 0x47, 0x77, 0x0, 0xff, 0x4f, 0x77, 0x0, 0xff, 0x57, 0x77, 0x0, 0xff, 0x5f, 0x77, 0x0, 0xff, 0x67, 0x77, 0x0, 0xff, 0x6f, 0x77, 0x0, 0xff, 0x77, 0x77, 0x0, 0xff, 0x7f, 0x77, 0x0, 0xff, 0x87, 0x77, 0x0, 0xff, 0x8f, 0x77, 0x0, 0xff, 0x97, 0x77, 0x0, 0xff, 0x9f, 0x77, 0x0, 0xff, 0xa7, 0x77, 0x0, 0xff, 0xaf, 0x77, 0x0, 0xff, 0xb7, 0x77, 0x0, 0xff, 0xbf, 0x77, 0x0, 0xff, 0xc7, 0x77, 0x0, 0xff, 0xcf, 0x77, 0x0, 0xff, 0xd7, 0x77, 0x0, 0xff, 0xdf, 0x77, 0x0, 0xff, 0xe7, 0x77, 0x0, 0xff, 0xef, 0x77, 0x0, 0xff, 0xf7, 0x77, 0x0, 0xff, 0x0, 0x7f, 0x0, 0xff, 0x7, 0x7f, 0x0, 0xff, 0xf, 0x7f, 0x0, 0xff, 0x17, 0x7f, 0x0, 0xff, 0x1f, 0x7f, 0x0, 0xff, 0x27, 0x7f, 0x0, 0xff, 0x2f, 0x7f, 0x0, 0xff, 0x37, 0x7f, 0x0, 0xff, 0x3f, 0x7f, 0x0, 0xff, 0x47, 0x7f, 0x0, 0xff, 0x4f, 0x7f, 0x0, 0xff, 0x57, 0x7f, 0x0, 0xff, 0x5f, 0x7f, 0x0, 0xff, 0x67, 0x7f, 0x0, 0xff, 0x6f, 0x7f, 0x0, 0xff, 0x77, 0x7f, 0x0, 0xff, 0x7f, 0x7f, 0x0, 0xff, 0x87, 0x7f, 0x0, 0xff, 0x8f, 0x7f, 0x0, 0xff, 0x97, 0x7f, 0x0, 0xff, 0x9f, 0x7f, 0x0, 0xff, 0xa7, 0x7f, 0x0, 0xff, 0xaf, 0x7f, 0x0, 0xff, 0xb7, 0x7f, 0x0, 0xff, 0xbf, 0x7f, 0x0, 0xff, 0xc7, 0x7f, 0x0, 0xff, 0xcf, 0x7f, 0x0, 0xff, 0xd7, 0x7f, 0x0, 0xff, 0xdf, 0x7f, 0x0, 0xff, 0xe7, 0x7f, 0x0, 0xff, 0xef, 0x7f, 0x0, 0xff, 0xf7, 0x7f, 0x0, 0xff, 0x0, 0x87, 0x0, 0xff, 0x7, 0x87, 0x0, 0xff, 0xf, 0x87, 0x0, 0xff, 0x17, 0x87, 0x0, 0xff, 0x1f, 0x87, 0x0, 0xff, 0x27, 0x87, 0x0, 0xff, 0x2f, 0x87, 0x0, 0xff, 0x37, 0x87, 0x0, 0xff, 0x3f, 0x87, 0x0, 0xff, 0x47, 0x87, 0x0, 0xff, 0x4f, 0x87, 0x0, 0xff, 0x57, 0x87, 0x0, 0xff, 0x5f, 0x87, 0x0, 0xff, 0x67, 0x87, 0x0, 0xff, 0x6f, 0x87, 0x0, 0xff, 0x77, 0x87, 0x0, 0xff, 0x7f, 0x87, 0x0, 0xff, 0x87, 0x87, 0x0, 0xff, 0x8f, 0x87, 0x0, 0xff, 0x97, 0x87, 0x0, 0xff, 0x9f, 0x87, 0x0, 0xff, 0xa7, 0x87, 0x0, 0xff, 0xaf, 0x87, 0x0, 0xff, 0xb7, 0x87, 0x0, 0xff, 0xbf, 0x87, 0x0, 0xff, 0xc7, 0x87, 0x0, 0xff, 0xcf, 0x87, 0x0, 0xff, 0xd7, 0x87, 0x0, 0xff, 0xdf, 0x87, 0x0, 0xff, 0xe7, 0x87, 0x0, 0xff, 0xef, 0x87, 0x0, 0xff, 0xf7, 0x87, 0x0, 0xff, 0x0, 0x8f, 0x0, 0xff, 0x7, 0x8f, 0x0, 0xff, 0xf, 0x8f, 0x0, 0xff, 0x17, 0x8f, 0x0, 0xff, 0x1f, 0x8f, 0x0, 0xff, 0x27, 0x8f, 0x0, 0xff, 0x2f, 0x8f, 0x0, 0xff, 0x37, 0x8f, 0x0, 0xff, 0x3f, 0x8f, 0x0, 0xff, 0x47, 0x8f, 0x0, 0xff, 0x4f, 0x8f, 0x0, 0xff, 0x57, 0x8f, 0x0, 0xff, 0x5f, 0x8f, 0x0, 0xff, 0x67, 0x8f, 0x0, 0xff, 0x6f, 0x8f, 0x0, 0xff, 0x77, 0x8f, 0x0, 0xff, 0x7f, 0x8f, 0x0, 0xff, 0x87, 0x8f, 0x0, 0xff, 0x8f, 0x8f, 0x0, 0xff, 0x97, 0x8f, 0x0, 0xff, 0x9f, 0x8f, 0x0, 0xff, 0xa7, 0x8f, 0x0, 0xff, 0xaf, 0x8f, 0x0, 0xff, 0xb7, 0x8f, 0x0, 0xff, 0xbf, 0x8f, 0x0, 0xff, 0xc7, 0x8f, 0x0, 0xff, 0xcf, 0x8f, 0x0, 0xff, 0xd7, 0x8f, 0x0, 0xff, 0xdf, 0x8f, 0x0, 0xff, 0xe7, 0x8f, 0x0, 0xff, 0xef, 0x8f, 0x0, 0xff, 0xf7, 0x8f, 0x0, 0xff, 0x0, 0x97, 0x0, 0xff, 0x7, 0x97, 0x0, 0xff, 0xf, 0x97, 0x0, 0xff, 0x17, 0x97, 0x0, 0xff, 0x1f, 0x97, 0x0, 0xff, 0x27, 0x97, 0x0, 0xff, 0x2f, 0x97, 0x0, 0xff, 0x37, 0x97, 0x0, 0xff, 0x3f, 0x97, 0x0, 0xff, 0x47, 0x97, 0x0, 0xff, 0x4f, 0x97, 0x0, 0xff, 0x57, 0x97, 0x0, 0xff, 0x5f, 0x97, 0x0, 0xff, 0x67, 0x97, 0x0, 0xff, 0x6f, 0x97, 0x0, 0xff, 0x77, 0x97, 0x0, 0xff, 0x7f, 0x97, 0x0, 0xff, 0x87, 0x97, 0x0, 0xff, 0x8f, 0x97, 0x0, 0xff, 0x97, 0x97, 0x0, 0xff, 0x9f, 0x97, 0x0, 0xff, 0xa7, 0x97, 0x0, 0xff, 0xaf, 0x97, 0x0, 0xff, 0xb7, 0x97, 0x0, 0xff, 0xbf, 0x97, 0x0, 0xff, 0xc7, 0x97, 0x0, 0xff, 0xcf, 0x97, 0x0, 0xff, 0xd7, 0x97, 0x0, 0xff, 0xdf, 0x97, 0x0, 0xff, 0xe7, 0x97, 0x0, 0xff, 0xef, 0x97, 0x0, 0xff, 0xf7, 0x97, 0x0, 0xff, 0x0, 0x9f, 0x0, 0xff, 0x7, 0x9f, 0x0, 0xff, 0xf, 0x9f, 0x0, 0xff, 0x17, 0x9f, 0x0, 0xff, 0x1f, 0x9f, 0x0, 0xff, 0x27, 0x9f, 0x0, 0xff, 0x2f, 0x9f, 0x0, 0xff, 0x37, 0x9f, 0x0, 0xff, 0x3f, 0x9f, 0x0, 0xff, 0x47, 0x9f, 0x0, 0xff, 0x4f, 0x9f, 0x0, 0xff, 0x57, 0x9f, 0x0, 0xff, 0x5f, 0x9f, 0x0, 0xff, 0x67, 0x9f, 0x0, 0xff, 0x6f, 0x9f, 0x0, 0xff, 0x77, 0x9f, 0x0, 0xff, 0x7f, 0x9f, 0x0, 0xff, 0x87, 0x9f, 0x0, 0xff, 0x8f, 0x9f, 0x0, 0xff, 0x97, 0x9f, 0x0, 0xff, 0x9f, 0x9f, 0x0, 0xff, 0xa7, 0x9f, 0x0, 0xff, 0xaf, 0x9f, 0x0, 0xff, 0xb7, 0x9f, 0x0, 0xff, 0xbf, 0x9f, 0x0, 0xff, 0xc7, 0x9f, 0x0, 0xff, 0xcf, 0x9f, 0x0, 0xff, 0xd7, 0x9f, 0x0, 0xff, 0xdf, 0x9f, 0x0, 0xff, 0xe7, 0x9f, 0x0, 0xff, 0xef, 0x9f, 0x0, 0xff, 0xf7, 0x9f, 0x0, 0xff, 0x0, 0xa7, 0x0, 0xff, 0x7, 0xa7, 0x0, 0xff, 0xf, 0xa7, 0x0, 0xff, 0x17, 0xa7, 0x0, 0xff, 0x1f, 0xa7, 0x0, 0xff, 0x27, 0xa7, 0x0, 0xff, 0x2f, 0xa7, 0x0, 0xff, 0x37, 0xa7, 0x0, 0xff, 0x3f, 0xa7, 0x0, 0xff, 0x47, 0xa7, 0x0, 0xff, 0x4f, 0xa7, 0x0, 0xff, 0x57, 0xa7, 0x0, 0xff, 0x5f, 0xa7, 0x0, 0xff, 0x67, 0xa7, 0x0, 0xff, 0x6f, 0xa7, 0x0, 0xff, 0x77, 0xa7, 0x0, 0xff, 0x7f, 0xa7, 0x0, 0xff, 0x87, 0xa7, 0x0, 0xff, 0x8f, 0xa7, 0x0, 0xff, 0x97, 0xa7, 0x0, 0xff, 0x9f, 0xa7, 0x0, 0xff, 0xa7, 0xa7, 0x0, 0xff, 0xaf, 0xa7, 0x0, 0xff, 0xb7, 0xa7, 0x0, 0xff, 0xbf, 0xa7, 0x0, 0xff, 0xc7, 0xa7, 0x0, 0xff, 0xcf, 0xa7, 0x0, 0xff, 0xd7, 0xa7, 0x0, 0xff, 0xdf, 0xa7, 0x0, 0xff, 0xe7, 0xa7, 0x0, 0xff, 0xef, 0xa7, 0x0, 0xff, 0xf7, 0xa7, 0x0, 0xff, 0x0, 0xaf, 0x0, 0xff, 0x7, 0xaf, 0x0, 0xff, 0xf, 0xaf, 0x0, 0xff, 0x17, 0xaf, 0x0, 0xff, 0x1f, 0xaf, 0x0, 0xff, 0x27, 0xaf, 0x0, 0xff, 0x2f, 0xaf, 0x0, 0xff, 0x37, 0xaf, 0x0, 0xff, 0x3f, 0xaf, 0x0, 0xff, 0x47, 0xaf, 0x0, 0xff, 0x4f, 0xaf, 0x0, 0xff, 0x57, 0xaf, 0x0, 0xff, 0x5f, 0xaf, 0x0, 0xff, 0x67, 0xaf, 0x0, 0xff, 0x6f, 0xaf, 0x0, 0xff, 0x77, 0xaf, 0x0, 0xff, 0x7f, 0xaf, 0x0, 0xff, 0x87, 0xaf, 0x0, 0xff, 0x8f, 0xaf, 0x0, 0xff, 0x97, 0xaf, 0x0, 0xff, 0x9f, 0xaf, 0x0, 0xff, 0xa7, 0xaf, 0x0, 0xff, 0xaf, 0xaf, 0x0, 0xff, 0xb7, 0xaf, 0x0, 0xff, 0xbf, 0xaf, 0x0, 0xff, 0xc7, 0xaf, 0x0, 0xff, 0xcf, 0xaf, 0x0, 0xff, 0xd7, 0xaf, 0x0, 0xff, 0xdf, 0xaf, 0x0, 0xff, 0xe7, 0xaf, 0x0, 0xff, 0xef, 0xaf, 0x0, 0xff, 0xf7, 0xaf, 0x0, 0xff, 0x0, 0xb7, 0x0, 0xff, 0x7, 0xb7, 0x0, 0xff, 0xf, 0xb7, 0x0, 0xff, 0x17, 0xb7, 0x0, 0xff, 0x1f, 0xb7, 0x0, 0xff, 0x27, 0xb7, 0x0, 0xff, 0x2f, 0xb7, 0x0, 0xff, 0x37, 0xb7, 0x0, 0xff, 0x3f, 0xb7, 0x0, 0xff, 0x47, 0xb7, 0x0, 0xff, 0x4f, 0xb7, 0x0, 0xff, 0x57, 0xb7, 0x0, 0xff, 0x5f, 0xb7, 0x0, 0xff, 0x67, 0xb7, 0x0, 0xff, 0x6f, 0xb7, 0x0, 0xff, 0x77, 0xb7, 0x0, 0xff, 0x7f, 0xb7, 0x0, 0xff, 0x87, 0xb7, 0x0, 0xff, 0x8f, 0xb7, 0x0, 0xff, 0x97, 0xb7, 0x0, 0xff, 0x9f, 0xb7, 0x0, 0xff, 0xa7, 0xb7, 0x0, 0xff, 0xaf, 0xb7, 0x0, 0xff, 0xb7, 0xb7, 0x0, 0xff, 0xbf, 0xb7, 0x0, 0xff, 0xc7, 0xb7, 0x0, 0xff, 0xcf, 0xb7, 0x0, 0xff, 0xd7, 0xb7, 0x0, 0xff, 0xdf, 0xb7, 0x0, 0xff, 0xe7, 0xb7, 0x0, 0xff, 0xef, 0xb7, 0x0, 0xff, 0xf7, 0xb7, 0x0, 0xff, 0x0, 0xbf, 0x0, 0xff, 0x7, 0xbf, 0x0, 0xff, 0xf, 0xbf, 0x0, 0xff, 0x17, 0xbf, 0x0, 0xff, 0x1f, 0xbf, 0x0, 0xff, 0x27, 0xbf, 0x0, 0xff, 0x2f, 0xbf, 0x0, 0xff, 0x37, 0xbf, 0x0, 0xff, 0x3f, 0xbf, 0x0, 0xff, 0x47, 0xbf, 0x0, 0xff, 0x4f, 0xbf, 0x0, 0xff, 0x57, 0xbf, 0x0, 0xff, 0x5f, 0xbf, 0x0, 0xff, 0x67, 0xbf, 0x0, 0xff, 0x6f, 0xbf, 0x0, 0xff, 0x77, 0xbf, 0x0, 0xff, 0x7f, 0xbf, 0x0, 0xff, 0x87, 0xbf, 0x0, 0xff, 0x8f, 0xbf, 0x0, 0xff, 0x97, 0xbf, 0x0, 0xff, 0x9f, 0xbf, 0x0, 0xff, 0xa7, 0xbf, 0x0, 0xff, 0xaf, 0xbf, 0x0, 0xff, 0xb7, 0xbf, 0x0, 0xff, 0xbf, 0xbf, 0x0, 0xff, 0xc7, 0xbf, 0x0, 0xff, 0xcf, 0xbf, 0x0, 0xff, 0xd7, 0xbf, 0x0, 0xff, 0xdf, 0xbf, 0x0, 0xff, 0xe7, 0xbf, 0x0, 0xff, 0xef, 0xbf, 0x0, 0xff, 0xf7, 0xbf, 0x0, 0xff, 0x0, 0xc7, 0x0, 0xff, 0x7, 0xc7, 0x0, 0xff, 0xf, 0xc7, 0x0, 0xff, 0x17, 0xc7, 0x0, 0xff, 0x1f, 0xc7, 0x0, 0xff, 0x27, 0xc7, 0x0, 0xff, 0x2f, 0xc7, 0x0, 0xff, 0x37, 0xc7, 0x0, 0xff, 0x3f, 0xc7, 0x0, 0xff, 0x47, 0xc7, 0x0, 0xff, 0x4f, 0xc7, 0x0, 0xff, 0x57, 0xc7, 0x0, 0xff, 0x5f, 0xc7, 0x0, 0xff, 0x67, 0xc7, 0x0, 0xff, 0x6f, 0xc7, 0x0, 0xff, 0x77, 0xc7, 0x0, 0xff, 0x7f, 0xc7, 0x0, 0xff, 0x87, 0xc7, 0x0, 0xff, 0x8f, 0xc7, 0x0, 0xff, 0x97, 0xc7, 0x0, 0xff, 0x9f, 0xc7, 0x0, 0xff, 0xa7, 0xc7, 0x0, 0xff, 0xaf, 0xc7, 0x0, 0xff, 0xb7, 0xc7, 0x0, 0xff, 0xbf, 0xc7, 0x0, 0xff, 0xc7, 0xc7, 0x0, 0xff, 0xcf, 0xc7, 0x0, 0xff, 0xd7, 0xc7, 0x0, 0xff, 0xdf, 0xc7, 0x0, 0xff, 0xe7, 0xc7, 0x0, 0xff, 0xef, 0xc7, 0x0, 0xff, 0xf7, 0xc7, 0x0, 0xff, 0x0, 0xcf, 0x0, 0xff, 0x7, 0xcf, 0x0, 0xff, 0xf, 0xcf, 0x0, 0xff, 0x17, 0xcf, 0x0, 0xff, 0x1f, 0xcf, 0x0, 0xff, 0x27, 0xcf, 0x0, 0xff, 0x2f, 0xcf, 0x0, 0xff, 0x37, 0xcf, 0x0, 0xff, 0x3f, 0xcf, 0x0, 0xff, 0x47, 0xcf, 0x0, 0xff, 0x4f, 0xcf, 0x0, 0xff, 0x57, 0xcf, 0x0, 0xff, 0x5f, 0xcf, 0x0, 0xff, 0x67, 0xcf, 0x0, 0xff, 0x6f, 0xcf, 0x0, 0xff, 0x77, 0xcf, 0x0, 0xff, 0x7f, 0xcf, 0x0, 0xff, 0x87, 0xcf, 0x0, 0xff, 0x8f, 0xcf, 0x0, 0xff, 0x97, 0xcf, 0x0, 0xff, 0x9f, 0xcf, 0x0, 0xff, 0xa7, 0xcf, 0x0, 0xff, 0xaf, 0xcf, 0x0, 0xff, 0xb7, 0xcf, 0x0, 0xff, 0xbf, 0xcf, 0x0, 0xff, 0xc7, 0xcf, 0x0, 0xff, 0xcf, 0xcf, 0x0, 0xff, 0xd7, 0xcf, 0x0, 0xff, 0xdf, 0xcf, 0x0, 0xff, 0xe7, 0xcf, 0x0, 0xff, 0xef, 0xcf, 0x0, 0xff, 0xf7, 0xcf, 0x0, 0xff, 0x0, 0xd7, 0x0, 0xff, 0x7, 0xd7, 0x0, 0xff, 0xf, 0xd7, 0x0, 0xff, 0x17, 0xd7, 0x0, 0xff, 0x1f, 0xd7, 0x0, 0xff, 0x27, 0xd7, 0x0, 0xff, 0x2f, 0xd7, 0x0, 0xff, 0x37, 0xd7, 0x0, 0xff, 0x3f, 0xd7, 0x0, 0xff, 0x47, 0xd7, 0x0, 0xff, 0x4f, 0xd7, 0x0, 0xff, 0x57, 0xd7, 0x0, 0xff, 0x5f, 0xd7, 0x0, 0xff, 0x67, 0xd7, 0x0, 0xff, 0x6f, 0xd7, 0x0, 0xff, 0x77, 0xd7, 0x0, 0xff, 0x7f, 0xd7, 0x0, 0xff, 0x87, 0xd7, 0x0, 0xff, 0x8f, 0xd7, 0x0, 0xff, 0x97, 0xd7, 0x0, 0xff, 0x9f, 0xd7, 0x0, 0xff, 0xa7, 0xd7, 0x0, 0xff, 0xaf, 0xd7, 0x0, 0xff, 0xb7, 0xd7, 0x0, 0xff, 0xbf, 0xd7, 0x0, 0xff, 0xc7, 0xd7, 0x0, 0xff, 0xcf, 0xd7, 0x0, 0xff, 0xd7, 0xd7, 0x0, 0xff, 0xdf, 0xd7, 0x0, 0xff, 0xe7, 0xd7, 0x0, 0xff, 0xef, 0xd7, 0x0, 0xff, 0xf7, 0xd7, 0x0, 0xff, 0x0, 0xdf, 0x0, 0xff, 0x7, 0xdf, 0x0, 0xff, 0xf, 0xdf, 0x0, 0xff, 0x17, 0xdf, 0x0, 0xff, 0x1f, 0xdf, 0x0, 0xff, 0x27, 0xdf, 0x0, 0xff, 0x2f, 0xdf, 0x0, 0xff, 0x37, 0xdf, 0x0, 0xff, 0x3f, 0xdf, 0x0, 0xff, 0x47, 0xdf, 0x0, 0xff, 0x4f, 0xdf, 0x0, 0xff, 0x57, 0xdf, 0x0, 0xff, 0x5f, 0xdf, 0x0, 0xff, 0x67, 0xdf, 0x0, 0xff, 0x6f, 0xdf, 0x0, 0xff, 0x77, 0xdf, 0x0, 0xff, 0x7f, 0xdf, 0x0, 0xff, 0x87, 0xdf, 0x0, 0xff, 0x8f, 0xdf, 0x0, 0xff, 0x97, 0xdf, 0x0, 0xff, 0x9f, 0xdf, 0x0, 0xff, 0xa7, 0xdf, 0x0, 0xff, 0xaf, 0xdf, 0x0, 0xff, 0xb7, 0xdf, 0x0, 0xff, 0xbf, 0xdf, 0x0, 0xff, 0xc7, 0xdf, 0x0, 0xff, 0xcf, 0xdf, 0x0, 0xff, 0xd7, 0xdf, 0x0, 0xff, 0xdf, 0xdf, 0x0, 0xff, 0xe7, 0xdf, 0x0, 0xff, 0xef, 0xdf, 0x0, 0xff, 0xf7, 0xdf, 0x0, 0xff, 0x0, 0xe7, 0x0, 0xff, 0x7, 0xe7, 0x0, 0xff, 0xf, 0xe7, 0x0, 0xff, 0x17, 0xe7, 0x0, 0xff, 0x1f, 0xe7, 0x0, 0xff, 0x27, 0xe7, 0x0, 0xff, 0x2f, 0xe7, 0x0, 0xff, 0x37, 0xe7, 0x0, 0xff, 0x3f, 0xe7, 0x0, 0xff, 0x47, 0xe7, 0x0, 0xff, 0x4f, 0xe7, 0x0, 0xff, 0x57, 0xe7, 0x0, 0xff, 0x5f, 0xe7, 0x0, 0xff, 0x67, 0xe7, 0x0, 0xff, 0x6f, 0xe7, 0x0, 0xff, 0x77, 0xe7, 0x0, 0xff, 0x7f, 0xe7, 0x0, 0xff, 0x87, 0xe7, 0x0, 0xff, 0x8f, 0xe7, 0x0, 0xff, 0x97, 0xe7, 0x0, 0xff, 0x9f, 0xe7, 0x0, 0xff, 0xa7, 0xe7, 0x0, 0xff, 0xaf, 0xe7, 0x0, 0xff, 0xb7, 0xe7, 0x0, 0xff, 0xbf, 0xe7, 0x0, 0xff, 0xc7, 0xe7, 0x0, 0xff, 0xcf, 0xe7, 0x0, 0xff, 0xd7, 0xe7, 0x0, 0xff, 0xdf, 0xe7, 0x0, 0xff, 0xe7, 0xe7, 0x0, 0xff, 0xef, 0xe7, 0x0, 0xff, 0xf7, 0xe7, 0x0, 0xff, 0x0, 0xef, 0x0, 0xff, 0x7, 0xef, 0x0, 0xff, 0xf, 0xef, 0x0, 0xff, 0x17, 0xef, 0x0, 0xff, 0x1f, 0xef, 0x0, 0xff, 0x27, 0xef, 0x0, 0xff, 0x2f, 0xef, 0x0, 0xff, 0x37, 0xef, 0x0, 0xff, 0x3f, 0xef, 0x0, 0xff, 0x47, 0xef, 0x0, 0xff, 0x4f, 0xef, 0x0, 0xff, 0x57, 0xef, 0x0, 0xff, 0x5f, 0xef, 0x0, 0xff, 0x67, 0xef, 0x0, 0xff, 0x6f, 0xef, 0x0, 0xff, 0x77, 0xef, 0x0, 0xff, 0x7f, 0xef, 0x0, 0xff, 0x87, 0xef, 0x0, 0xff, 0x8f, 0xef, 0x0, 0xff, 0x97, 0xef, 0x0, 0xff, 0x9f, 0xef, 0x0, 0xff, 0xa7, 0xef, 0x0, 0xff, 0xaf, 0xef, 0x0, 0xff, 0xb7, 0xef, 0x0, 0xff, 0xbf, 0xef, 0x0, 0xff, 0xc7, 0xef, 0x0, 0xff, 0xcf, 0xef, 0x0, 0xff, 0xd7, 0xef, 0x0, 0xff, 0xdf, 0xef, 0x0, 0xff, 0xe7, 0xef, 0x0, 0xff, 0xef, 0xef, 0x0, 0xff, 0xf7, 0xef, 0x0, 0xff, 0x0, 0xf7, 0x0, 0xff, 0x7, 0xf7, 0x0, 0xff, 0xf, 0xf7, 0x0, 0xff, 0x17, 0xf7, 0x0, 0xff, 0x1f, 0xf7, 0x0, 0xff, 0x27, 0xf7, 0x0, 0xff, 0x2f, 0xf7, 0x0, 0xff, 0x37, 0xf7, 0x0, 0xff, 0x3f, 0xf7, 0x0, 0xff, 0x47, 0xf7, 0x0, 0xff, 0x4f, 0xf7, 0x0, 0xff, 0x57, 0xf7, 0x0, 0xff, 0x5f, 0xf7, 0x0, 0xff, 0x67, 0xf7, 0x0, 0xff, 0x6f, 0xf7, 0x0, 0xff, 0x77, 0xf7, 0x0, 0xff, 0x7f, 0xf7, 0x0, 0xff, 0x87, 0xf7, 0x0, 0xff, 0x8f, 0xf7, 0x0, 0xff, 0x97, 0xf7, 0x0, 0xff, 0x9f, 0xf7, 0x0, 0xff, 0xa7, 0xf7, 0x0, 0xff, 0xaf, 0xf7, 0x0, 0xff, 0xb7, 0xf7, 0x0, 0xff, 0xbf, 0xf7, 0x0, 0xff, 0xc7, 0xf7, 0x0, 0xff, 0xcf, 0xf7, 0x0, 0xff, 0xd7, 0xf7, 0x0, 0xff, 0xdf, 0xf7, 0x0, 0xff, 0xe7, 0xf7, 0x0, 0xff, 0xef, 0xf7, 0x0, 0xff, 0xf7, 0xf7, 0x0, 0xff}
	wantStride := 128
	wantRect := image.Rect(0, 0, 32, 32)

	assert.Equal(t, wantPix, img.Pix)
	assert.Equal(t, wantStride, img.Stride)
	assert.Equal(t, wantRect, img.Rect)
}
