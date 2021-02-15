package shading

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/ray"
)

// Preset materials
var (
	Glass = &PhongMat{
		Ambient:      0,
		Diffuse:      0.1,
		Specular:     0.7,
		Shininess:    300,
		Reflectivity: 0.1,
		Transparency: 0.9,
		IOR:          1.5,
		Color:        color.NewColor(0.85, 0.9, 0.85),
		Pattern:      nil,
	}
)

type Material interface {
	Lighting(light *PointLight, h *ray.Hit) *color.Color
}
