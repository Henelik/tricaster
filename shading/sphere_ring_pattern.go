package shading

import (
	"math"

	"github.com/Henelik/tricaster/util"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type SphereRingPattern struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Patterns []Pattern
}

func NewSphereRingPattern(m *matrix.Matrix, ps ...Pattern) *SphereRingPattern {
	result := &SphereRingPattern{
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

func (p *SphereRingPattern) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *SphereRingPattern) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *SphereRingPattern) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(math.Sqrt(tpos.X*tpos.X+tpos.Y*tpos.Y+tpos.Z*tpos.Z)))%len(p.Patterns)].Process(pos)
}
