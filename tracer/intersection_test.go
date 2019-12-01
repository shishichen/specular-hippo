package tracer

import (
	"testing"
)

func TestNewIntersection(t *testing.T) {
	type args struct {
		t float64
		s *Sphere
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{3.5, NewSphere()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIntersection(tt.args.t, tt.args.s)
			if !equals(got.T(), tt.args.t) {
				t.Errorf("Intersection.T() = %v, want %v", got.T(), tt.args.t)
			}
			if got.Sphere() != tt.args.s {
				t.Errorf("Intersection.Sphere() = %p, want %p", got.Sphere(), tt.args.s)
			}
		})
	}
}

func TestIntersection_Equals(t *testing.T) {
	s := NewSphere()
	type args struct {
		j *Intersection
	}
	tests := []struct {
		name string
		i    *Intersection
		args args
		want bool
	}{
		{"case1", NewIntersection(3.0, s), args{NewIntersection(3.0, s)}, true},
		{"case2", NewIntersection(3.0, s), args{NewIntersection(-5.0, s)}, false},
		{"case3", NewIntersection(3.0, NewSphere()), args{NewIntersection(3.0, NewSphere())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equals(tt.args.j); got != tt.want {
				t.Errorf("Intersection.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIntersections(t *testing.T) {
	s := NewSphere()
	type args struct {
		i []*Intersection
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{[]*Intersection{NewIntersection(1.0, s), NewIntersection(2.0, s)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIntersections(tt.args.i)
			if len(got) != len(tt.args.i) {
				t.Errorf("length = %v, want %v", len(got), len(tt.args.i))
			}
			for i := range got {
				if !got[i].Equals(tt.args.i[i]) {
					t.Errorf("intersection at %v = %v, want %v", i, got[i], tt.args.i[i])
				}
			}
		})
	}
}

func TestIntersections_Equals(t *testing.T) {
	s := NewSphere()
	type args struct {
		j Intersections
	}
	tests := []struct {
		name string
		i    Intersections
		args args
		want bool
	}{
		{"case1", NewIntersections([]*Intersection{NewIntersection(1.0, s), NewIntersection(3.0, s)}),
			args{NewIntersections([]*Intersection{NewIntersection(1.0, s), NewIntersection(3.0, s)})}, true},
		{"case2", NewIntersections([]*Intersection{NewIntersection(1.0, s), NewIntersection(3.0, s)}),
			args{NewIntersections([]*Intersection{NewIntersection(3.0, s), NewIntersection(1.0, s)})}, true},
		{"case3", NewIntersections([]*Intersection{NewIntersection(1.0, s), NewIntersection(3.0, s)}),
			args{NewIntersections([]*Intersection{NewIntersection(3.0, s)})}, false},
		{"case4", NewIntersections([]*Intersection{NewIntersection(1.0, s), NewIntersection(3.0, s)}),
			args{NewIntersections([]*Intersection{NewIntersection(3.0, s), NewIntersection(-5.0, s)})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equals(tt.args.j); got != tt.want {
				t.Errorf("Intersections.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersections_Hit(t *testing.T) {
	s := NewSphere()
	tests := []struct {
		name string
		i    Intersections
		want *Intersection
	}{
		{"case1", NewIntersections([]*Intersection{NewIntersection(1.0, s), NewIntersection(2.0, s)}),
			NewIntersection(1.0, s)},
		{"case2", NewIntersections([]*Intersection{NewIntersection(-1.0, s), NewIntersection(1.0, s)}),
			NewIntersection(1.0, s)},
		{"case3", NewIntersections([]*Intersection{NewIntersection(-2.0, s), NewIntersection(-1.0, s)}), nil},
		{"case4", NewIntersections([]*Intersection{NewIntersection(5.0, s), NewIntersection(7.0, s),
			NewIntersection(-3.0, s), NewIntersection(2.0, s)}), NewIntersection(2.0, s)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Hit(); ((got == nil) != (tt.want == nil)) || (got != nil && !got.Equals(tt.want)) {
				t.Errorf("Intersections.Hit() = %v, want %v", got, tt.want)
			}
		})
	}
}
