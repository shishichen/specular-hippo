package tracer

import "math"

// GradientPattern represents a gradient pattern.
type GradientPattern struct {
	transformable
	a Pattern
	b Pattern
}

// NewGradientPattern constructs a new gradient pattern.
func NewGradientPattern(a Pattern, b Pattern) *GradientPattern {
	return &GradientPattern{defaultTransformable(), a, b}
}

// WithTransform sets this pattern's transform to a transform.
// May return nil without setting the transform if the transform is invalid.
func (g *GradientPattern) WithTransform(t *Matrix4) *GradientPattern {
	if !g.setTransform(t) {
		return nil
	}
	return g
}

// ColorAt implements the Pattern interface.
func (g *GradientPattern) ColorAt(p *Point) *Color {
	p = g.toLocalPoint(p)
	a, b := g.a.ColorAt(p), g.b.ColorAt(p)
	f := math.Abs(math.Mod(math.Abs(p.X()), 2.0) - 1.0)
	return b.PlusColor(a.MinusColor(b).TimesScalar(f))
}

// EqualsPattern implements the Pattern interface.
func (g *GradientPattern) EqualsPattern(p Pattern) bool {
	h, ok := p.(*GradientPattern)
	if !ok {
		return false
	}
	return g.a.EqualsPattern(h.a) && g.b.EqualsPattern(h.b)
}

// RadialGradientPattern represents a radial gradient pattern.
type RadialGradientPattern struct {
	transformable
	a Pattern
	b Pattern
}

// NewRadialGradientPattern constructs a new gradient pattern.
func NewRadialGradientPattern(a Pattern, b Pattern) *RadialGradientPattern {
	return &RadialGradientPattern{defaultTransformable(), a, b}
}

// WithTransform sets this pattern's transform to a transform.
// May return nil without setting the transform if the transform is invalid.
func (r *RadialGradientPattern) WithTransform(t *Matrix4) *RadialGradientPattern {
	if !r.setTransform(t) {
		return nil
	}
	return r
}

// ColorAt implements the Pattern interface.
func (r *RadialGradientPattern) ColorAt(p *Point) *Color {
	p = r.toLocalPoint(p)
	a, b := r.a.ColorAt(p), r.b.ColorAt(p)
	f := math.Abs(math.Mod(math.Sqrt(p.X()*p.X()+p.Z()*p.Z()), 2.0) - 1.0)
	return b.PlusColor(a.MinusColor(b).TimesScalar(f))
}

// EqualsPattern implements the Pattern interface.
func (r *RadialGradientPattern) EqualsPattern(p Pattern) bool {
	s, ok := p.(*RadialGradientPattern)
	if !ok {
		return false
	}
	return r.a.EqualsPattern(s.a) && r.b.EqualsPattern(s.b)
}
