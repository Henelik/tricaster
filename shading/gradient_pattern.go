package shading

import (
	"git.maze.io/go/math32"
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type GradientPattern struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Pattern1 Pattern
	Pattern2 Pattern
}

func NewGradientPattern(m *matrix.Matrix, c1, c2 Pattern) *GradientPattern {
	result := &GradientPattern{
		m:        matrix.Identity,
		im:       matrix.Identity,
		Pattern1: c1,
		Pattern2: c2,
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
	return p.Pattern1.Process(pos).Lerp(p.Pattern2.Process(pos), tpos.X-math32.Floor(tpos.X))
}
