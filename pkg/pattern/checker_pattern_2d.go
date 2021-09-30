package pattern

import (
	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/tuple"
	"github.com/Henelik/tricaster/pkg/util"
)

type CheckerPattern2D struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Patterns []Pattern
}

func NewCheckerPattern2D(m *matrix.Matrix, p1, p2 Pattern) *CheckerPattern2D {
	result := &CheckerPattern2D{
		Patterns: []Pattern{p1, p2},
	}
	if m != nil {
		result.SetMatrix(m)
	} else {
		result.SetMatrix(matrix.Identity)
	}
	return result
}

func (p *CheckerPattern2D) SetMatrix(m *matrix.Matrix) {
	p.m = m.Mult(matrix.Translation(100000, 100000, 100000)) // fix to break symmetry around pattern origin
	p.im = p.m.Inverse()
}

func (p *CheckerPattern2D) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *CheckerPattern2D) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(tpos.X)+int(tpos.Y))%2].Process(pos)
}
