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
	o := worldToCanvas.TimesPoint(tracer.NewPoint(0, 0, 10))
	s := tracer.NewSphere()
	if !s.SetTransform(tracer.NewTranslation(0, 0, 3).Concatenate(worldToCanvas)) {
		fmt.Printf("Error: unable to set sphere transform to: %v\n", s.Transform())
		return c
	}
	for i := 0; i < c.Width(); i++ {
		for j := 0; j < c.Height(); j++ {
			p := tracer.NewPoint(float64(i), float64(c.Height()-j-1), 0)
			r := tracer.NewRay(o, p.MinusPoint(o).Normalize())
			if len(s.Intersects(r)) > 0 {
				c.SetColor(i, j, red)
			}
		}
	}

	return c
}

func main() {
	c := drawSphere()
	c.ToFile()
}
