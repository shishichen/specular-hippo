package tracer

import "math"

// Sphere represents a sphere.
type Sphere struct {
	internalShape
}

// NewSphere constructs a new sphere.
func NewSphere() *Sphere {
	return &Sphere{defaultInternalShape()}
}

// WithMaterial sets this sphere's material to a material.
func (s *Sphere) WithMaterial(m *Material) *Sphere {
	s.setMaterial(m)
	return s
}

// WithTransform sets this sphere's transform to a matrix.
// May return nil without setting the transform if the transform is invalid.
func (s *Sphere) WithTransform(t *Matrix4) *Sphere {
	if !s.setTransform(t) {
		return nil
	}
	return s
}

// Intersect implements the Shape interface.
func (s *Sphere) Intersect(r *Ray) Intersections {
	or := r // save so we can return as part of the intersection
	r = s.toLocalRay(r)
	sphereToRay := r.Origin().MinusPoint(NewPoint(0.0, 0.0, 0.0))
	a := r.Direction().DotVector(r.Direction())
	b := 2.0 * r.Direction().DotVector(sphereToRay)
	c := sphereToRay.DotVector(sphereToRay) - 1.0
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return NewIntersections()
	}
	t1 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2.0 * a)
	return NewIntersections(NewIntersection(s, or, t1), NewIntersection(s, or, t2))
}

// NormalAt implements the Shape interface.
func (s *Sphere) NormalAt(p *Point) *Vector {
	p = s.toLocalPoint(p)
	normal := p.MinusPoint(NewPoint(0.0, 0.0, 0.0))
	return s.toWorldVector(normal)
}
