package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/ray"
)

type Material interface {
	Lighting(light *PointLight, h *ray.Hit) *color.Color
}
