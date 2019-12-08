package tracer

// Plane represents a plane.
type Plane struct {
	internalShape
}

// NewPlane constructs a new plane.
func NewPlane() *Plane {
	return &Plane{defaultInternalShape()}
}

// WithMaterial sets this plane's material to a material.
func (p *Plane) WithMaterial(m *Material) *Plane {
	p.setMaterial(m)
	return p
}

// WithTransform sets this plane's transform to a matrix.
// May return nil without setting the transform if the transform is invalid.
func (p *Plane) WithTransform(t *Matrix4) *Plane {
	p.setTransform(t)
	return p
}

// Intersect implements the Shape interface.
func (p *Plane) Intersect(r *Ray) Intersections {
	or := r
	r = p.toLocalRay(r)
	if equals(r.Direction().Y(), 0.0) {
		return NewIntersections()
	}
	t := -r.Origin().Y() / r.Direction().Y()
	return NewIntersections(NewIntersection(p, or, t))
}

// NormalAt implements the Shape interface.
func (p *Plane) NormalAt(*Point) *Vector {
	normal := NewVector(0.0, 1.0, 0.0)
	return p.toWorldVector(normal)
}
