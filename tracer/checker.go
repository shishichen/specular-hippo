package tracer

import "math"

// CheckerPattern represents a checker pattern.
type CheckerPattern struct {
	transformable
	a Pattern
	b Pattern
}

// NewCheckerPattern constructs a new checker pattern.
func NewCheckerPattern(a Pattern, b Pattern) *CheckerPattern {
	return &CheckerPattern{defaultTransformable(), a, b}
}

// WithTransform sets this pattern's transform to a transform.
// May return nil without setting the transform if the transform is invalid.
func (c *CheckerPattern) WithTransform(t *Matrix4) *CheckerPattern {
	if !c.setTransform(t) {
		return nil
	}
	return c
}

// ColorAt implements the Pattern interface.
func (c *CheckerPattern) ColorAt(p *Point) *Color {
	p = c.toLocalPoint(p)
	a, b := c.a.ColorAt(p), c.b.ColorAt(p)
	if equals(math.Mod(math.Floor(p.X())+math.Floor(p.Y())+math.Floor(p.Z()), 2.0), 0.0) {
		return a
	}
	return b
}

// EqualsPattern implements the Pattern interface.
func (c *CheckerPattern) EqualsPattern(p Pattern) bool {
	d, ok := p.(*CheckerPattern)
	if !ok {
		return false
	}
	return c.a.EqualsPattern(d.a) && c.b.EqualsPattern(d.b)
}
