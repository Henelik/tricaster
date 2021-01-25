package canvas

import (
	"testing"

	"github.com/Henelik/tricaster/color"
	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	w := 10
	h := 20

	c, err := NewCanvas(w, h)
	assert.Equal(t, nil, err)

	assert.Equal(t, w, c.W)
	assert.Equal(t, h, c.H)

	assert.Equal(t, w*h, len(c.Pix))
	assert.Equal(t, color.NewColor(0, 0, 0), &c.Pix[0])
}

func TestGet(t *testing.T) {
	w := 10
	h := 20
	black := color.NewColor(0, 0, 0)

	c, err := NewCanvas(w, h)
	assert.Equal(t, nil, err)

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

	c, err := NewCanvas(w, h)
	assert.Equal(t, nil, err)

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
