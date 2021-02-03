package shading

import (
	"math"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/util"
)

type CheckerPattern3D struct {
	m      *matrix.Matrix
	im     *matrix.Matrix
	Colors []*color.Color
}

func NewCheckerPattern3D(m *matrix.Matrix, c1, c2 *color.Color) *CheckerPattern3D {
	result := &CheckerPattern3D{
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

func (p *CheckerPattern3D) SetMatrix(m *matrix.Matrix) {
	p.m = m
	p.im = m.Inverse()
}

func (p *CheckerPattern3D) GetMatrix() *matrix.Matrix {
	return p.m
}

func (p *CheckerPattern3D) Process(pos *tuple.Tuple) *color.Color {
	tpos := p.im.MultTuple(pos)
	return p.Colors[util.AbsInt(int(math.Round(tpos.X))+int(math.Round(tpos.Y))+int(math.Round(tpos.Z)))%2]
}