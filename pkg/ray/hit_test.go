package ray

import (
	"fmt"
	"testing"

	"github.com/Henelik/tricaster/pkg/geometry"
	"github.com/Henelik/tricaster/pkg/material"
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/tuple"

	"github.com/stretchr/testify/assert"
)

func TestRemovePrimitiveFromArr(t *testing.T) {
	s1 := geometry.NewSphere(nil, nil)
	s2 := geometry.NewSphere(nil, nil)
	s3 := geometry.NewSphere(nil, nil)

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
	a := geometry.NewSphere(
		matrix.ScalingU(2),
		material.Glass.Copy())
	a.Mat.(*material.PhongMat).IOR = 1.5

	b := geometry.NewSphere(
		matrix.Translation(0, 0, -0.25),
		material.Glass.Copy())
	b.Mat.(*material.PhongMat).IOR = 2.0

	c := geometry.NewSphere(
		matrix.Translation(0, 0, 0.25),
		material.Glass.Copy())
	c.Mat.(*material.PhongMat).IOR = 2.5

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
		name := fmt.Sprintf("%v,%v,%v", tc.index, tc.wantN1, tc.wantN2)
		t.Run(name, func(t *testing.T) {
			h := xs[tc.index].ToHit(r, xs)
			assert.Equal(t, tc.wantN1, h.N1)
			assert.Equal(t, tc.wantN2, h.N2)
		})
	}
}
