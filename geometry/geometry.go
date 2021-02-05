package geometry

import (
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/shading"
	"github.com/Henelik/tricaster/tuple"

	"github.com/Henelik/tricaster/ray"
)

// Primitive geometry type which defines an intersection function
type Primitive interface {
	SetMatrix(m *matrix.Matrix)
	GetMatrix() *matrix.Matrix
	// Intersects returns an array of intersections where the ray meets the primitive
	Intersects(r *ray.Ray) []ray.Intersection
	// NormalAt returns the normal vector at a given scene point
	NormalAt(pos *tuple.Tuple) *tuple.Tuple
	Shade(light *shading.PointLight, h *ray.Hit) *color.Color
}
