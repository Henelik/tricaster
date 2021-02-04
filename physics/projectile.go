//+build !test

package physics

import (
	"errors"

	"github.com/Henelik/tricaster/tuple"
)

type Projectile struct {
	Pos *tuple.Tuple
	Vel *tuple.Tuple
}

func NewProjectile(pos, vel *tuple.Tuple) (*Projectile, error) {
	if !pos.IsPoint() {
		return nil, errors.New("pos must be a point")
	}
	if !vel.IsVector() {
		return nil, errors.New("vel must be a vector")
	}
	return &Projectile{pos, vel}, nil
}

func (p *Projectile) Tick(env *Environment) *Projectile {
	return &Projectile{
		Pos: p.Pos.Add(p.Vel),
		Vel: p.Vel.Add(env.Gravity).Add(env.Wind),
	}
}
