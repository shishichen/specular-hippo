package tracer

import (
	"math"
	"testing"
)

func TestPlane_Intersect(t *testing.T) {
	var (
		r1 = NewRay(NewPoint(0.0, 10.0, 0.0), NewVector(0.0, 0.0, 1.0))
		r2 = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))
		r3 = NewRay(NewPoint(0.0, 0.0, 10.0), NewVector(0.0, 1.0, 0.0))
		r4 = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 1.0, 0.0))
		r5 = NewRay(NewPoint(0.0, 1.0, 0.0), NewVector(0.0, -1.0, 0.0))
		r6 = NewRay(NewPoint(0.0, -1.0, 0.0), NewVector(0.0, 1.0, 0.0))
		r7 = NewRay(NewPoint(0.0, 0.0, 1.0), NewVector(0.0, 0.0, -1.0))
		r8 = NewRay(NewPoint(0.0, 0.0, -1.0), NewVector(0.0, 0.0, 1.0))
		p1 = NewPlane()
		p2 = NewPlane().WithTransform(NewRotateX(math.Pi / 2.0))
	)
	type args struct {
		r *Ray
	}
	tests := []struct {
		name string
		p    *Plane
		args args
		want Intersections
	}{
		{"case1", p1, args{r1}, NewIntersections()},
		{"case2", p1, args{r2}, NewIntersections()},
		{"case3", p2, args{r3}, NewIntersections()},
		{"case4", p2, args{r4}, NewIntersections()},
		{"case5", p1, args{r5}, NewIntersections(NewIntersection(p1, r5, 1.0))},
		{"case6", p1, args{r6}, NewIntersections(NewIntersection(p1, r6, 1.0))},
		{"case7", p2, args{r7}, NewIntersections(NewIntersection(p2, r7, 1.0))},
		{"case8", p2, args{r8}, NewIntersections(NewIntersection(p2, r8, 1.0))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Intersect(tt.args.r); !got.Equals(tt.want) {
				t.Errorf("Plane.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlane_NormalAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		p    *Plane
		args args
		want *Vector
	}{
		{"case1", NewPlane(), args{NewPoint(0.0, 0.0, 0.0)}, NewVector(0.0, 1.0, 0.0)},
		{"case2", NewPlane(), args{NewPoint(10.0, 0.0, -10.0)}, NewVector(0.0, 1.0, 0.0)},
		{"case3", NewPlane(), args{NewPoint(-5.0, 0.0, 150.0)}, NewVector(0.0, 1.0, 0.0)},
		{"case4", NewPlane().WithTransform(NewRotateX(math.Pi / 2.0)), args{NewPoint(0.0, 0.0, 0.0)}, NewVector(0.0, 0.0, 1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.NormalAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("Plane.NormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
