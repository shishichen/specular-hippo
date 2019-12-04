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

// WithTransform sets this sphere's transform to a matrix.
// May return nil without setting the transform if the transform is invalid.
func (s *Sphere) WithTransform(t *Matrix4) *Sphere {
	if !t.HasInverse() {
		return nil
	}
	s.transform = t
	return s
}

// WithMaterial sets this sphere's material to a material.
func (s *Sphere) WithMaterial(m *Material) *Sphere {
	s.m = m
	return s
}

// Intersect returns this sphere's intersection points with a ray.
func (s *Sphere) Intersect(r *Ray) Intersections {
	rObject := s.Transform().Inverse().TimesRay(r)
	sphereToRay := rObject.Origin().MinusPoint(NewPoint(0.0, 0.0, 0.0))
	a := rObject.Direction().DotVector(rObject.Direction())
	b := 2.0 * rObject.Direction().DotVector(sphereToRay)
	c := sphereToRay.DotVector(sphereToRay) - 1.0
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return NewIntersections()
	}
	t1 := (-1.0*b - math.Sqrt(discriminant)) / (2.0 * a)
	t2 := (-1.0*b + math.Sqrt(discriminant)) / (2.0 * a)
	return NewIntersections(NewIntersection(r, t1, s), NewIntersection(r, t2, s))
}

// NormalAt returns the normal at a point on this sphere.
func (s *Sphere) NormalAt(p *Point) *Vector {
	objectPoint := s.Transform().Inverse().TimesPoint(p)
	objectNormal := objectPoint.MinusPoint(NewPoint(0.0, 0.0, 0.0))
	worldNormal := s.Transform().Inverse().Transpose().TimesVector(objectNormal)
	return worldNormal.Normalize()
}
