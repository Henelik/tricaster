package material

import (
	"math"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/pattern"
	"github.com/Henelik/tricaster/pkg/ray"
)

var (
	DefaultPhong = &PhongMat{
		Ambient:      0.1,
		Diffuse:      0.9,
		Specular:     0.9,
		Shininess:    200,
		Reflectivity: 0,
		Transparency: 0,
		IOR:          1,
		Color:        color.White,
	}
	Glass = &PhongMat{
		Ambient:      0.1,
		Diffuse:      0.9,
		Specular:     0.8,
		Shininess:    300,
		Reflectivity: .9,
		Transparency: .9,
		IOR:          1.5,
		Color:        color.Grey(.9),
		Pattern:      nil,
	}
)

type PhongMat struct {
	Ambient      float64
	Diffuse      float64
	Specular     float64
	Shininess    float64
	Reflectivity float64
	Transparency float64
	IOR          float64
	Color        *color.Color // used as a fallback if there is no pattern
	Pattern      pattern.Pattern
}

func (m *PhongMat) Lighting(light *light.PointLight, h *ray.Hit) *color.Color {
	var col *color.Color
	if m.Pattern != nil {
		col = m.Pattern.Process(h.Pos)
	} else {
		col = m.Color
	}
	if h.InShadow {
		return col.MultCol(light.Color).MultF(m.Ambient)
	}
	effectiveColor := col.MultCol(light.Color)
	lightV := light.Pos.Sub(h.Pos).Norm()
	ambient := effectiveColor.MultF(m.Ambient)
	// light_dot_normal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the
	// light is on the other side of the surface.
	lightDotNormal := lightV.DotProd(h.NormalV)
	diffuse := color.Black
	specular := color.Black
	if lightDotNormal >= 0 {
		// compute the diffuse contribution
		diffuse = effectiveColor.MultF(m.Diffuse * lightDotNormal)

		if m.Specular != 0 {
			// reflect_dot_eye represents the cosine of the angle between the
			// reflection vector and the eye vector. A negative number means the
			// light reflects away from the eye.
			reflectDotEye := lightV.Neg().Reflect(h.NormalV).DotProd(h.EyeV)

			if reflectDotEye > 0 {
				// compute the specular reflection
				factor := math.Pow(reflectDotEye, m.Shininess)
				specular = light.Color.MultF(m.Specular * factor)
			}
		}
	}
	return ambient.Add(diffuse.Add(specular))
}

// Copy returns a new duplicate material
func (m *PhongMat) Copy() *PhongMat {
	mat := *m
	return &mat
}

// CopyWithColor returns a new material with modified color
func (m *PhongMat) CopyWithColor(c *color.Color) *PhongMat {
	mat := *m
	mat.Color = c
	return &mat
}

func (m *PhongMat) GetIOR() float64 {
	return m.IOR
}
