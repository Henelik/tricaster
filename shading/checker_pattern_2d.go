package shading

import (
	"math"

	"github.com/Henelik/tricaster/util"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type CheckerPattern2D struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Patterns []Pattern
}

func NewCheckerPattern2D(m *matrix.Matrix, p1, p2 Pattern) *CheckerPattern2D {
	result := &CheckerPattern2D{
		m:        matrix.Identity,
		im:       matrix.Identity,
		Patterns: []Pattern{p1, p2},
	}
	if m != nil {
		result.m = m
		result.im = m.Inverse()
	}
	return result
}

func (p *CheckerPattern2D) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *CheckerPattern2D) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *CheckerPattern2D) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(math.Round(tpos.X))+int(math.Round(tpos.Y)))%2].Process(pos)
}
