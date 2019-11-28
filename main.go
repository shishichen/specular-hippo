package main

import (
	"github.com/shishichen/specular-hippo/projectile"
	"github.com/shishichen/specular-hippo/tracer"
)

func main() {
	e := projectile.Environment{Gravity: tracer.NewVector(0.0, -0.1, 0.0), Wind: tracer.NewVector(-0.01, 0.0, 0.0)}
	p := projectile.Projectile{Position: tracer.NewPoint(0.0, 1.0, 0.0), Velocity: tracer.NewVector(1.0, 1.8, 0.0).Normalize().TimesScalar(11.25)}

	c := tracer.NewCanvas(900, 550)
	white := tracer.NewColor(1.0, 1.0, 1.0)

	for i := 1; p.Position.Y() >= 0.0; i++ {
		c.SetColor(int(p.Position.X()), c.Height()-int(p.Position.Y()), white)
		p = projectile.Tick(e, p)
	}

	c.ToFile()
}
