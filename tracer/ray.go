package tracer

// Ray represents a ray.
type Ray struct {
	o *Point
	d *Vector
}

// NewRay constructs a new ray.
func NewRay(o *Point, d *Vector) *Ray {
	return &Ray{o, d}
}

// Origin returns the origin of this ray.
func (r *Ray) Origin() *Point {
	return r.o
}

// Direction returns the origin of this ray.
func (r *Ray) Direction() *Vector {
	return r.d
}

// Equals returns whether a ray is approximately equal to this ray.
func (r *Ray) Equals(s *Ray) bool {
	return r.Origin().Equals(s.Origin()) && r.Direction().Equals(s.Direction())
}

// Position returns a point representing the position a distance along this ray.
func (r *Ray) Position(t float64) *Point {
	return r.o.PlusVector(r.d.TimesScalar(t))
}
