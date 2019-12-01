package tracer

import "math"

// Light represents a point light.
type Light struct {
	p *Point
	i *Color
}

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

// Illuminate returns the color of a point given a material, normal, and eye vector.
func (l *Light) Illuminate(p *Point, m *Material, normal *Vector, eye *Vector) *Color {
	color := m.Color().TimesColor(l.Intensity())
	ambient := color.TimesScalar(m.Ambient())
	light := l.Position().MinusPoint(p).Normalize()
	lightDotNormal := light.DotVector(normal)
	if lightDotNormal < 0.0 {
		return ambient
	}
	diffuse := color.TimesScalar(m.Diffuse()).TimesScalar(lightDotNormal)
	reflect := normal.Reflect(light.TimesScalar(-1.0))
	reflectDotEye := reflect.DotVector(eye)
	if reflectDotEye < 0.0 {
		return ambient.PlusColor(diffuse)
	}
	specular := l.Intensity().TimesScalar(m.Specular()).TimesScalar(math.Pow(reflectDotEye, m.Shininess()))
	return ambient.PlusColor(diffuse).PlusColor(specular)
}
