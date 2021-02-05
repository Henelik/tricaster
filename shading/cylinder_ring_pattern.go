package shading

import (
	"git.maze.io/go/math32"
	"github.com/Henelik/tricaster/util"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

type CylinderRingPattern struct {
	m        *matrix.Matrix
	im       *matrix.Matrix
	Patterns []Pattern
}

func NewCylinderRingPattern(m *matrix.Matrix, ps ...Pattern) *CylinderRingPattern {
	result := &CylinderRingPattern{
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

func (p *CylinderRingPattern) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *CylinderRingPattern) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *CylinderRingPattern) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Patterns[util.AbsInt(int(math32.Sqrt(tpos.X*tpos.X+tpos.Y*tpos.Y)))%len(p.Patterns)].Process(pos)
}
