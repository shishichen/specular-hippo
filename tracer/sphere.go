package tracer

import "math"

// Sphere represents a sphere.
type Sphere struct {
	transform *Matrix4
	m         *Material
}

// NewSphere constructs a new sphere.
func NewSphere() *Sphere {
	return &Sphere{NewIdentity(), NewDefaultMaterial()}
}

// Transform returns this sphere's transform.
func (s *Sphere) Transform() *Matrix4 {
	return s.transform
}

// Material returns this sphere's material.
func (s *Sphere) Material() *Material {
	return s.m
}

// SetTransform sets this sphere's transform to a matrix.
func (s *Sphere) SetTransform(t *Matrix4) bool {
	if !t.HasInverse() {
		return false
	}
	s.transform = t
	return true
}

// SetMaterial sets this sphere's material to a material.
func (s *Sphere) SetMaterial(m *Material) {
	s.m = m
}

// Intersects returns intersection points with a ray.
func (s *Sphere) Intersects(r *Ray) Intersections {
	r = s.Transform().Inverse().TimesRay(r)
	sphereToRay := r.Origin().MinusPoint(NewPoint(0.0, 0.0, 0.0))
	a := r.Direction().DotVector(r.Direction())
	b := 2.0 * r.Direction().DotVector(sphereToRay)
	c := sphereToRay.DotVector(sphereToRay) - 1.0
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return NewIntersections([]*Intersection{})
	}
	t1 := (-1.0*b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-1.0*b + math.Sqrt(discriminant)) / (2.0 * a)
	return NewIntersections([]*Intersection{NewIntersection(t1, s), NewIntersection(t2, s)})
}

// NormalAt returns the normal at a point on this sphere.
func (s *Sphere) NormalAt(p *Point) *Vector {
	objectPoint := s.Transform().Inverse().TimesPoint(p)
	objectNormal := objectPoint.MinusPoint(NewPoint(0.0, 0.0, 0.0))
	worldNormal := s.Transform().Inverse().Transpose().TimesVector(objectNormal)
	return worldNormal.Normalize()
}
