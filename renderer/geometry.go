package renderer

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/tuple"
)

// Primitive geometry type which defines an intersection function
type Primitive interface {
	SetMatrix(m *matrix.Matrix)
	GetMatrix() *matrix.Matrix
	// Intersects returns an array of intersections where the ray meets the primitive
	Intersects(r *Ray) []Intersection
	// NormalAt returns the normal vector at a given scene point
	NormalAt(pos *tuple.Tuple) *tuple.Tuple
	Shade(light *PointLight, h *Hit) *color.Color
	GetMaterial() Material
}
