package pattern

import (
	"testing"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
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
