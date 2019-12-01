package tracer

import (
	"fmt"
	"sort"
)

// Intersection represents an intersection.
type Intersection struct {
	t float64
	s *Sphere // treated as opaque pointer
}

// Intersections represents an ordered collection of intersections.
type Intersections []*Intersection

// NewIntersection constructs a new intersection.
func NewIntersection(t float64, s *Sphere) *Intersection {
	return &Intersection{t, s}
}

// T returns the intersection point.
func (i *Intersection) T() float64 {
	return i.t
}

// Sphere returns the sphere intersected with.
func (i *Intersection) Sphere() *Sphere {
	return i.s
}

// Equals returns whether an intersection is approximately equal to this intersection.
func (i *Intersection) Equals(j *Intersection) bool {
	fmt.Printf("%v %v", i.Sphere(), j.Sphere())
	return equals(i.T(), j.T()) && i.Sphere() == j.Sphere()
}

// NewIntersections constructs a new collection of intersections.
func NewIntersections(i []*Intersection) Intersections {
	r := Intersections(i)
	sort.Slice(r, func(i, j int) bool { return r[i].T() < r[j].T() })
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
func (i Intersections) Hit() *Intersection {
	for _, x := range i {
		if x.T() >= 0.0 {
			return x
		}
	}
	return nil
}
