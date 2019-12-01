package main

import (
	"fmt"
	"math"

	"github.com/shishichen/specular-hippo/projectile"
	"github.com/shishichen/specular-hippo/tracer"
)

var (
	white = tracer.NewColor(1.0, 1.0, 1.0)
	red   = tracer.NewColor(1.0, 0.0, 0.0)
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
	t := tracer.NewScaling(150, 150, 150).Concatenate(tracer.NewTranslation(200, 200, 0))
	r := tracer.NewRotationZ(math.Pi / -6)
	for i := 0; i < 12; i++ {
		q := t.TimesPoint(p)
		c.SetColor(int(q.X()), int(q.Y()), white)
		p = r.TimesPoint(p)
	}

	return c
}

func drawSphere() *tracer.Canvas {
	c := tracer.NewCanvas(400, 400)
	worldToCanvas := tracer.NewScaling(100, 100, 100).Concatenate(tracer.NewTranslation(200, 200, 0))

	origin := worldToCanvas.TimesPoint(tracer.NewPoint(0, 0, 0))
	sphere := tracer.NewSphere()
	if !sphere.SetTransform(tracer.NewTranslation(0, 0, 7).Concatenate(worldToCanvas)) {
		fmt.Printf("Error: unable to set sphere transform to: %v\n", sphere.Transform())
		return c
	}
	sphere.SetMaterial(tracer.NewMaterial(tracer.NewColor(0.2, 1, 1), 0.1, 0.9, 0.9, 200))
	light := tracer.NewLight(worldToCanvas.TimesPoint(tracer.NewPoint(-10, 10, 0)), tracer.NewColor(1, 1, 1))
	for i := 0; i < c.Width(); i++ {
		for j := 0; j < c.Height(); j++ {
			pixel := tracer.NewPoint(float64(i), float64(c.Height()-j-1), 1000)
			ray := tracer.NewRay(origin, pixel.MinusPoint(origin).Normalize())
			hit := sphere.Intersects(ray).Hit()
			if hit == nil {
				continue
			}
			point := ray.Position(hit.T())
			normal := hit.Sphere().NormalAt(point)
			eye := ray.Direction().TimesScalar(-1)
			color := light.Illuminate(point, sphere.Material(), normal, eye)
			c.SetColor(i, j, color)
		}
	}

	return c
}

func main() {
	c := drawSphere()
	c.ToFile()
}
