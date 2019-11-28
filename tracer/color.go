package tracer

import "math"

// Color represents a color.
type Color [3]float64

// NewColor constructs a new color.
func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

// R returns the red value.
func (c *Color) R() float64 {
	return c[0]
}

// G returns the green value.
func (c *Color) G() float64 {
	return c[1]
}

// B returns the blue value.
func (c *Color) B() float64 {
	return c[2]
}

// Equals returns whether a color is approximately equal to this color.
func (c *Color) Equals(d *Color) bool {
	return equals(c.R(), d.R()) && equals(c.G(), d.G()) && equals(c.B(), d.B())
}

// PlusColor returns a color representing this color plus another color.
func (c *Color) PlusColor(d *Color) *Color {
	return NewColor(c.R()+d.R(), c.G()+d.G(), c.B()+d.B())
}

// MinusColor returns a color representing this color minus another color.
func (c *Color) MinusColor(d *Color) *Color {
	return NewColor(c.R()-d.R(), c.G()-d.G(), c.B()-d.B())
}

// TimesScalar returns this color multiplied by a scalar.
func (c *Color) TimesScalar(f float64) *Color {
	return NewColor(c.R()*f, c.G()*f, c.B()*f)
}

// TimesColor returns this color multiplied by another color.
func (c *Color) TimesColor(d *Color) *Color {
	return NewColor(c.R()*d.R(), c.G()*d.G(), c.B()*d.B())
}

func convert(c float64) uint32 {
	const min = 0
	const max = 0xffff
	return uint32(math.Min(math.Max(max*c, min), max))
}

// RGBA implements the image.Color interface.
func (c *Color) RGBA() (r, g, b, a uint32) {
	return convert(c.R()), convert(c.G()), convert(c.B()), convert(1.0)
}
