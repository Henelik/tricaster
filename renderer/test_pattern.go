package renderer

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type TestPattern struct {
	m  *matrix.Matrix
	im *matrix.Matrix
}

func NewTestPattern(m *matrix.Matrix) *TestPattern {
	result := &TestPattern{
		m:  matrix.Identity,
		im: matrix.Identity,
	}
	if m != nil {
		result.m = m
		result.im = m.Inverse()
	}
	return result
}

func (p *TestPattern) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *TestPattern) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *TestPattern) Process(pos *tuple.Tuple) *color.Color {
	return color.NewColor(pos.X, pos.Y, pos.Z)
}
