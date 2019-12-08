package tracer

import "math"

// RingPattern represents a ring pattern.
type RingPattern struct {
	transformable
	a Pattern
	b Pattern
}

// NewRingPattern constructs a new ring pattern.
func NewRingPattern(a Pattern, b Pattern) *RingPattern {
	return &RingPattern{defaultTransformable(), a, b}
}

// WithTransform sets this pattern's transform to a transform.
// May return nil without setting the transform if the transform is invalid.
func (r *RingPattern) WithTransform(t *Matrix4) *RingPattern {
	if !r.setTransform(t) {
		return nil
	}
	return r
}

// ColorAt implements the Pattern interface.
func (r *RingPattern) ColorAt(p *Point) *Color {
	p = r.toLocalPoint(p)
	a, b := r.a.ColorAt(p), r.b.ColorAt(p)
	if equals(math.Mod(math.Floor(math.Sqrt(p.X()*p.X()+p.Z()*p.Z())), 2.0), 0.0) {
		return a
	}
	return b
}

// EqualsPattern implements the Pattern interface.
func (r *RingPattern) EqualsPattern(p Pattern) bool {
	s, ok := p.(*RingPattern)
	if !ok {
		return false
	}
	return r.a.EqualsPattern(s.a) && r.b.EqualsPattern(s.b)
}
