package shading

import (
	"github.com/Henelik/tricaster/tuple"
	"github.com/Henelik/tricaster/color"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestPhong(t *testing.T){
	testCases := []struct{
		name    string
		eyeV    *tuple.Tuple
		normalV *tuple.Tuple
		light   *PointLight
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
			want: color.NewColor(1.9, 1.9, 1.9),
		},
		{
			name:    "Lighting with the eye between light and surface, eye offset 45°",
			eyeV:    tuple.NewVector(0, math.Sqrt2/2, -math.Sqrt2/2),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 0, -10),
				color.White,
			},
			want: color.NewColor(1.0, 1.0, 1.0),
		},
		{
			name:    "Lighting with eye opposite surface, light offset 45°",
			eyeV:    tuple.NewVector(0, 0, -1),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 10, -10),
				color.White,
			},
			want: color.NewColor(0.7363961030678927,0.7363961030678927,0.7363961030678927),
		},
		{
			name:    "Lighting with eye in the path of the reflection vector",
			eyeV:    tuple.NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 10, -10),
				color.White,
			},
			want: color.NewColor(1.6363961030678928,1.6363961030678928,1.6363961030678928),
		},
		{
			name:    "Lighting with the light behind the surface",
			eyeV:    tuple.NewVector(0, 0, -1),
			normalV: tuple.NewVector(0, 0, -1),
			light: &PointLight{
				tuple.NewPoint(0, 0, 10),
				color.White,
			},
			want: color.NewColor(0.1, 0.1, 0.1),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DefaultPhong.Lighting(
				tc.light,
				tuple.Origin,
				tc.eyeV,
				tc.normalV)
			assert.True(t, tc.want.Equal(result))
			// assert.Equal(t, tc.want, result)
		})
	}
}