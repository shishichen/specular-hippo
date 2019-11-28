package projectile

import (
	"github.com/shishichen/specular-hippo/tracer"
)

// Environment represents the environment.
type Environment struct {
	Gravity *tracer.Vector
	Wind    *tracer.Vector
}

// Projectile represents the projectile.
type Projectile struct {
	Position *tracer.Point
	Velocity *tracer.Vector
}

// Tick returns a projectile representing the state after a unit of time.
func Tick(e Environment, p Projectile) Projectile {
	position := p.Position.PlusVector(p.Velocity)
	velocity := p.Velocity.PlusVector(e.Gravity).PlusVector(e.Wind)
	return Projectile{position, velocity}
}
