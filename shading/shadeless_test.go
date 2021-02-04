package shading

import (
	"math"
	"testing"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/tuple"
	"github.com/stretchr/testify/assert"
)

func TestShadeless(t *testing.T) {
	testCases := []struct {
		name    string
		eyeV    *tuple.Tuple
		normalV *tuple.Tuple
		light   *PointLight
		shadow  bool
		want    *color.Color
	}{
		{
			name:    "Lighting with the eye between the light and the surface",
			eyeV:    tuple.NewVector(0, 0, -1),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 0, -10),
				color.White,
			},
			want: color.White,
		},
		{
			name:    "Lighting with the eye between light and surface, eye offset 45°",
			eyeV:    tuple.NewVector(0, math.Sqrt2/2, -math.Sqrt2/2),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 0, -10),
				color.White,
			},
			want: color.White,
		},
		{
			name:    "Lighting with eye opposite surface, light offset 45°",
			eyeV:    tuple.NewVector(0, 0, -1),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 10, -10),
				color.White,
			},
			want: color.White,
		},
		{
			name:    "Lighting with eye in the path of the reflection vector",
			eyeV:    tuple.NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 10, -10),
				color.White,
			},
			want: color.White,
		},
		{
			name:    "Lighting with the light behind the surface",
			eyeV:    tuple.NewVector(0, 0, -1),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 0, 10),
				color.White,
			},
			want: color.White,
		},
		{
			name:    "Lighting with the surface in shadow",
			eyeV:    tuple.NewVector(0, 0, -1),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 0, 10),
				color.White,
			},
			shadow: true,
			want:   color.White,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DefaultShadeless.Lighting(
				tc.light,
				tuple.Origin,
				tc.eyeV,
				tc.normalV,
				tc.shadow)
			assert.True(t, tc.want.Equal(result))
		})
	}
}
