package material

import (
	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/ray"
)

type Material interface {
	Lighting(light *light.PointLight, h *ray.Hit) *color.Color
	GetIOR() float64
}
