package tracer

// Shape is an interface for shapes.
type Shape interface {
	Material() *Material
	Intersect(r *Ray) Intersections
	NormalAt(p *Point) *Vector
}

// Shapes represents a collection of shapes.
type Shapes []Shape

// Intersect returns a collection of shapes' intersection points with a ray.
func (s Shapes) Intersect(r *Ray) Intersections {
	intersections := make([]Intersections, len(s))
	for i, shape := range s {
		intersections[i] = shape.Intersect(r)
	}
	return MergeIntersections(intersections)
}

// internalShape represents the some of the internals of shapes.
type internalShape struct {
	m        *Material
	t        *Matrix4
	inverseT *Matrix4
}

func defaultInternalShape() internalShape {
	return internalShape{NewDefaultMaterial(), NewIdentity(), NewIdentity()}
}

// Material returns this shape's material.
func (s *internalShape) Material() *Material {
	return s.m
}

// Transform returns this shape's transform.
func (s *internalShape) Transform() *Matrix4 {
	return s.t
}

func (s *internalShape) inverseTransform() *Matrix4 {
	return s.inverseT
}
func (s *internalShape) setMaterial(m *Material) {
	s.m = m
}

func (s *internalShape) setTransform(t *Matrix4) bool {
	if !t.HasInverse() {
		return false
	}
	s.t = t
	s.inverseT = t.Inverse()
	return true
}

func (s *internalShape) worldToObjectPoint(p *Point) *Point {
	return s.inverseTransform().TimesPoint(p)
}

func (s *internalShape) worldToObjectRay(r *Ray) *Ray {
	return s.inverseTransform().TimesRay(r)
}

func (s *internalShape) objectToWorldVector(v *Vector) *Vector {
	return s.inverseTransform().Transpose().TimesVector(v).Normalize()
}
