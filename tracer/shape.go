package tracer

// Shape is an interface for shapes.
type Shape interface {
	Material() *Material
	NormalAt(p *Point) *Vector
	Intersect(r *Ray) Intersections
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
