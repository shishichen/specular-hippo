package tracer

// Shape is an interface for shapes.
type Shape interface {
	// Material returns this shape's material.
	Material() *Material
	// Intersect returns this shape's intersection points with a ray.
	Intersect(*Ray) Intersections
	// ColorAt returns the color at a point on this shape.
	ColorAt(*Point) *Color
	// NormalAt returns the normal at a point on this shape.
	NormalAt(*Point) *Vector
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
	transformable
	m *Material
}

func defaultInternalShape() internalShape {
	return internalShape{defaultTransformable(), NewDefaultMaterial()}
}

// Material implements the Shape interface.
func (s *internalShape) Material() *Material {
	return s.m
}

// ColorAt implements the Shape interface.
func (s *internalShape) ColorAt(p *Point) *Color {
	p = s.toLocalPoint(p)
	return s.Material().ColorAt(p)
}

func (s *internalShape) setMaterial(m *Material) {
	s.m = m
}
