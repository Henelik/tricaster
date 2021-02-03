package shading

import (
	"math"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type GradientPattern struct {
	m      *matrix.Matrix
	im     *matrix.Matrix
	Color1 *color.Color
	Color2 *color.Color
}

func NewGradientPattern(m *matrix.Matrix, c1, c2 *color.Color) *GradientPattern {
	result := &GradientPattern{
		m:      matrix.Identity,
		im:     matrix.Identity,
		Color1: c1,
		Color2: c2,
	}
	if m != nil {
		result.m = m
		result.im = m.Inverse()
	}
	return result
}

func (p *GradientPattern) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *GradientPattern) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *GradientPattern) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Color1.Lerp(p.Color2, tpos.X-math.Floor(tpos.X))
}
