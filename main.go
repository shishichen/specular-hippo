package main

import (
	"math"

	"github.com/shishichen/specular-hippo/projectile"
	"github.com/shishichen/specular-hippo/tracer"
)

var (
	black = tracer.NewColor(0, 0, 0)
	white = tracer.NewColor(1, 1, 1)
	red   = tracer.NewColor(1, 0, 0)
	blue  = tracer.NewColor(0, 0, 1)
)

func drawProjectile() *tracer.Canvas {
	c := tracer.NewCanvas(900, 550)

	e := projectile.Environment{Gravity: tracer.NewVector(0.0, -0.1, 0.0), Wind: tracer.NewVector(-0.01, 0.0, 0.0)}
	p := projectile.Projectile{Position: tracer.NewPoint(0.0, 1.0, 0.0), Velocity: tracer.NewVector(1.0, 1.8, 0.0).Normalize().TimesScalar(11.25)}
	for i := 1; p.Position.Y() >= 0.0; i++ {
		c.SetColor(int(p.Position.X()), c.Height()-int(p.Position.Y())-1, white)
		p = projectile.Tick(e, p)
	}

	return c
}

func drawClock() *tracer.Canvas {
	c := tracer.NewCanvas(400, 400)

	p := tracer.NewPoint(0, 1, 0)
	t := tracer.NewScale(150, 150, 150).Translate(200, 200, 0)
	r := tracer.NewRotateZ(math.Pi / -6)
	for i := 0; i < 12; i++ {
		q := t.TimesPoint(p)
		c.SetColor(int(q.X()), int(q.Y()), white)
		p = r.TimesPoint(p)
	}

	return c
}

func drawSphere() *tracer.Canvas {
	world := tracer.NewWorld(
		tracer.Shapes{
			tracer.NewSphere().WithTransform(tracer.NewScale(4.0, 4.0, 4.0).Translate(0, 0, -7)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(0.2, 1, 1)), 0.1, 0.9, 0.9, 200))},
		tracer.Lights{tracer.NewLight(tracer.NewPoint(10, 10, 10), tracer.NewColor(1, 1, 1))})
	camera := tracer.NewCamera(400, 400, math.Pi/2).
		WithTransformFromParameters(tracer.NewPoint(0, 0, 0), tracer.NewPoint(0, 0, -1), tracer.NewVector(0, 1, 0))
	return camera.Render(world)
}

func drawSixSpheres() *tracer.Canvas {
	world := tracer.NewWorld(
		tracer.Shapes{
			tracer.NewSphere().
				WithTransform(tracer.NewScale(10, 0.01, 10)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.9, 0.9)), 0.2, 0.9, 0, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(10, 0.01, 10).RotateX(math.Pi/2).RotateY(math.Pi/-4).Translate(0, 0, 5)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.9, 0.9)), 0.2, 0.9, 0, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(10, 0.01, 10).RotateX(math.Pi/2).RotateY(math.Pi/4).Translate(0, 0, 5)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.9, 0.9)), 0.2, 0.9, 0, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewTranslate(-0.5, 1, 0.5)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(0.1, 1, 0.5)), 0.2, 0.7, 0.3, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(0.5, 1, 0.1)), 0.2, 0.7, 0.3, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.8, 0.1)), 0.2, 0.7, 0.3, 200)),
		},
		tracer.Lights{
			tracer.NewLight(tracer.NewPoint(-10, 10, -10), tracer.NewColor(1, 1, 1)),
		})
	camera := tracer.NewCamera(1000, 500, math.Pi/3).
		WithTransformFromParameters(tracer.NewPoint(0, 1.5, -5), tracer.NewPoint(0, 1, 0), tracer.NewVector(0, 1, 0))
	return camera.WithThreads(2).Render(world)
}

func drawPlaneAndThreeSpheres() *tracer.Canvas {
	world := tracer.NewWorld(
		tracer.Shapes{
			tracer.NewPlane().
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.9, 0.9)), 0.2, 0.9, 0, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewTranslate(-0.5, 1, 0.5)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(0.1, 1, 0.5)), 0.2, 0.7, 0.3, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(0.5, 1, 0.1)), 0.2, 0.7, 0.3, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75)).
				WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.8, 0.1)), 0.2, 0.7, 0.3, 200)),
		},
		tracer.Lights{
			tracer.NewLight(tracer.NewPoint(-10, 10, -10), tracer.NewColor(1, 1, 1)),
		})
	camera := tracer.NewCamera(800, 800, math.Pi/3).
		WithTransformFromParameters(tracer.NewPoint(0, 1.5, -5), tracer.NewPoint(0, 1, 0), tracer.NewVector(0, 1, 0))
	return camera.WithThreads(2).Render(world)
}

func drawHexagonalRoom() *tracer.Canvas {
	shapes := [7]tracer.Shape{}
	for i := 0; i < 6; i++ {
		angle := float64(i) * math.Pi / 3
		shapes[i] = tracer.NewPlane().
			WithTransform(tracer.NewRotateZ(math.Pi/2).RotateY(angle).Translate(-math.Cos(angle), 0, math.Sin(angle))).
			WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.9, 0.9)), 0.2, 0.9, 0, 200))
	}
	shapes[6] = tracer.NewPlane().
		WithMaterial(tracer.NewMaterial(tracer.NewSolidPattern(tracer.NewColor(1, 0.9, 0.9)), 0.2, 0.9, 0, 200))
	world := tracer.NewWorld(
		shapes[:],
		tracer.Lights{
			tracer.NewLight(tracer.NewPoint(0.5, 10, 0.5), tracer.NewColor(1, 1, 1)),
		})
	camera := tracer.NewCamera(800, 800, math.Pi/3).
		WithTransformFromParameters(tracer.NewPoint(0, 5, 0), tracer.NewPoint(0, 0, 0), tracer.NewVector(0, 0, 1))
	return camera.WithThreads(2).Render(world)
}

func drawStripedSpheres() *tracer.Canvas {
	world := tracer.NewWorld(
		tracer.Shapes{
			tracer.NewPlane().
				WithMaterial(tracer.NewMaterial(
					tracer.NewBlendedPattern(
						tracer.NewStripePattern(tracer.NewSolidPattern(white), tracer.NewSolidPattern(blue)),
						tracer.NewStripePattern(tracer.NewSolidPattern(white), tracer.NewSolidPattern(blue)).
							WithTransform(tracer.NewRotateY(math.Pi/2.0))).
						WithTransform(tracer.NewRotateY(math.Pi/4)),
					0.2, 0.9, 0, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewTranslate(-0.5, 1, 0.5)).
				WithMaterial(tracer.NewMaterial(
					tracer.NewPerturbedPattern(
						tracer.NewRingPattern(tracer.NewSolidPattern(black), tracer.NewSolidPattern(white)).
							WithTransform(tracer.NewScale(0.21, 0.21, 0.21).RotateX(-math.Pi/4)),
						true),
					0.2, 0.7, 0.3, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5)).
				WithMaterial(tracer.NewMaterial(
					tracer.NewPerturbedPattern(
						tracer.NewGradientPattern(tracer.NewSolidPattern(white), tracer.NewSolidPattern(black)).
							WithTransform(tracer.NewScale(0.25, 0.25, 0.25).RotateY(math.Pi/6).RotateZ(-math.Pi/2)),
						true),
					0.2, 0.7, 0.3, 200)),
			tracer.NewSphere().
				WithTransform(tracer.NewScale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75)).
				WithMaterial(tracer.NewMaterial(
					tracer.NewPerturbedPattern(
						tracer.NewCheckerPattern(tracer.NewSolidPattern(white), tracer.NewSolidPattern(black)).
							WithTransform(tracer.NewScale(0.7, 0.7, 0.7).RotateZ(math.Pi/3).Translate(0.5, 0, 0)),
						true),
					0.2, 0.7, 0.3, 200)),
		},
		tracer.Lights{
			tracer.NewLight(tracer.NewPoint(-10, 10, -10), tracer.NewColor(1, 1, 1)),
		})
	camera := tracer.NewCamera(800, 800, math.Pi/3).
		WithTransformFromParameters(tracer.NewPoint(0, 1.5, -5), tracer.NewPoint(0, 1, 0), tracer.NewVector(0, 1, 0))
	return camera.WithThreads(2).Render(world)
}

func main() {
	c := drawStripedSpheres()
	c.ToFile()
}
