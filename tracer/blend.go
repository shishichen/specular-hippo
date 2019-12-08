package tracer

// BlendedPattern represents a blended pattern.
type BlendedPattern struct {
	transformable
	a Pattern
	b Pattern
}

// NewBlendedPattern constructs a new blended pattern.
func NewBlendedPattern(a Pattern, b Pattern) *BlendedPattern {
	return &BlendedPattern{defaultTransformable(), a, b}
}

// WithTransform sets this pattern's transform to a transform.
// May return nil without setting the transform if the transform is invalid.
func (l *BlendedPattern) WithTransform(t *Matrix4) *BlendedPattern {
	if !l.setTransform(t) {
		return nil
	}
	return l
}

// ColorAt implements the Pattern interface.
func (l *BlendedPattern) ColorAt(p *Point) *Color {
	p = l.toLocalPoint(p)
	a, b := l.a.ColorAt(p), l.b.ColorAt(p)
	return a.PlusColor(b).TimesScalar(0.5)
}

// EqualsPattern implements the Pattern interface.
func (l *BlendedPattern) EqualsPattern(p Pattern) bool {
	m, ok := p.(*BlendedPattern)
	if !ok {
		return false
	}
	return l.a.EqualsPattern(m.a) && l.b.EqualsPattern(m.b)
}
