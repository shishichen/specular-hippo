package main

import (
	"fmt"

	"github.com/shishichen/specular-hippo/tracer"
)

type environment struct {
	gravity *tracer.Vector
	wind    *tracer.Vector
}

type projectile struct {
	position *tracer.Point
	velocity *tracer.Vector
}

func tick(e environment, p projectile) projectile {
	position := p.position.PlusVector(p.velocity)
	velocity := p.velocity.PlusVector(e.gravity).PlusVector(e.wind)
	return projectile{position, velocity}
}

func main() {
	const factor = 4.0
	e := environment{tracer.NewVector(0.0, -0.1, 0.0), tracer.NewVector(-0.01, 0.0, 0.0)}
	p := projectile{tracer.NewPoint(0.0, 1.0, 0.0), tracer.NewVector(1.0, 1.0, 0.0).Normalize().TimesScalar(factor)}
	fmt.Printf("initial position: %v\n", *p.position)
	for i := 1; p.position.Y() > 0.0; i++ {
		p = tick(e, p)
		fmt.Printf("new position after %v ticks: %v\n", i, *p.position)
	}
}
