package tracer

// Point represents a point.
type Point [3]float64

// NewPoint constructs a new point.
func NewPoint(x float64, y float64, z float64) *Point {
	return &Point{x, y, z}
}

// X returns the x coordinate.
func (p *Point) X() float64 {
	return p[0]
}

// Y returns the y coordinate.
func (p *Point) Y() float64 {
	return p[1]
}

// Z returns the z coordinate.
func (p *Point) Z() float64 {
	return p[2]
}

// W returns the w coordinate.
func (p *Point) W() float64 {
	return 1.0
}

// Equals returns whether a point is approximately equal to this point.
func (p *Point) Equals(q *Point) bool {
	return equals(p.X(), q.X()) && equals(p.Y(), q.Y()) && equals(p.Z(), q.Z())
}

// PlusVector returns a point representing this point plus another vector.
func (p *Point) PlusVector(v *Vector) *Point {
	return NewPoint(p.X()+v.X(), p.Y()+v.Y(), p.Z()+v.Z())
}

// MinusPoint returns a vector representing this point minus another point.
func (p *Point) MinusPoint(q *Point) *Vector {
	return NewVector(p.X()-q.X(), p.Y()-q.Y(), p.Z()-q.Z())
}

// MinusVector returns a point representing this point minus a vector.
func (p *Point) MinusVector(v *Vector) *Point {
	return NewPoint(p.X()-v.X(), p.Y()-v.Y(), p.Z()-v.Z())
}

// TimesScalar returns this point multiplied by a scalar.
func (p *Point) TimesScalar(f float64) *Point {
	return NewPoint(p.X()*f, p.Y()*f, p.Z()*f)
}

// DividedByScalar returns this point divided by a scalar.
func (p *Point) DividedByScalar(f float64) *Point {
	return NewPoint(p.X()/f, p.Y()/f, p.Z()/f)
}
