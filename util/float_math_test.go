package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	var f1 float32 = 1.0 + .5*Epsilon
	var f2 float32 = 1.0

	assert.True(t, Equal(f1, f2))
	assert.NotEqual(t, f1, f2)

	f1 = 1.0 + 2*Epsilon

	assert.False(t, Equal(f1, f2))
	assert.NotEqual(t, f1, f2)
}

func TestClamp(t *testing.T) {
	testCases := []struct {
		name string
		n    float32
		min  float32
		max  float32
		want float32
	}{
		{
			name: "unclamped return",
			n:    0.5,
			min:  0,
			max:  1,
			want: 0.5,
		},
		{
			name: "min return return",
			n:    -0.5,
			min:  0,
			max:  1,
			want: 0,
		},
		{
			name: "max return",
			n:    1.5,
			min:  0,
			max:  1,
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Clamp(tc.n, tc.min, tc.max)
			assert.Equal(t, tc.want, got)
		})
	}
}
