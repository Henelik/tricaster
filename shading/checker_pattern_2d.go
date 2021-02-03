package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/util"
)

type CheckerPattern2D struct {
	m      *matrix.Matrix
	im     *matrix.Matrix
	Colors []*color.Color
}

func NewCheckerPattern2D(m *matrix.Matrix, c1, c2 *color.Color) *CheckerPattern2D {
	result := &CheckerPattern2D{
		m:      matrix.Identity,
		im:     matrix.Identity,
		Colors: []*color.Color{c1, c2},
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
	return p.Colors[util.AbsInt(int(tpos.X)+int(tpos.Y))%2]
}
