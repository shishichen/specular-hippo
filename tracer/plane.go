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

// Intersect returns this plane's intersection points with a ray.
func (p *Plane) Intersect(r *Ray) Intersections {
	or := r
	r = p.worldToObjectRay(r)
	if equals(r.Direction().Y(), 0.0) {
		return NewIntersections()
	}
	t := -1.0 * r.Origin().Y() / r.Direction().Y()
	return NewIntersections(NewIntersection(p, or, t))
}

// NormalAt returns the normal at a point on this plane.
func (p *Plane) NormalAt(*Point) *Vector {
	normal := NewVector(0.0, 1.0, 0.0)
	return p.objectToWorldVector(normal)
}
