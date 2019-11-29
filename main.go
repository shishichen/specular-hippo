package main

import (
	"math"

	"github.com/shishichen/specular-hippo/projectile"
	"github.com/shishichen/specular-hippo/tracer"
)

var (
	white = tracer.NewColor(1.0, 1.0, 1.0)
)

func drawProjectile() *tracer.Canvas {
	c := tracer.NewCanvas(900, 550)

	e := projectile.Environment{Gravity: tracer.NewVector(0.0, -0.1, 0.0), Wind: tracer.NewVector(-0.01, 0.0, 0.0)}
	p := projectile.Projectile{Position: tracer.NewPoint(0.0, 1.0, 0.0), Velocity: tracer.NewVector(1.0, 1.8, 0.0).Normalize().TimesScalar(11.25)}
	for i := 1; p.Position.Y() >= 0.0; i++ {
		c.SetColor(int(p.Position.X()), c.Height()-int(p.Position.Y()), white)
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

func main() {
	c := drawClock()
	c.ToFile()
}
