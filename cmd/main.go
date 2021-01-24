package main

import (
	"log"

	"github.com/Henelik/tricaster/physics"
	"github.com/Henelik/tricaster/tuple"
)

func main() {
	physics_test()
}

func physics_test() {
	// projectile starts one unit above the origin
	// velocity is normalized to 1  unit/tick
	p, err := physics.CreateProjectile(
		tuple.NewPoint(0, 0, 1),
		tuple.NewVector(1, 0, 1),
	)
	if err != nil {
		log.Fatal(err)
	}

	// gravity -0.1 unit/tick, and wind is -0.01 unit/tick
	e, err := physics.CreateEnvironment(
		tuple.NewVector(0, 0, -0.1),
		tuple.NewVector(-0.01, 0, 0),
	)

	for p.Pos.Z > 0 {
		p = p.Tick(e)
		log.Printf("projectile position: X: %f, Y: %f, Z:%f, W:%f\n", p.Pos.X, p.Pos.Y, p.Pos.Z, p.Pos.W)
	}
}
