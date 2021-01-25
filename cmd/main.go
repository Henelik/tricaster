package main

import (
	"image/png"
	"log"
	"os"

	"github.com/Henelik/tricaster/canvas"
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/physics"
	"github.com/Henelik/tricaster/tuple"
)

func main() {
	imageTest()
}

func physicsTest() {
	// projectile starts one unit above the origin
	// velocity is normalized to 1  unit/tick
	p, err := physics.NewProjectile(
		tuple.NewPoint(0, 0, 1),
		tuple.NewVector(1, 0, 1),
	)
	if err != nil {
		log.Fatal(err)
	}

	// gravity -0.1 unit/tick, and wind is -0.01 unit/tick
	e, err := physics.NewEnvironment(
		tuple.NewVector(0, 0, -0.1),
		tuple.NewVector(-0.01, 0, 0),
	)

	for p.Pos.Z > 0 {
		p = p.Tick(e)
		log.Printf("projectile position: %s\n", p.Pos.Fmt())
	}
}

// generate a test UV space
func imageTest() {
	canv, err := canvas.NewCanvas(32, 32)
	if err != nil {
		log.Fatal(err)
	}

	for x := 0; x < canv.W; x++ {
		for y := 0; y < canv.H; y++ {
			col := color.NewColor(
				float64(x)/float64(canv.W),
				float64(y)/float64(canv.H),
				0)
			canv.Set(x, y, col)
		}
	}

	img := canv.ToImage()

	outputFile, err := os.Create("test.png")
	if err != nil {
		panic(err)
	}

	png.Encode(outputFile, img)

	outputFile.Close()
}
