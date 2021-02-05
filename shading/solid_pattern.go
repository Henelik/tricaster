package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/tuple"
)

type SolidPattern struct {
	Color *color.Color
}

func NewSolidPattern(col *color.Color) *SolidPattern {
	return &SolidPattern{Color: col}
}

func SolidPat(r, g, b float32) *SolidPattern {
	return NewSolidPattern(color.NewColor(r, g, b))
}

func (p *SolidPattern) Process(pos *tuple.Tuple) *color.Color {
	return p.Color
}
