package tracer

import "math"

// Light represents a point light.
type Light struct {
	p *Point
	i *Color
}

// Lights represents a collection of lights.
type Lights []*Light

// NewLight constructs a new light.
func NewLight(p *Point, i *Color) *Light {
	return &Light{p, i}
}

// Position returns this light's position.
func (l *Light) Position() *Point {
	return l.p
}

// Intensity returns this light's intensity.
func (l *Light) Intensity() *Color {
	return l.i
}

// Illuminate returns the color of a point given the shape the point is on, normal vector, eye vector, and whether it's in shadow.
func (l *Light) Illuminate(s Shape, p *Point, normal *Vector, eye *Vector, isShadowed bool) *Color {
	color := s.ColorAt(p).TimesColor(l.Intensity())
	ambient := color.TimesScalar(s.Material().Ambient())
	if isShadowed {
		return ambient
	}
	light := l.Position().MinusPoint(p).Normalize()
	lightDotNormal := light.DotVector(normal)
	if lightDotNormal < 0.0 {
		return ambient
	}
	diffuse := color.TimesScalar(s.Material().Diffuse()).TimesScalar(lightDotNormal)
	reflect := normal.Reflect(light.TimesScalar(-1.0))
	reflectDotEye := reflect.DotVector(eye)
	if reflectDotEye < 0.0 {
		return ambient.PlusColor(diffuse)
	}
	specular := l.Intensity().TimesScalar(s.Material().Specular()).TimesScalar(math.Pow(reflectDotEye, s.Material().Shininess()))
	return ambient.PlusColor(diffuse).PlusColor(specular)
}
