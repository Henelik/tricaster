package main

import (
	"github.com/Henelik/tricaster/shading"
	"image/png"
	"log"
	"os"

	"github.com/Henelik/tricaster/canvas"
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/geometry"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/physics"
	"github.com/Henelik/tricaster/ray"
	"github.com/Henelik/tricaster/tuple"
)

func main() {
	drawSphereTest()
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
	if err != nil {
		log.Fatal(err)
	}

	for p.Pos.Z > 0 {
		p = p.Tick(e)
		log.Printf("projectile position: %s\n", p.Pos.Fmt())
	}
}

// generate a test UV space
func imageTest() {
	canv := canvas.NewCanvas(256, 256)

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

func projectilePlot() {
	p, err := physics.NewProjectile(
		tuple.NewPoint(0, 0, 1),
		tuple.NewVector(1, 0, 1.8).Norm().Mult(11.25),
	)
	if err != nil {
		log.Fatal(err)
	}

	e, err := physics.NewEnvironment(
		tuple.NewVector(0, 0, -0.1),
		tuple.NewVector(-0.01, 0, 0),
	)
	if err != nil {
		log.Fatal(err)
	}

	canv := canvas.NewCanvas(900, 550)

	for p.Pos.Z > 0 {
		p = p.Tick(e)
		canv.SetSafe(int(p.Pos.X), canv.H-int(p.Pos.Z), color.Red)
		canv.SetSafe(int(p.Pos.X)+1, canv.H-int(p.Pos.Z), color.Red)
		canv.SetSafe(int(p.Pos.X), canv.H-int(p.Pos.Z)+1, color.Red)
		canv.SetSafe(int(p.Pos.X)+1, canv.H-int(p.Pos.Z)+1, color.Red)
	}

	img := canv.ToImage()

	outputFile, err := os.Create("projectile.png")
	if err != nil {
		panic(err)
	}

	png.Encode(outputFile, img)

	outputFile.Close()
}

func drawSphereTest() {
	s := geometry.NewSphere(
		matrix.Translation(0, 5, 0),
		&shading.PhongMat{
			Ambient: 0.1,
			Diffuse: 0.9,
			Specular: 0.9,
			Shininess: 200,
			Color: color.Red,
		})

	light := &shading.PointLight{
		tuple.NewPoint(-10, -5, 10),
		color.White,
	}

	w := 512
	h := 512

	canv := canvas.NewCanvas(w, h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			// shoot a ray at the sphere
			xPos := (float64(x) - float64(w)/2.0) / (float64(w) * 0.4)
			yPos := -(float64(y) - float64(h)/2.0) / (float64(h) * 0.4)
			r := ray.NewRay(
				tuple.NewPoint(xPos, 0, yPos),
				tuple.NewVector(0, 1, 0),
			)
			h := geometry.Hit(s.Intersects(r))
			if h != geometry.NilHit {
				// canv.Set(x, y, color.Red)
				hitPoint := r.Position(h.T)
				c := s.Mat.Lighting(
					light,
					hitPoint,
					r.Direction.Neg(),
					s.NormalAt(hitPoint))
				// fmt.Printf("col: (%v, %v, %v)\n", c.R, c.G, c.B)
				canv.Set(x, y, c)
			}
		}
	}

	img := canv.ToImage()

	outputFile, err := os.Create("sphere.png")
	if err != nil {
		panic(err)
	}

	png.Encode(outputFile, img)

	outputFile.Close()
}
