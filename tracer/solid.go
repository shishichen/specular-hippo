package tracer

// SolidPattern represents a solid pattern.
type SolidPattern struct {
	c *Color
}

// NewSolidPattern constructs a new solid pattern.
func NewSolidPattern(c *Color) *SolidPattern {
	return &SolidPattern{c}
}

// ColorAt implements the Pattern interface.
func (s *SolidPattern) ColorAt(p *Point) *Color {
	return s.c
}

// EqualsPattern implements the Pattern interface.
func (s *SolidPattern) EqualsPattern(p Pattern) bool {
	t, ok := p.(*SolidPattern)
	if !ok {
		return false
	}
	return s.c.Equals(t.c)
}
