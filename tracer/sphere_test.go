package tracer

import (
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewSphere()
			if !got.Transform().Equals(NewIdentity()) {
				t.Errorf("Sphere.Transform() = %v, want identity", got.Transform())
			}
			if !got.Material().Equals(NewDefaultMaterial()) {
				t.Errorf("Sphere.Material() = %v, want default", got.Material())
			}
		})
	}
}

func TestSphere_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name    string
		s       *Sphere
		args    args
		success bool
	}{
		{"case1", NewSphere(), args{NewTranslate(2.0, 3.0, 4.0)}, true},
		{"case2", NewSphere(), args{NewScale(0.0, 0.0, 0.0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (tt.s.WithTransform(tt.args.t) != nil); got != tt.success {
				t.Errorf("Sphere.WithTransform() = %v, should succeed: %v", got, tt.success)
			}
			got := tt.s.Transform()
			if tt.success && !got.Equals(tt.args.t) {
				t.Errorf("transform = %v, want %v", got, tt.args.t)
			} else if !tt.success && !got.Equals(NewIdentity()) {
				t.Errorf("transform = %v, want identity", got)
			}
		})
	}
}

func TestSphere_WithMaterial(t *testing.T) {
	type args struct {
		m *Material
	}
	tests := []struct {
		name string
		s    *Sphere
		args args
	}{
		{"case1", NewSphere(), args{NewMaterial(NewColor(1.0, 1.0, 1.0), 1.0, 0.9, 0.9, 200.0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.WithMaterial(tt.args.m)
			if got := tt.s.Material(); !got.Equals(tt.args.m) {
				t.Errorf("material = %v, want %v", got, tt.args.m)
			}
		})
	}
}

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
		{"case1", s1, args{r1}, NewIntersections(NewIntersection(r1, 4.0, s1), NewIntersection(r1, 6.0, s1))},
		{"case2", s1, args{r2}, NewIntersections(NewIntersection(r2, 5.0, s1), NewIntersection(r2, 5.0, s1))},
		{"case3", s1, args{r3}, NewIntersections()},
		{"case4", s1, args{r4}, NewIntersections(NewIntersection(r4, -1.0, s1), NewIntersection(r4, 1.0, s1))},
		{"case5", s1, args{r5}, NewIntersections(NewIntersection(r5, -6.0, s1), NewIntersection(r5, -4.0, s1))},
		{"case6", s2, args{r1}, NewIntersections(NewIntersection(r1, 3.0, s2), NewIntersection(r1, 7.0, s2))},
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
