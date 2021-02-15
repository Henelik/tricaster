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
		Patterns: []Pattern{c1, c2},
	}
	if m != nil {
		result.SetMatrix(m)
	} else {
		result.SetMatrix(matrix.Identity)
	}
	return result
}

func (p *CheckerPattern3D) SetMatrix(m *matrix.Matrix) {
	p.m = m.Mult(matrix.Translation(100000, 100000, 100000)) // fix to break symmetry around pattern origin
	p.im = p.m.Inverse()
}

func (p *CheckerPattern3D) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *CheckerPattern3D) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(tpos.X)+int(tpos.Y)+int(tpos.Z))%2].Process(pos)
}
