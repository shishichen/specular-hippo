package tracer

// Material represents a material.
type Material struct {
	p         Pattern
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

// NewMaterial constructs a new material.
func NewMaterial(p Pattern, ambient float64, diffuse float64, specular float64, shininess float64) *Material {
	if ambient < 0 || diffuse < 0 || specular < 0 || shininess < 0 {
		return nil
	}
	return &Material{p, ambient, diffuse, specular, shininess}
}

// NewDefaultMaterial constructs a new default material.
func NewDefaultMaterial() *Material {
	return NewMaterial(NewSolidPattern(white), 0.1, 0.9, 0.9, 200.0)
}

// Pattern this material's pattern.
func (m *Material) Pattern() Pattern {
	return m.p
}

// Ambient returns this material's ambient.
func (m *Material) Ambient() float64 {
	return m.ambient
}

// Diffuse returns this material's diffuse.
func (m *Material) Diffuse() float64 {
	return m.diffuse
}

// Specular returns this material's specular.
func (m *Material) Specular() float64 {
	return m.specular
}

// Shininess returns this material's shininess.
func (m *Material) Shininess() float64 {
	return m.shininess
}

// Equals returns whether a material is approximately equal to this material.
func (m *Material) Equals(n *Material) bool {
	return m.Pattern().EqualsPattern(n.Pattern()) &&
		equals(m.Ambient(), n.Ambient()) && equals(m.Diffuse(), n.Diffuse()) &&
		equals(m.Specular(), n.Specular()) && equals(m.Shininess(), n.Shininess())
}

// ColorAt returns this material's color at a point.
func (m *Material) ColorAt(p *Point) *Color {
	return m.p.ColorAt(p)
}
