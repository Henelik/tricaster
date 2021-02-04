package shading

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/tuple"
)

func TestSolidPattern_Process(t *testing.T) {
	p := NewSolidPattern(color.White)
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
			want: color.White,
		},
		{
			name: "0, 1, 0",
			pos:  tuple.NewPoint(0, 1, 0),
			want: color.White,
		},
		{
			name: "1, 1, 0",
			pos:  tuple.NewPoint(1, 1, 0),
			want: color.White,
		},
		{
			name: "0, 0, 1",
			pos:  tuple.NewPoint(0, 0, 1),
			want: color.White,
		},
		{
			name: "1, 0, 1",
			pos:  tuple.NewPoint(1, 0, 1),
			want: color.White,
		},
		{
			name: "0, 1, 1",
			pos:  tuple.NewPoint(0, 1, 1),
			want: color.White,
		},
		{
			name: "1, 1, 1",
			pos:  tuple.NewPoint(1, 1, 1),
			want: color.White,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, p.Process(tc.pos))
		})
	}
}
