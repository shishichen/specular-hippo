package tracer

// Shape is an interface for shapes.
type Shape interface {
	Material() *Material
	NormalAt(p *Point) *Vector
	Intersect(r *Ray) Intersections
}

// Shapes represents a collection of shapes.
type Shapes []Shape
