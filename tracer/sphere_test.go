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

func TestSphere_SetTransform(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		s    *Sphere
		args args
		want bool
	}{
		{"case1", NewSphere(), args{NewTranslation(2.0, 3.0, 4.0)}, true},
		{"case2", NewSphere(), args{NewScaling(0.0, 0.0, 0.0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SetTransform(tt.args.t); got != tt.want {
				t.Errorf("Sphere.SetTransform() = %v, want %v", got, tt.want)
			}
			got := tt.s.Transform()
			if tt.want && !got.Equals(tt.args.t) {
				t.Errorf("transform = %v, want %v", got, tt.args.t)
			} else if !tt.want && !got.Equals(NewIdentity()) {
				t.Errorf("transform = %v, want %v", got, tt.args.t)
			}
		})
	}
}

func TestSphere_SetMaterial(t *testing.T) {
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
			tt.s.SetMaterial(tt.args.m)
			if got := tt.s.Material(); !got.Equals(tt.args.m) {
				t.Errorf("material = %v, want %v", got, tt.args.m)
			}
		})
	}
}

func TestSphere_Intersects(t *testing.T) {
	s := NewSphere()
	type args struct {
		r *Ray
	}
	tests := []struct {
		name string
		s    *Sphere
		args args
		want Intersections
	}{
		{"case1", s, args{NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0))},
			NewIntersections([]*Intersection{NewIntersection(4.0, s), NewIntersection(6.0, s)})},
		{"case2", s, args{NewRay(NewPoint(0.0, 1.0, -5.0), NewVector(0.0, 0.0, 1.0))},
			NewIntersections([]*Intersection{NewIntersection(5.0, s), NewIntersection(5.0, s)})},
		{"case3", s, args{NewRay(NewPoint(0.0, 2.0, -5.0), NewVector(0.0, 0.0, 1.0))},
			NewIntersections([]*Intersection{})},
		{"case4", s, args{NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))},
			NewIntersections([]*Intersection{NewIntersection(-1.0, s), NewIntersection(1.0, s)})},
		{"case5", s, args{NewRay(NewPoint(0.0, 0.0, 5.0), NewVector(0.0, 0.0, 1.0))},
			NewIntersections([]*Intersection{NewIntersection(-6.0, s), NewIntersection(-4.0, s)})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Intersects(tt.args.r); !got.Equals(tt.want) {
				t.Errorf("Sphere.Intersects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_TransformedIntersects(t *testing.T) {
	s := NewSphere()
	type args struct {
		r *Ray
	}
	tests := []struct {
		name string
		s    *Sphere
		t    *Matrix4
		args args
		want Intersections
	}{
		{"case1", s, NewScaling(2.0, 2.0, 2.0), args{NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0))},
			NewIntersections([]*Intersection{NewIntersection(3.0, s), NewIntersection(7.0, s)})},
		{"case2", s, NewTranslation(5.0, 0.0, 0.0), args{NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0))},
			NewIntersections([]*Intersection{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.s.SetTransform(tt.t) {
				t.Errorf("unable to set transform to %v", tt.t)
			}
			if got := tt.s.Intersects(tt.args.r); !got.Equals(tt.want) {
				t.Errorf("Sphere.TransformedIntersects() = %v, want %v", got, tt.want)
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.NormalAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("Sphere.NormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_TransformedNormalAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		s    *Sphere
		t    *Matrix4
		args args
		want *Vector
	}{
		{"case1", NewSphere(), NewTranslation(0.0, 1.0, 0.0), args{NewPoint(0.0, 1.70711, -0.70711)}, NewVector(0.0, 0.70711, -0.70711)},
		{"case2", NewSphere(), NewRotationZ(math.Pi / 5.0).Concatenate(NewScaling(1.0, 0.5, 1.0)),
			args{NewPoint(0.0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/-2.0)}, NewVector(0.0, 0.97014, -0.24254)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.s.SetTransform(tt.t) {
				t.Errorf("unable to set transform to %v", tt.t)
			}
			if got := tt.s.NormalAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("Sphere.TransformedNormalAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
