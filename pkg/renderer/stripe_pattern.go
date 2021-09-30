package renderer

import (
	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/pattern"
	"github.com/Henelik/tricaster/pkg/tuple"
	"github.com/Henelik/tricaster/pkg/util"
)

type StripePattern struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Patterns []pattern.Pattern
}

func NewStripePattern(m *matrix.Matrix, ps ...pattern.Pattern) *StripePattern {
	result := &StripePattern{
		Patterns: ps,
	}
	if m != nil {
		result.SetMatrix(m)
	} else {
		result.SetMatrix(matrix.Identity)
	}
	return result
}

func (p *StripePattern) SetMatrix(m *matrix.Matrix) {
	p.m = m.Mult(matrix.Translation(100000, 100000, 100000)) // fix to break symmetry around pattern origin
	p.im = p.m.Inverse()
}

func (p *StripePattern) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *StripePattern) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(tpos.X)%len(p.Patterns))].Process(pos)
}
