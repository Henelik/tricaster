package renderer

import (
	"strconv"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/geometry"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/pattern"
	"github.com/Henelik/tricaster/pkg/tuple"
)

type Configuration struct {
	Name    string         `yaml:"name"`
	World   WorldConfig    `yaml:"world"`
	Camera  CameraConfig   `yaml:"camera"`
	Objects []ObjectConfig `yaml:"objects"`
}

// world

type WorldConfig struct {
	Shadows   bool        `yaml:"shadows"`
	MaxBounce int         `yaml:"max_bounce"`
	Light     LightConfig `yaml:"light"`
}

func (w *WorldConfig) ToWorld() *World {
	return &World{
		Light:  w.Light.ToLight(),
		Config: w,
	}
}

// light

type LightConfig struct {
	Color    ColorConfig
	Position PointConfig
}

func (l *LightConfig) ToLight() *light.PointLight {
	return &light.PointLight{
		Pos:   l.Position.ToPoint(),
		Color: l.Color.ToColor(),
	}
}

// camera

type CameraConfig struct {
	Height            int
	Width             int
	AALevel           int `yaml:"aa_level"`
	NumWorkers        int `yaml:"num_workers"`
	SubdivisionNumber int `yaml:"subdivision_number"`
	FOV               float64
	Transform         *ViewTransformConfig
}

func (c *CameraConfig) ToCamera() *Camera {
	return NewCamera(c)
}

type ViewTransformConfig struct {
	From PointConfig
	To   PointConfig
	Up   VectorConfig
}

func (v *ViewTransformConfig) ToMatrix() *matrix.Matrix {
	return matrix.ViewTransform(
		tuple.NewPoint(v.From.X, v.From.Y, v.From.Z),
		tuple.NewPoint(v.To.X, v.To.Y, v.To.Z),
		tuple.NewVector(v.Up.X, v.Up.Y, v.Up.Z),
	)
}

// object

type ObjectConfig struct {
	Type      string
	Material  MaterialConfig
	Transform TransformConfig
}

func (o *ObjectConfig) ToPrimitive() Primitive {
	switch o.Type {
	case "sphere":
		return geometry.NewSphere(o.Transform.ToMatrix(), o.Material.ToMaterial())
	case "cube":
		return geometry.NewCube(o.Transform.ToMatrix(), o.Material.ToMaterial())
	case "plane":
		return geometry.NewPlane(o.Transform.ToMatrix(), o.Material.ToMaterial())
	default:
		panic("unknown object type: " + o.Type)
	}
}

// material

type MaterialConfig struct {
	Type         string
	Ambient      float64
	Diffuse      float64
	Specular     float64
	Shininess    float64
	Reflectivity float64
	Transparency float64
	IOR          float64
	Color        ColorConfig
	Pattern      *PatternConfig
}

func (m *MaterialConfig) ToMaterial() material.Material {
	switch m.Type {
	case "phong":
		mat := &material.PhongMat{
			Ambient:      m.Ambient,
			Diffuse:      m.Diffuse,
			Specular:     m.Specular,
			Shininess:    m.Shininess,
			Reflectivity: m.Reflectivity,
			Transparency: m.Transparency,
			IOR:          m.IOR,
			Color:        m.Color.ToColor(),
			Pattern:      nil,
		}

		if m.Pattern != nil {
			mat.Pattern = m.Pattern.ToPattern()
		}

		return mat
	default:
		panic("unrecognized material type: " + m.Type)
	}
}

// transform

type TransformConfig struct {
	Position PointConfig
	Rotation PointConfig
	Scale    PointConfig
}

func (t *TransformConfig) ToMatrix() *matrix.Matrix {
	return matrix.Compose(
		matrix.Translation(t.Position.X, t.Position.Y, t.Position.Z),
		matrix.RotationX(t.Rotation.X),
		matrix.RotationY(t.Rotation.Y),
		matrix.RotationZ(t.Rotation.Z),
		matrix.Scaling(t.Scale.X, t.Scale.Y, t.Scale.Z),
	)
}

// tuples

type PointConfig struct {
	X float64
	Y float64
	Z float64
}

func (p *PointConfig) ToPoint() *tuple.Tuple {
	return tuple.NewPoint(p.X, p.Y, p.Z)
}

type VectorConfig struct {
	X float64
	Y float64
	Z float64
}

func (v *VectorConfig) ToVector() *tuple.Tuple {
	return tuple.NewVector(v.X, v.Y, v.Z)
}

// color

type ColorConfig struct {
	R float64
	G float64
	B float64
}

func (config *ColorConfig) ToColor() *color.Color {
	return color.NewColor(config.R, config.G, config.B)
}

// pattern

type PatternConfig struct {
	Type        string
	Color       ColorConfig
	Translation TransformConfig
	SubPatterns []PatternConfig `yaml:"sub_patterns"`
}

func (p *PatternConfig) ToPattern() pattern.Pattern {
	switch p.Type {
	case "solid":
		return pattern.NewSolidPattern(p.Color.ToColor())

	case "checker_2d":
		numSub := len(p.SubPatterns)

		if numSub < 2 {
			panic("not enough sub-patterns for checker_2d: " + strconv.Itoa(numSub))
		}

		return pattern.NewCheckerPattern2D(
			p.Translation.ToMatrix(),
			p.SubPatterns[0].ToPattern(),
			p.SubPatterns[1].ToPattern())

	case "checker_3d":
		numSub := len(p.SubPatterns)

		if numSub < 2 {
			panic("not enough sub-patterns for checker_3d: " + strconv.Itoa(numSub))
		}

		return pattern.NewCheckerPattern3D(
			p.Translation.ToMatrix(),
			p.SubPatterns[0].ToPattern(),
			p.SubPatterns[1].ToPattern())

	case "cylinder_ring":
		numSub := len(p.SubPatterns)

		if numSub < 2 {
			panic("not enough sub-patterns for cylinder_ring: " + strconv.Itoa(numSub))
		}

		subPatterns := make([]pattern.Pattern, 0, numSub)

		for _, pConfig := range p.SubPatterns {
			subPatterns = append(subPatterns, pConfig.ToPattern())
		}

		return pattern.NewCylinderRingPattern(
			p.Translation.ToMatrix(),
			subPatterns...)

	case "sphere_ring":
		numSub := len(p.SubPatterns)

		if numSub < 2 {
			panic("not enough sub-patterns for sphere_ring: " + strconv.Itoa(numSub))
		}

		subPatterns := make([]pattern.Pattern, 0, numSub)

		for _, pConfig := range p.SubPatterns {
			subPatterns = append(subPatterns, pConfig.ToPattern())
		}

		return pattern.NewCylinderRingPattern(
			p.Translation.ToMatrix(),
			subPatterns...)

	case "gradient":
		numSub := len(p.SubPatterns)

		if numSub < 2 {
			panic("not enough sub-patterns for gradient: " + strconv.Itoa(numSub))
		}

		return pattern.NewGradientPattern(
			p.Translation.ToMatrix(),
			p.SubPatterns[0].ToPattern(),
			p.SubPatterns[1].ToPattern())

	default:
		panic("unrecognized pattern type: " + p.Type)
	}
}