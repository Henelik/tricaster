package shading

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/tuple"
)

func TestCylinderRingPattern_Process(t *testing.T) {
	p := NewCylinderRingPattern(
		nil,
		NewSolidPattern(color.White),
		NewSolidPattern(color.Black))
	testCases := []struct {
		name string
		pos  *tuple.Tuple
		want *color.Color
	}{
		{
			name: "0, 0, 0",
			pos:  tuple.NewPoint(0, 0, 0),
			want: color.White,
		},
		{
			name: "1, 0, 0",
			pos:  tuple.NewPoint(1, 0, 0),
			want: color.Black,
		},
		{
			name: "-1, 0, 0",
			pos:  tuple.NewPoint(-1, 0, 0),
			want: color.Black,
		},
		{
			name: "0, 1, 0",
			pos:  tuple.NewPoint(0, 1, 0),
			want: color.Black,
		},
		{
			name: "0, -1, 0",
			pos:  tuple.NewPoint(0, -1, 0),
			want: color.Black,
		},
		{
			name: "0, 0, 1",
			pos:  tuple.NewPoint(0, 0, 1),
			want: color.White,
		},
		{
			name: "0, 0, -1",
			pos:  tuple.NewPoint(0, 0, -1),
			want: color.White,
		},
		{
			name: "1, 1, 0",
			pos:  tuple.NewPoint(1, 1, 0),
			want: color.Black,
		},
		{
			name: "2, 2, 0",
			pos:  tuple.NewPoint(2, 2, 0),
			want: color.White,
		},
		{
			name: "1.25, 1.25, 1.25",
			pos:  tuple.NewPoint(1.25, 1.25, 1.25),
			want: color.Black,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, p.Process(tc.pos))
		})
	}
}
