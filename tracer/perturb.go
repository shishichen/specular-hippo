package tracer

import (
	"math/rand"
	"time"

	opensimplex "github.com/ojrac/opensimplex-go"
)

// PerturbedPattern represents a perturbed pattern.
type PerturbedPattern struct {
	transformable
	p Pattern
	n [3]opensimplex.Noise
}

// NewPerturbedPattern constructs a new perturbed pattern.
func NewPerturbedPattern(p Pattern, randomize bool) *PerturbedPattern {
	var s rand.Source
	if randomize {
		s = rand.NewSource(time.Now().Unix())
	} else {
		s = rand.NewSource(0)
	}
	return &PerturbedPattern{defaultTransformable(), p, [3]opensimplex.Noise{
		opensimplex.NewNormalized(s.Int63()), opensimplex.NewNormalized(s.Int63()), opensimplex.NewNormalized(s.Int63())}}

}

// WithTransform sets this pattern's transform to a transform.
// May return nil without setting the transform if the transform is invalid.
func (e *PerturbedPattern) WithTransform(t *Matrix4) *PerturbedPattern {
	if !e.setTransform(t) {
		return nil
	}
	return e
}

// ColorAt implements the Pattern interface.
func (e *PerturbedPattern) ColorAt(p *Point) *Color {
	p = e.toLocalPoint(p)
	var f [3]float64
	for i := 0; i < 3; i++ {
		f[i] = 0.5 - e.n[i].Eval3(p.X(), p.Y(), p.Z())
	}
	return e.p.ColorAt(NewPoint(p.X()+f[0], p.Y()+f[1], p.Z()+f[2]))
}

// EqualsPattern implements the Pattern interface.
func (e *PerturbedPattern) EqualsPattern(p Pattern) bool {
	f, ok := p.(*PerturbedPattern)
	if !ok {
		return false
	}
	return e.p.EqualsPattern(f.p)
}
