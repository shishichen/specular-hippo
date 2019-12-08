package tracer

import "math"

// StripePattern represents a stripe pattern in the X dimension.
type StripePattern struct {
	transformable
	a Pattern
	b Pattern
}

// NewStripePattern constructs a new stripe pattern.
func NewStripePattern(a Pattern, b Pattern) *StripePattern {
	return &StripePattern{defaultTransformable(), a, b}
}

// WithTransform sets this pattern's transform to a transform.
// May return nil without setting the transform if the transform is invalid.
func (s *StripePattern) WithTransform(t *Matrix4) *StripePattern {
	if !s.setTransform(t) {
		return nil
	}
	return s
}

// ColorAt implements the Pattern interface.
func (s *StripePattern) ColorAt(p *Point) *Color {
	p = s.toLocalPoint(p)
	a, b := s.a.ColorAt(p), s.b.ColorAt(p)
	if equals(math.Mod(math.Floor(p.X()), 2.0), 0.0) {
		return a
	}
	return b
}

// EqualsPattern implements the Pattern interface.
func (s *StripePattern) EqualsPattern(p Pattern) bool {
	t, ok := p.(*StripePattern)
	if !ok {
		return false
	}
	return s.a.EqualsPattern(t.a) && s.b.EqualsPattern(t.b)
}
