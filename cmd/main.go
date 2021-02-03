//+build !test

package main

import (
	"image/png"
	"log"
	"math"
	"os"

	"github.com/Henelik/tricaster/scene"
	"github.com/Henelik/tricaster/shading"

	"github.com/Henelik/tricaster/canvas"
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/geometry"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/physics"
	"github.com/Henelik/tricaster/tuple"
)

func main() {
	drawTestScene()
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
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.Red,
		})

	light := &shading.PointLight{
		tuple.NewPoint(10, -5, 10),
		color.White,
	}

	world := &scene.World{
		Geometry: []geometry.Primitive{s},
		Light:    light,
	}

	cam := scene.NewCamera(512, 512, 0,
		matrix.ViewTransform(
			tuple.Origin,
			tuple.NewPoint(0, 5, 0),
			tuple.Up), 16)

	cam.Render(world).SaveImage("new_sphere.png")
}

func drawTestScene() {
	floorMat := &shading.PhongMat{
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.0,
		Shininess: 10,
		Color:     color.NewColor(1, 0.9, 0.9),
		Pattern: shading.NewCheckerPattern3D(
			matrix.ScalingU(20),
			color.NewColor(1, 0.9, 0.9),
			color.NewColor(0.2, 0.19, 0.19)),
	}
	floorMat.Specular = 0
	floor := geometry.NewPlane(
		matrix.Identity,
		floorMat)

	middle := geometry.NewSphere(
		matrix.Translation(5, 5, 2).Mult(matrix.Scaling(2, 2, 2)),
		&shading.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.0,
			Shininess: 10,
			Color:     color.NewColor(0.1, 1, 0.5),
			Pattern: shading.NewCheckerPattern3D(
				matrix.Compose(
					matrix.RotationZ(-math.Pi/6),
					matrix.RotationY(-math.Pi/6),
					matrix.ScalingU(.5),
				),
				color.NewColor(0.1, 1, 0.5),
				color.NewColor(0.1, 0.5, 0.4)),
		})

	left := geometry.NewSphere(
		matrix.Translation(2, -2, 1),
		&shading.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(1, 0.1, 0.1),
		})

	right := geometry.NewSphere(
		matrix.Translation(-4, 3, 1.25).Mult(matrix.ScalingU(1.25)),
		&shading.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(0.2, 0.2, 1),
			Pattern: shading.NewStripePattern(
				matrix.Compose(
					matrix.Translation(0, 0, .25),
					matrix.RotationY(math.Pi/2),
					matrix.ScalingU(.5),
				),
				color.NewColor(0.2, 0.2, 1),
				color.NewColor(0.2, 0.2, 0.4)),
		})

	w := &scene.World{
		Geometry: []geometry.Primitive{
			floor,
			middle,
			left,
			right,
		},
		Light: &shading.PointLight{
			Pos:   tuple.NewPoint(0, -10, 10),
			Color: color.White,
		},
		Config: scene.WorldConfig{
			Shadows: true,
		},
	}

	c := scene.NewCamera(1000, 500, math.Pi/3,
		matrix.ViewTransform(
			tuple.NewPoint(-15, -10, 5),
			tuple.NewPoint(3, 3, 2),
			tuple.Up), 4)

	c.GoRender(2, w).SaveImage("new_scene_aax4.png")
}
