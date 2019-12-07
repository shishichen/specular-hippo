package tracer

import (
	"math"
	"testing"
)

func TestSphere_Intersect(t *testing.T) {
	var (
		r1 = NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0))
		r2 = NewRay(NewPoint(0.0, 1.0, -5.0), NewVector(0.0, 0.0, 1.0))
		r3 = NewRay(NewPoint(0.0, 2.0, -5.0), NewVector(0.0, 0.0, 1.0))
		r4 = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))
		r5 = NewRay(NewPoint(0.0, 0.0, 5.0), NewVector(0.0, 0.0, 1.0))
		s1 = NewSphere()
		s2 = NewSphere().WithTransform(NewScale(2.0, 2.0, 2.0))
		s3 = NewSphere().WithTransform(NewTranslate(5.0, 0.0, 0.0))
	)
	type args struct {
		r *Ray
	}
	tests := []struct {
		name string
		s    *Sphere
		args args
		want Intersections
	}{
		{"case1", s1, args{r1}, NewIntersections(NewIntersection(s1, r1, 4.0), NewIntersection(s1, r1, 6.0))},
		{"case2", s1, args{r2}, NewIntersections(NewIntersection(s1, r2, 5.0), NewIntersection(s1, r2, 5.0))},
		{"case3", s1, args{r3}, NewIntersections()},
		{"case4", s1, args{r4}, NewIntersections(NewIntersection(s1, r4, -1.0), NewIntersection(s1, r4, 1.0))},
		{"case5", s1, args{r5}, NewIntersections(NewIntersection(s1, r5, -6.0), NewIntersection(s1, r5, -4.0))},
		{"case6", s2, args{r1}, NewIntersections(NewIntersection(s2, r1, 3.0), NewIntersection(s2, r1, 7.0))},
		{"case7", s3, args{r1}, NewIntersections()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersect(tt.args.r); !got.Equals(tt.want) {
				t.Errorf("Sphere.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_NormalAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		s    *Sphere
		args args
		want *Vector
	}{
		{"case1", NewSphere(), args{NewPoint(1.0, 0.0, 0.0)}, NewVector(1.0, 0.0, 0.0)},
		{"case2", NewSphere(), args{NewPoint(0.0, 1.0, 0.0)}, NewVector(0.0, 1.0, 0.0)},
		{"case3", NewSphere(), args{NewPoint(0.0, 0.0, 1.0)}, NewVector(0.0, 0.0, 1.0)},
		{"case4", NewSphere(), args{NewPoint(math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0)},
			NewVector(math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0, math.Sqrt(3.0)/3.0)},
		{"case5", NewSphere().WithTransform(NewTranslate(0.0, 1.0, 0.0)), args{NewPoint(0.0, 1.70711, -0.70711)},
			NewVector(0.0, 0.70711, -0.70711)},
		{"case6", NewSphere().WithTransform(NewRotateZ(math.Pi/5.0).Scale(1.0, 0.5, 1.0)),
			args{NewPoint(0.0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/-2.0)}, NewVector(0.0, 0.97014, -0.24254)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NormalAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("Sphere.NormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
