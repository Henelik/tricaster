package group

import (
	"github.com/Henelik/tricaster/pkg/matrix"
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
)

type BasicGroup struct {
	Parent                 GroupInterface
	Children               []Intersecter
	matrix                 *matrix.Matrix
	inverseMatrix          *matrix.Matrix
	inverseTransposeMatrix *matrix.Matrix
}

func NewBasicGroup(m *matrix.Matrix, parent GroupInterface, children ...Intersecter) *BasicGroup {
	group := &BasicGroup{
		Parent:   parent,
		Children: children,
		matrix:   matrix.Identity,
	}

	if m != nil {
		group.matrix = m
	}

	return group
}

func (group *BasicGroup) SetMatrix(m *matrix.Matrix) {
	group.matrix = m
	group.inverseMatrix = m.Inverse()
	group.inverseTransposeMatrix = group.inverseMatrix.Transpose()
}

func (group *BasicGroup) WorldToGroup(p *tuple.Tuple) *tuple.Tuple {
	if group.Parent != nil {
		return group.inverseMatrix.MultTuple(group.Parent.WorldToGroup(p))
	}

	return group.inverseMatrix.MultTuple(p)
}

func (group *BasicGroup) GroupToWorld(p *tuple.Tuple) *tuple.Tuple {
	if group.Parent != nil {
		return group.inverseTransposeMatrix.MultTuple(group.Parent.WorldToGroup(p))
	}

	return group.inverseTransposeMatrix.MultTuple(p)
}

func (group *BasicGroup) Intersects(r *ray.Ray) []ray.Intersection {
	rt := r.Transform(group.matrix.Inverse())

	result := make([]ray.Intersection, 0, len(group.Children)*2)

	for _, child := range group.Children {
		result = append(result, child.Intersects(rt)...)
	}

	return result
}
