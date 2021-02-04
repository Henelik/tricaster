package shading

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Henelik/tricaster/tuple"

	"github.com/Henelik/tricaster/color"
)

func TestStripePattern_Process(t *testing.T) {
	p := NewStripePattern(nil, NewSolidPattern(color.White), NewSolidPattern(color.Black))
	testCases := []struct {
		name string
		pos  *tuple.Tuple
		want *color.Color
	}{
		{
			name: "A stripe pattern is constant in y",
			pos:  tuple.NewPoint(0, 2, 0),
			want: color.White,
		},
		{
			name: "A stripe pattern is constant in z",
			pos:  tuple.NewPoint(0, 0, 2),
			want: color.White,
		},
		{
			name: "A stripe pattern is white at 0 x",
			pos:  tuple.NewPoint(0, 0, 0),
			want: color.White,
		},
		{
			name: "A stripe pattern is black at 1 x",
			pos:  tuple.NewPoint(1, 0, 0),
			want: color.Black,
		},
		{
			name: "A stripe pattern alternates back to white at 2 x",
			pos:  tuple.NewPoint(2, 0, 0),
			want: color.White,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, p.Process(tc.pos))
		})
	}
}
