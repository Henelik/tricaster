package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/util"
)

type CheckerPattern3D struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Patterns []Pattern
}

func NewCheckerPattern3D(m *matrix.Matrix, c1, c2 Pattern) *CheckerPattern3D {
	result := &CheckerPattern3D{
		m:        matrix.Identity,
		im:       matrix.Identity,
		Patterns: []Pattern{c1, c2},
	}
	if m != nil {
		result.m = m
		result.im = m.Inverse()
	}
	return result
}

func (p *CheckerPattern3D) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *CheckerPattern3D) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *CheckerPattern3D) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(tpos.X)+int(tpos.Y)+int(tpos.Z))%2].Process(pos)
}
