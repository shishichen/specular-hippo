package tracer

import "sort"

// Intersection represents an intersection.
type Intersection struct {
	s        Shape
	r        *Ray
	t        float64
	p        *Point
	normal   *Vector
	eye      *Vector
	inside   bool
	shiftedP *Point
}

// Intersections represents an ordered collection of intersections.
type Intersections []*Intersection

// NewIntersection constructs a new intersection.
func NewIntersection(s Shape, r *Ray, t float64) *Intersection {
	return &Intersection{s, r, t, nil, nil, nil, false, nil}
}

// Shape returns the shape intersected with to create this intersection.
func (i *Intersection) Shape() Shape {
	return i.s
}

// Ray returns the ray used to create this intersection.
func (i *Intersection) Ray() *Ray {
	return i.r
}

// T returns the intersection point on the ray used to create this intersection.
func (i *Intersection) T() float64 {
	return i.t
}

// Point returns the point in world space where this intersection is.
// Returns nil if ComputeMetadata has not yet been called.
func (i *Intersection) Point() *Point {
	return i.p
}

// Normal returns the normal vector at this intersection.
// Returns nil if ComputeMetadata has not yet been called.
func (i *Intersection) Normal() *Vector {
	return i.normal
}

// Eye returns the eye vector at this intersection.
// Returns nil if ComputeMetadata has not yet been called.
func (i *Intersection) Eye() *Vector {
	return i.eye
}

// Inside returns whether the intersection is on the inside of the shape.
// Returns false if ComputeMetadata has not yet been called.
func (i *Intersection) Inside() bool {
	return i.inside
}

// ShiftedPoint returns the point in world space just over actual point.
// Returns nil if ComputeMetadata has not yet been called.
func (i *Intersection) ShiftedPoint() *Point {
	return i.shiftedP
}

// Equals returns whether an intersection is approximately equal to this intersection.
func (i *Intersection) Equals(j *Intersection) bool {
	return i.Ray() == j.Ray() && equals(i.T(), j.T()) && i.Shape() == j.Shape()
}

// ComputeMetadata computes and sets the metadata about this intersection.
func (i *Intersection) ComputeMetadata() {
	i.p = i.Ray().Position(i.T())
	i.normal = i.Shape().NormalAt(i.Point())
	i.eye = i.Ray().Direction().TimesScalar(-1.0)
	i.inside = i.Normal().DotVector(i.Eye()) < 0.0
	if i.Inside() {
		i.normal = i.Normal().TimesScalar(-1.0)
	}
	i.shiftedP = i.p.PlusVector(i.Normal().TimesScalar(epsilon))
}

func sortIntersections(i Intersections) {
	sort.Slice(i, func(x, y int) bool { return i[x].T() < i[y].T() })
}

// NewIntersections constructs a new ordered collection of intersections.
func NewIntersections(is ...*Intersection) Intersections {
	r := Intersections(is)
	sortIntersections(r)
	return r
}

// Equals returns whether a collection of intersections is approximately equal to this collection of intersections.
func (i Intersections) Equals(j Intersections) bool {
	if len(i) != len(j) {
		return false
	}
	for x := range i {
		if !i[x].Equals(j[x]) {
			return false
		}
	}
	return true
}

// Hit returns the hit intersection from a collection of intersections.
// May return nil.
func (i Intersections) Hit() *Intersection {
	for _, x := range i {
		if x.T() >= 0.0 {
			return x
		}
	}
	return nil
}

// MergeIntersections merges a collection of collections of intersections.
func MergeIntersections(i []Intersections) Intersections {
	r := NewIntersections()
	for _, e := range i {
		r = append(r, e...)
	}
	sortIntersections(r)
	return r
}
