package tracer

import (
	"testing"
)

func TestNewRay(t *testing.T) {
	type args struct {
		origin    *Point
		direction *Vector
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{NewPoint(1.0, 2.0, 3.0), NewVector(4.0, 5.0, 6.0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRay(tt.args.origin, tt.args.direction)
			if !got.Origin().Equals(tt.args.origin) {
				t.Errorf("Ray.Origin() = %v, want %v", got.o, tt.args.origin)
			}
			if !got.Direction().Equals(tt.args.direction) {
				t.Errorf("Ray.Direction() = %v, want %v", got.d, tt.args.direction)
			}
		})
	}
}

func TestRay_Equals(t *testing.T) {
	type args struct {
		s *Ray
	}
	tests := []struct {
		name string
		r    *Ray
		args args
		want bool
	}{
		{"case1", NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(1.0, 0.0, 0.0)),
			args{NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(1.0, 0.0, 0.0))}, true},
		{"case1", NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(1.0, 0.0, 0.0)),
			args{NewRay(NewPoint(0.0, 0.0, -1.0), NewVector(1.0, 0.0, 0.0))}, false},
		{"case1", NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(1.0, 0.0, 0.0)),
			args{NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(1.0, -2.0, 0.0))}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Equals(tt.args.s); got != tt.want {
				t.Errorf("Ray.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRay_Position(t *testing.T) {
	type args struct {
		t float64
	}
	tests := []struct {
		name string
		r    *Ray
		args args
		want *Point
	}{
		{"case1", NewRay(NewPoint(2.0, 3.0, 4.0), NewVector(1.0, 0.0, 0.0)), args{0.0}, NewPoint(2.0, 3.0, 4.0)},
		{"case2", NewRay(NewPoint(2.0, 3.0, 4.0), NewVector(1.0, 0.0, 0.0)), args{1.0}, NewPoint(3.0, 3.0, 4.0)},
		{"case3", NewRay(NewPoint(2.0, 3.0, 4.0), NewVector(1.0, 0.0, 0.0)), args{-1.0}, NewPoint(1.0, 3.0, 4.0)},
		{"case4", NewRay(NewPoint(2.0, 3.0, 4.0), NewVector(1.0, 0.0, 0.0)), args{2.5}, NewPoint(4.5, 3.0, 4.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Position(tt.args.t); !got.Equals(tt.want) {
				t.Errorf("Ray.Position() = %v, want %v", got, tt.want)
			}
		})
	}
}
