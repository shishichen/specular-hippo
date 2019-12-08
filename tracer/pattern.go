package tracer

// Pattern is an interface for coloriing patterns.
type Pattern interface {
	// ColorAt returns the color of this pattern at a point.
	ColorAt(*Point) *Color
	// EqualsPattern returns whether this pattern is approximately equal to another pattern.
	EqualsPattern(Pattern) bool
}
