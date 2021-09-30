package pattern

import (
	"testing"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestGradientPattern_Process(t *testing.T) {
	// A gradient linearly interpolates between colors
	p := NewGradientPattern(nil, NewSolidPattern(color.White), NewSolidPattern(color.Black))
	assert.Equal(t, color.White, p.Process(tuple.NewPoint(0, 0, 0)))
	assert.Equal(t, color.Grey(0.75), p.Process(tuple.NewPoint(0.25, 0, 0)))
	assert.Equal(t, color.Grey(0.5), p.Process(tuple.NewPoint(0.5, 0, 0)))
	assert.Equal(t, color.Grey(0.25), p.Process(tuple.NewPoint(0.75, 0, 0)))
}
