package renderer

import (
	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
)

// Primitive geometry type which defines an intersection function
type Primitive interface {
	SetMatrix(m *matrix.Matrix)
	GetMatrix() *matrix.Matrix
	// Intersects returns an array of intersections where the ray meets the primitive
	Intersects(r *ray.Ray) []ray.Intersection
	// NormalAt returns the normal vector at a given scene point
	NormalAt(pos *tuple.Tuple) *tuple.Tuple
	Shade(light *light.PointLight, h *ray.Hit) *color.Color
	GetMaterial() material.Material
	GetIOR() float64
}
