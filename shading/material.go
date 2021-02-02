package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/tuple"
)

type Material interface {
	Lighting(light *PointLight, pos, eyeV, normalV *tuple.Tuple, inShadow bool) *color.Color
}
