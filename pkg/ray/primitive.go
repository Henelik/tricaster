package ray

import (
	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/tuple"
)

// Primitive geometry type which defines an intersection function
type Primitive interface {
	// Intersects returns an array of intersections where the ray meets the primitive
	Intersects(r *Ray) []Intersection
	// NormalAt returns the normal vector at a given scene point
	NormalAt(pos *tuple.Tuple) *tuple.Tuple
	Shade(light *light.PointLight, h *Hit) *color.Color
}
