package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/util"
)

type StripePattern struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Patterns []Pattern
}

func NewStripePattern(m *matrix.Matrix, ps ...Pattern) *StripePattern {
	result := &StripePattern{
		m:        matrix.Identity,
		im:       matrix.Identity,
		Patterns: ps,
	}
	if m != nil {
		result.m = m
		result.im = m.Inverse()
	}
	return result
}

func (p *StripePattern) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *StripePattern) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *StripePattern) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(tpos.X)%len(p.Patterns))].Process(pos)
}
