package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/tuple"
)

var DefaultShadeless = &ShadelessMat{
	Color: color.White,
}

type ShadelessMat struct {
	Color   *color.Color // used as a fallback if there is no pattern
	Pattern Pattern
}

func (m ShadelessMat) Lighting(light *PointLight, pos, eyeV, normalV *tuple.Tuple, inShadow bool) *color.Color {
	if m.Pattern != nil {
		return m.Pattern.Process(pos)
	}
	return m.Color
}

// CopyWithColor returns a new material with modified color
func (m *ShadelessMat) CopyWithColor(c *color.Color) *ShadelessMat {
	mat := *m
	mat.Color = c
	return &mat
}
