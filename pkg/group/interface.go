package group

import (
	"github.com/Henelik/tricaster/pkg/ray"
	"github.com/Henelik/tricaster/pkg/tuple"
)

type Intersecter interface {
	Intersects(r *ray.Ray) []ray.Intersection
}

type GroupInterface interface {
	Intersects(r *ray.Ray) []ray.Intersection
	WorldToGroup(p *tuple.Tuple) *tuple.Tuple
	GroupToWorld(p *tuple.Tuple) *tuple.Tuple
}
