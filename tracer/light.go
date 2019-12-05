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

// IsShadowed returns whether a point is in shadow of this light given a collection of shapes.
func (l *Light) IsShadowed(p *Point, s Shapes) bool {
	light := l.Position().MinusPoint(p)
	r := NewRay(p, light.Normalize())
	hit := s.Intersect(r).Hit()
	return hit != nil && hit.T() < light.Magnitude()
}

// Illuminate returns the color of a point given a material, normal vector, eye vector, and whether it's in shadow.
func (l *Light) Illuminate(m *Material, p *Point, normal *Vector, eye *Vector, s Shapes) *Color {
	color := m.Color().TimesColor(l.Intensity())
	ambient := color.TimesScalar(m.Ambient())
	if l.IsShadowed(p, s) {
		return ambient
	}
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
