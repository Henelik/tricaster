package renderer

import (
	"github.com/Henelik/tricaster/color"
)

type Material interface {
	Lighting(light *PointLight, h *Hit) *color.Color
}
