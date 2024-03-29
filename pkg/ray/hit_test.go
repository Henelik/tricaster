package ray

import (
	"fmt"
	"testing"

	"github.com/Henelik/tricaster/pkg/color"
	"github.com/Henelik/tricaster/pkg/light"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

type TestPrimitive struct {
	IOR float64
}

func (prim *TestPrimitive) Intersects(r *Ray) []Intersection {
	return nil
}

func (prim *TestPrimitive) NormalAt(pos *tuple.Tuple) *tuple.Tuple {
	return tuple.NewVector(0, 0, -1)
}

func (prim *TestPrimitive) Shade(light *light.PointLight, h *Hit) *color.Color {
	return nil
}

func (prim *TestPrimitive) GetIOR() float64 {
	return prim.IOR
}

func TestRemovePrimitiveFromArr(t *testing.T) {
	s1 := &TestPrimitive{1}
	s2 := &TestPrimitive{1}
	s3 := &TestPrimitive{1}

	testCases := []struct {
		name     string
		item     IORHaver
		arr      []IORHaver
		wantArr  []IORHaver
		wantBool bool
	}{
		{
			name:     "empty array",
			item:     s1,
			arr:      []IORHaver{},
			wantArr:  []IORHaver{},
			wantBool: false,
		},
		{
			name:     "item not in array",
			item:     s1,
			arr:      []IORHaver{s2, s3},
			wantArr:  []IORHaver{s2, s3},
			wantBool: false,
		},
		{
			name:     "item at beginning of array",
			item:     s1,
			arr:      []IORHaver{s1, s2},
			wantArr:  []IORHaver{s2},
			wantBool: true,
		},
		{
			name:     "item at end of array",
			item:     s1,
			arr:      []IORHaver{s2, s1},
			wantArr:  []IORHaver{s2},
			wantBool: true,
		},
		{
			name:     "item in middle of array",
			item:     s1,
			arr:      []IORHaver{s2, s1, s3},
			wantArr:  []IORHaver{s2, s3},
			wantBool: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotArr, gotBool := removeIORHaverFromArr(tc.item, tc.arr)
			assert.Equal(t, tc.wantArr, gotArr)
			assert.Equal(t, tc.wantBool, gotBool)
		})
	}
}

func TestComputeRefractIOR(t *testing.T) {
	a := &TestPrimitive{1.5}
	b := &TestPrimitive{2.0}
	c := &TestPrimitive{2.5}

	r := NewRay(
		tuple.NewPoint(0, 0, -4),
		tuple.NewVector(0, 0, 1))

	xs := []Intersection{
		{2, a},
		{2.75, b},
		{3.25, c},
		{4.75, b},
		{5.25, c},
		{6, a},
	}

	testCases := []struct {
		index  int
		wantN1 float64
		wantN2 float64
	}{
		{
			index:  0,
			wantN1: 1.0,
			wantN2: 1.5,
		},
		{
			index:  1,
			wantN1: 1.5,
			wantN2: 2.0,
		},
		{
			index:  2,
			wantN1: 2.0,
			wantN2: 2.5,
		},
		{
			index:  3,
			wantN1: 2.5,
			wantN2: 2.5,
		},
		{
			index:  4,
			wantN1: 2.5,
			wantN2: 1.5,
		},
		{
			index:  5,
			wantN1: 1.5,
			wantN2: 1.0,
		},
	}
	for _, tc := range testCases {
		name := fmt.Sprintf("%v, %v, %v", tc.index, tc.wantN1, tc.wantN2)
		t.Run(name, func(t *testing.T) {
			h := NewHit(r, xs, tc.index)
			assert.Equal(t, tc.wantN1, h.N1)
			assert.Equal(t, tc.wantN2, h.N2)
		})
	}
}
