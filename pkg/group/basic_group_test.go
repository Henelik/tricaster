package group

import (
	"testing"

	"github.com/Henelik/tricaster/pkg/matrix"

	"github.com/Henelik/tricaster/pkg/geometry"

	"github.com/stretchr/testify/assert"

	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
)

func TestBasicGroup_Intersects(t *testing.T) {
	s1 := geometry.NewSphere(nil, nil)
	s2 := geometry.NewSphere(matrix.Translation(0, 0, -3), nil)
	s3 := geometry.NewSphere(matrix.Translation(5, 0, 0), nil)

	testCases := []struct {
		name  string
		group *BasicGroup
		ray   *ray.Ray
		want  []ray.Intersection
	}{
		{
			name:  "Intersecting a ray with an empty group",
			group: NewBasicGroup(nil, nil),
			ray: ray.NewRay(
				tuple.NewPoint(0, 0, 0),
				tuple.NewVector(0, 0, 1),
			),
			want: []ray.Intersection{},
		},
		{
			name:  "Intersecting a ray with a nonempty group",
			group: NewBasicGroup(nil, nil, s1, s2, s3),
			ray: ray.NewRay(
				tuple.NewPoint(0, 0, -5),
				tuple.NewVector(0, 0, 1),
			),
			want: []ray.Intersection{
				{4, s1},
				{6, s1},
				{1, s2},
				{3, s2},
			},
		},
		{
			name:  "Intersecting a transformed group",
			group: NewBasicGroup(matrix.Scaling(2, 2, 2), nil, s3),
			ray: ray.NewRay(
				tuple.NewPoint(10, 0, -10),
				tuple.NewVector(0, 0, 1),
			),
			want: []ray.Intersection{
				{8, s3},
				{12, s3},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.group.Intersects(tc.ray)

			assert.Equal(t, tc.want, got)
		})
	}
}
