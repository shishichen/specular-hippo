package tracer

// Material represents a material.
type Material struct {
	c         *Color
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

// NewMaterial constructs a new material.
func NewMaterial(c *Color, ambient float64, diffuse float64, specular float64, shininess float64) *Material {
	if ambient < 0 || diffuse < 0 || specular < 0 || shininess < 0 {
		return nil
	}
	return &Material{c, ambient, diffuse, specular, shininess}
}

// NewDefaultMaterial constructs a new default material.
func NewDefaultMaterial() *Material {
	return NewMaterial(NewColor(1.0, 1.0, 1.0), 0.1, 0.9, 0.9, 200.0)
}

// Color returns this material's color.
func (m *Material) Color() *Color {
	return m.c
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
	return m.Color().Equals(n.Color()) &&
		equals(m.Ambient(), n.Ambient()) && equals(m.Diffuse(), n.Diffuse()) &&
		equals(m.Specular(), n.Specular()) && equals(m.Shininess(), n.Shininess())
}
