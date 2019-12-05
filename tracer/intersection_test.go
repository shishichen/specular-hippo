package tracer

import (
	"testing"
)

func TestNewIntersection(t *testing.T) {
	type args struct {
		r *Ray
		t float64
		s Shape
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0)), 3.5, NewSphere()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIntersection(tt.args.r, tt.args.t, tt.args.s)
			if !got.Ray().Equals(tt.args.r) {
				t.Errorf("Intersection.Ray() = %v, want %v", got.Ray(), tt.args.r)
			}
			if !equals(got.T(), tt.args.t) {
				t.Errorf("Intersection.T() = %v, want %v", got.T(), tt.args.t)
			}
			if got.Shape() != tt.args.s {
				t.Errorf("Intersection.Shape() = %p, want %p", got.Shape(), tt.args.s)
			}
			if got.Point() != nil {
				t.Errorf("Intersection.Point() = %v, want nil", got.Point())
			}
			if got.Normal() != nil {
				t.Errorf("Intersection.Normal() = %v, want nil", got.Normal())
			}
			if got.Eye() != nil {
				t.Errorf("Intersection.Eye() = %v, want nil", got.Eye())
			}
			if got.Inside() != false {
				t.Errorf("Intersection.Inside() = %v, want false", got.Inside())
			}
		})
	}
}

func TestIntersection_Equals(t *testing.T) {
	var (
		r = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))
		s = NewSphere()
	)
	type args struct {
		j *Intersection
	}
	tests := []struct {
		name string
		i    *Intersection
		args args
		want bool
	}{
		{"case1", NewIntersection(r, 3.0, s), args{NewIntersection(r, 3.0, s)}, true},
		{"case2", NewIntersection(NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0)), 3.0, s),
			args{NewIntersection(NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0)), 3.0, s)}, false},
		{"case3", NewIntersection(r, 3.0, s), args{NewIntersection(r, -5.0, s)}, false},
		{"case4", NewIntersection(r, 3.0, NewSphere()), args{NewIntersection(r, 3.0, NewSphere())}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Equals(tt.args.j); got != tt.want {
				t.Errorf("Intersection.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersection_ComputeMetadata(t *testing.T) {
	type want struct {
		point  *Point
		normal *Vector
		eye    *Vector
		inside bool
	}
	tests := []struct {
		name string
		i    *Intersection
		want want
	}{
		{"case1", NewIntersection(NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0)), 4.0, NewSphere()),
			want{NewPoint(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), false}},
		{"case2", NewIntersection(NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0)), 1.0, NewSphere()),
			want{NewPoint(0.0, 0.0, 1.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ComputeMetadata()
			if tt.i.Point() == nil || !tt.i.Point().Equals(tt.want.point) {
				t.Errorf("point = %v, want %v", tt.i.Point(), tt.want.point)
			}
			if tt.i.Normal() == nil || !tt.i.Normal().Equals(tt.want.normal) {
				t.Errorf("normal = %v, want %v", tt.i.Normal(), tt.want.normal)
			}
			if tt.i.Eye() == nil || !tt.i.Eye().Equals(tt.want.eye) {
				t.Errorf("eye = %v, want %v", tt.i.Eye(), tt.want.eye)
			}
			if tt.i.Inside() != tt.want.inside {
				t.Errorf("inside = %v, want %v", tt.i.Inside(), tt.want.inside)
			}
		})
	}
}

func TestIntersection_ShiftedPoint(t *testing.T) {
	type want struct {
		negative bool
		dz       float64
	}
	tests := []struct {
		name string
		i    *Intersection
		want want
	}{
		{"case1", NewIntersection(NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0)), 5.0,
			NewSphere().WithTransform(NewTranslate(0.0, 0.0, 1.0))), want{true, epsilon / 2.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.i.ComputeMetadata()
			if tt.i.ShiftedPoint() == nil ||
				(tt.want.negative && tt.i.ShiftedPoint().Z() >= tt.i.Point().Z()-tt.want.dz) ||
				(!tt.want.negative && tt.i.ShiftedPoint().Z() <= tt.i.Point().Z()+tt.want.dz) {
				t.Errorf("shifted point = %v, want difference from %v of %v", tt.i.ShiftedPoint().Z(), tt.i.Point().Z(), tt.want.dz)
			}
		})
	}
}

func TestNewIntersections(t *testing.T) {
	var (
		r = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))
		s = NewSphere()
	)
	type args struct {
		i1 *Intersection
		i2 *Intersection
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{NewIntersection(r, 1.0, s), NewIntersection(r, 2.0, s)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIntersections(tt.args.i1, tt.args.i2)
			if len(got) != 2 {
				t.Errorf("length = %v, want 2", len(got))
			}
			if !got[0].Equals(tt.args.i1) {
				t.Errorf("intersection at 0 = %v, want %v", got[0], tt.args.i1)
			}
			if !got[1].Equals(tt.args.i2) {
				t.Errorf("intersection at 1 = %v, want %v", got[1], tt.args.i2)
			}
		})
	}
}

func TestIntersections_Equals(t *testing.T) {
	var (
		r = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))
		s = NewSphere()
	)
	type args struct {
		j Intersections
	}
	tests := []struct {
		name string
		i    Intersections
		args args
		want bool
	}{
		{"case1", NewIntersections(NewIntersection(r, 1.0, s), NewIntersection(r, 3.0, s)),
			args{NewIntersections(NewIntersection(r, 1.0, s), NewIntersection(r, 3.0, s))}, true},
		{"case2", NewIntersections(NewIntersection(r, 1.0, s), NewIntersection(r, 3.0, s)),
			args{NewIntersections(NewIntersection(r, 3.0, s), NewIntersection(r, 1.0, s))}, true},
		{"case3", NewIntersections(NewIntersection(r, 1.0, s), NewIntersection(r, 3.0, s)),
			args{NewIntersections(NewIntersection(r, 3.0, s))}, false},
		{"case4", NewIntersections(NewIntersection(r, 1.0, s), NewIntersection(r, 3.0, s)),
			args{NewIntersections(NewIntersection(r, 3.0, s), NewIntersection(r, -5.0, s))}, false},
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
	var (
		r = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))
		s = NewSphere()
	)
	tests := []struct {
		name string
		i    Intersections
		want *Intersection
	}{
		{"case1", NewIntersections(NewIntersection(r, 1.0, s), NewIntersection(r, 2.0, s)), NewIntersection(r, 1.0, s)},
		{"case2", NewIntersections(NewIntersection(r, -1.0, s), NewIntersection(r, 1.0, s)), NewIntersection(r, 1.0, s)},
		{"case3", NewIntersections(NewIntersection(r, -2.0, s), NewIntersection(r, -1.0, s)), nil},
		{"case4", NewIntersections(NewIntersection(r, 5.0, s), NewIntersection(r, 7.0, s), NewIntersection(r, -3.0, s),
			NewIntersection(r, 2.0, s)), NewIntersection(r, 2.0, s)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.Hit(); ((got == nil) != (tt.want == nil)) || (got != nil && !got.Equals(tt.want)) {
				t.Errorf("Intersections.Hit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MergeIntersections(t *testing.T) {
	var (
		r = NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))
		s = NewSphere()
	)
	type args struct {
		i []Intersections
	}
	tests := []struct {
		name string
		args args
		want Intersections
	}{
		{"case1", args{[]Intersections{NewIntersections(NewIntersection(r, 3.0, s)), NewIntersections(NewIntersection(r, 1.0, s))}},
			NewIntersections(NewIntersection(r, 1.0, s), NewIntersection(r, 3.0, s))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntersections(tt.args.i); !got.Equals(tt.want) {
				t.Errorf("MergeIntersections() = %v, want %v", got, tt.want)
			}
		})
	}
}
