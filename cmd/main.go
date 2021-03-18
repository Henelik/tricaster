//+build !test

package main

import (
	"image/png"
	"log"
	"math"
	"os"
	"time"

	"github.com/Henelik/tricaster/canvas"
	"github.com/Henelik/tricaster/color"
	"github.com/Henelik/tricaster/matrix"
	"github.com/Henelik/tricaster/physics"
	"github.com/Henelik/tricaster/renderer"
	"github.com/Henelik/tricaster/tuple"
)

func main() {
	start := time.Now()

	// reflectionScene()
	refractionScene()

	t := time.Now()
	log.Printf("Render time was %v\n", t.Sub(start))
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
	s := renderer.NewSphere(
		matrix.Translation(0, 5, 0),
		&renderer.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.Red,
		})

	light := &renderer.PointLight{
		tuple.NewPoint(10, -5, 10),
		color.White,
	}

	world := &renderer.World{
		Geometry: []renderer.Primitive{s},
		Light:    light,
	}

	cam := renderer.NewCamera(512, 512, 0,
		matrix.ViewTransform(
			tuple.Origin,
			tuple.NewPoint(0, 5, 0),
			tuple.Up), 16)

	cam.Render(world).SaveImage("new_sphere.png")
}

func RGBSphereScene() {
	floorMat := &renderer.PhongMat{
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.0,
		Shininess: 10,
		Color:     color.NewColor(1, 0.9, 0.9),
		Pattern: renderer.NewCheckerPattern3D(
			matrix.ScalingU(20),
			renderer.NewStripePattern(
				matrix.RotationZ(math.Pi/4),
				renderer.SolidPat(1, 0.9, 0.9),
				renderer.SolidPat(.75, 0.7, 0.7),
			),
			renderer.NewStripePattern(
				matrix.RotationZ(-math.Pi/4),
				renderer.SolidPat(0.2, 0.19, 0.19),
				renderer.SolidPat(.3, 0.3, 0.3),
			)),
	}
	floor := renderer.NewPlane(
		matrix.Identity,
		floorMat)

	middle := renderer.NewSphere(
		matrix.Translation(5, 5, 2).Mult(matrix.Scaling(2, 2, 2)),
		&renderer.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.0,
			Shininess: 10,
			Color:     color.NewColor(0.1, 1, 0.5),
			Pattern: renderer.NewCheckerPattern3D(
				matrix.Compose(
					matrix.RotationZ(-math.Pi/6),
					matrix.RotationY(-math.Pi/6),
					matrix.ScalingU(.5),
				),
				renderer.SolidPat(0.1, 1, 0.5),
				renderer.SolidPat(0.1, 0.5, 0.4)),
		})

	left := renderer.NewSphere(
		matrix.Translation(2, -2, 1),
		&renderer.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(1, 0.1, 0.1),
		})

	right := renderer.NewSphere(
		matrix.Translation(-4, 3, 1.25).Mult(matrix.ScalingU(1.25)),
		&renderer.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			Color:     color.NewColor(0.2, 0.2, 1),
			Pattern: renderer.NewStripePattern(
				matrix.Compose(
					matrix.Translation(0, 0, .25),
					matrix.RotationY(math.Pi/2),
					matrix.ScalingU(.5),
				),
				renderer.NewGradientPattern(
					matrix.Compose(
						matrix.RotationY(math.Pi/2),
						matrix.ScalingU(3),
					),
					renderer.SolidPat(0.9, 0.9, 0.9),
					renderer.SolidPat(0.2, 0.2, 1),
				),
				renderer.SolidPat(0.2, 0.2, 0.4)),
		})

	w := &renderer.World{
		Geometry: []renderer.Primitive{
			floor,
			middle,
			left,
			right,
		},
		Light: &renderer.PointLight{
			Pos:   tuple.NewPoint(0, -10, 10),
			Color: color.White,
		},
		Config: renderer.WorldConfig{
			Shadows: true,
		},
	}

	c := renderer.NewCamera(1920, 1080, math.Pi/3,
		matrix.ViewTransform(
			tuple.NewPoint(-15, -10, 5),
			tuple.NewPoint(3, 3, 2),
			tuple.Up), 0)

	c.GoRender(w, 2).SaveImage("new_scene_aax4.png")
}

func reflectionScene() {
	floorMat := &renderer.PhongMat{
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.0,
		Shininess: 10,
		// Reflectivity: .01,
		Color: color.NewColor(1, 0.9, 0.9),
		Pattern: renderer.NewCheckerPattern3D(
			matrix.ScalingU(5).Mult(
				matrix.Translation(2.5, 2.5, 2.5)),
			renderer.SolidPat(1, 0.9, 0.9),
			renderer.SolidPat(0.2, 0.19, 0.19),
		),
	}
	floor := renderer.NewPlane(
		matrix.Identity,
		floorMat)
	lWall := renderer.NewPlane(
		matrix.Compose(
			matrix.Translation(20, 0, 0),
			matrix.RotationY(math.Pi/2),
		),
		floorMat)
	rWall := renderer.NewPlane(
		matrix.Compose(
			matrix.Translation(0, 20, 0),
			matrix.RotationX(math.Pi/2),
		),
		floorMat)
	blWall := renderer.NewPlane(
		matrix.Compose(
			matrix.Translation(-20, 0, 0),
			matrix.RotationY(math.Pi/2),
		),
		floorMat)
	brWall := renderer.NewPlane(
		matrix.Compose(
			matrix.Translation(0, -20, 0),
			matrix.RotationX(math.Pi/2),
		),
		floorMat)
	ceiling := renderer.NewPlane(
		matrix.Translation(0, 0, 40),
		floorMat)

	mirrorBall := renderer.NewSphere(
		matrix.Translation(5, 5, 2).Mult(matrix.Scaling(2, 2, 2)),
		&renderer.PhongMat{
			Ambient:      0.1,
			Diffuse:      0.9,
			Specular:     0.8,
			Shininess:    300,
			Reflectivity: .9,
			Color:        color.NewColor(0.7, 0.7, 0.7),
			Pattern:      nil,
		})

	middle := renderer.NewSphere(
		matrix.Translation(-10, -10, 2).Mult(matrix.Scaling(2, 2, 2)),
		&renderer.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.0,
			Shininess: 10,
			// Reflectivity: .01,
			Color: color.NewColor(0.1, 1, 0.5),
			Pattern: renderer.NewCheckerPattern3D(
				matrix.Compose(
					matrix.RotationZ(-math.Pi/6),
					matrix.RotationY(-math.Pi/6),
					matrix.ScalingU(.5),
				),
				renderer.SolidPat(0.1, 1, 0.5),
				renderer.SolidPat(0.1, 0.5, 0.4)),
		})
	left := renderer.NewSphere(
		matrix.Translation(7, -7, 1),
		&renderer.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			// Reflectivity: .05,
			Color: color.NewColor(1, 0.1, 0.1),
		})
	right := renderer.NewSphere(
		matrix.Translation(-4, 3, 1.25).Mult(matrix.ScalingU(1.25)),
		&renderer.PhongMat{
			Ambient:   0.1,
			Diffuse:   0.9,
			Specular:  0.9,
			Shininess: 200,
			// Reflectivity: .025,
			Color: color.NewColor(0.2, 0.2, 1),
			Pattern: renderer.NewStripePattern(
				matrix.Compose(
					matrix.Translation(0, 0, .25),
					matrix.RotationY(math.Pi/2),
					matrix.ScalingU(.5),
				),
				renderer.NewGradientPattern(
					matrix.Compose(
						matrix.RotationY(math.Pi/2),
						matrix.ScalingU(3),
					),
					renderer.SolidPat(0.9, 0.9, 0.9),
					renderer.SolidPat(0.2, 0.2, 1),
				),
				renderer.SolidPat(0.2, 0.2, 0.4)),
		})

	w := &renderer.World{
		Geometry: []renderer.Primitive{
			floor,
			ceiling,
			lWall,
			rWall,
			blWall,
			brWall,
			mirrorBall,
			middle,
			right,
			left,
		},
		Light: &renderer.PointLight{
			Pos:   tuple.NewPoint(0, -10, 10),
			Color: color.White,
		},
		Config: renderer.WorldConfig{
			Shadows:   true,
			MaxBounce: 3,
		},
	}

	c := renderer.NewCamera(1920, 1080, math.Pi/3,
		matrix.ViewTransform(
			tuple.NewPoint(-15, -10, 5),
			tuple.NewPoint(3, 3, 2),
			tuple.Up), 0)

	c.GoRender(w, 8).SaveImage("reflection.png")
}

func refractionScene() {
	roomMat := &renderer.PhongMat{
		Ambient: 0.1,
		Diffuse: 0.9,
		IOR:     1,
		Color:   color.NewColor(1, 0.9, 0.9),
		Pattern: renderer.NewCheckerPattern3D(
			matrix.ScalingU(1).Mult(
				matrix.Translation(2.5, 2.5, 2.5)),
			renderer.SolidPat(0.91, 0.9, 0.9),
			renderer.SolidPat(0.2, 0.19, 0.19),
		),
	}
	room := renderer.NewCube(
		matrix.Compose(
			matrix.Translation(0, 0, 20),
			matrix.ScalingU(20)),
		roomMat,
	)

	glassBall := renderer.NewSphere(
		matrix.Translation(0, 0, 3).Mult(matrix.Scaling(2, 2, 2)),
		&renderer.PhongMat{
			Ambient:      0.1,
			Diffuse:      0.9,
			Specular:     0.8,
			Shininess:    300,
			Reflectivity: .9,
			Transparency: .9,
			IOR:          1.5,
			Color:        color.Grey(.1),
			Pattern:      nil,
		})

	airBall := renderer.NewSphere(
		matrix.Translation(0, 0, 3),
		&renderer.PhongMat{
			Ambient:      0.1,
			Diffuse:      0.9,
			Specular:     0.8,
			Shininess:    300,
			Reflectivity: 0,
			Transparency: 1,
			IOR:          1,
			Color:        color.Grey(.1),
			Pattern:      nil,
		})

	middle := renderer.NewSphere(
		matrix.Translation(-10, -10, 2).Mult(matrix.Scaling(2, 2, 2)),
		&renderer.PhongMat{
			Ambient:      0.1,
			Diffuse:      0.9,
			Specular:     0.0,
			Shininess:    10,
			Reflectivity: .05,
			Color:        color.NewColor(0.1, 1, 0.5),
			Pattern: renderer.NewCheckerPattern3D(
				matrix.Compose(
					matrix.RotationZ(-math.Pi/6),
					matrix.RotationY(-math.Pi/6),
					matrix.ScalingU(.5),
				),
				renderer.SolidPat(0.1, 1, 0.5),
				renderer.SolidPat(0.1, 0.5, 0.4)),
		})
	left := renderer.NewSphere(
		matrix.Translation(7, -7, 1),
		&renderer.PhongMat{
			Ambient:      0.1,
			Diffuse:      0.9,
			Specular:     0.9,
			Shininess:    200,
			Reflectivity: .1,
			Color:        color.NewColor(1, 0.1, 0.1),
		})
	right := renderer.NewSphere(
		matrix.Translation(-4, 3, 1.25).Mult(matrix.ScalingU(1.25)),
		&renderer.PhongMat{
			Ambient:      0.1,
			Diffuse:      0.9,
			Specular:     0.9,
			Shininess:    200,
			Reflectivity: .1,
			Color:        color.NewColor(0.2, 0.2, 1),
			Pattern: renderer.NewStripePattern(
				matrix.Compose(
					matrix.Translation(0, 0, .25),
					matrix.RotationY(math.Pi/2),
					matrix.ScalingU(.5),
				),
				renderer.NewGradientPattern(
					matrix.Compose(
						matrix.RotationY(math.Pi/2),
						matrix.ScalingU(3),
					),
					renderer.SolidPat(0.9, 0.9, 0.9),
					renderer.SolidPat(0.2, 0.2, 1),
				),
				renderer.SolidPat(0.2, 0.2, 0.4)),
		})

	w := &renderer.World{
		Geometry: []renderer.Primitive{
			room,
			glassBall,
			airBall,
			middle,
			right,
			left,
		},
		Light: &renderer.PointLight{
			Pos:   tuple.NewPoint(0, -10, 10),
			Color: color.White,
		},
		Config: renderer.WorldConfig{
			Shadows:   true,
			MaxBounce: 5,
		},
	}

	c := renderer.NewCamera(1920, 1080, math.Pi/3,
		matrix.ViewTransform(
			tuple.NewPoint(-15, -10, 5),
			tuple.NewPoint(3, 3, 2),
			tuple.Up), 2)

	c.GoRender(w, 8).SaveImage("refraction.png")
}
