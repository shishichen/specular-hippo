package tracer

import (
	"testing"
)

func TestNewPoint(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{4.3, -4.2, 3.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPoint(tt.args.x, tt.args.y, tt.args.z)
			if got.X() != tt.args.x {
				t.Errorf("Point.X() = %v, want %v", got.X(), tt.args.x)
			}
			if got.Y() != tt.args.y {
				t.Errorf("Point.Y() = %v, want %v", got.Y(), tt.args.y)
			}
			if got.Z() != tt.args.z {
				t.Errorf("Point.Z() = %v, want %v", got.Z(), tt.args.z)
			}
			if got.W() != 1.0 {
				t.Errorf("Point.W() = %v, want %v", got.W(), 1.0)
			}
		})
	}
}

func TestPoint_Equals(t *testing.T) {
	type args struct {
		q *Point
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want bool
	}{
		{"case1", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.0, 2.0, -3.0)}, true},
		{"case2", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.000001, 2.0, -3.0)}, true},
		{"case3", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.0, 1.999998, -3.0)}, true},
		{"case4", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.0, 2.0, -3.000005)}, true},
		{"case5", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.000001, 1.999998, -3.000005)}, true},
		{"case6", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.0001, 1.999998, -3.000005)}, false},
		{"case7", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.000001, 1.99998, -3.000005)}, false},
		{"case8", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.000001, 1.999998, -3.00005)}, false},
		{"case8", NewPoint(1.0, 2.0, -3.0), args{NewPoint(1.0001, 1.99998, -3.00005)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Equals(tt.args.q); got != tt.want {
				t.Errorf("Point.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_PlusVector(t *testing.T) {
	type args struct {
		v *Vector
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(3.0, -2.0, 5.0), args{NewVector(-2.0, 3.0, 1.0)}, NewPoint(1.0, 1.0, 6.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.PlusVector(tt.args.v); !got.Equals(tt.want) {
				t.Errorf("Point.PlusVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_MinusPoint(t *testing.T) {
	type args struct {
		q *Point
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Vector
	}{
		{"case1", NewPoint(3.0, 2.0, 1.0), args{NewPoint(5.0, 6.0, 7.0)}, NewVector(-2.0, -4.0, -6.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.MinusPoint(tt.args.q); !got.Equals(tt.want) {
				t.Errorf("Point.MinusPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_MinusVector(t *testing.T) {
	type args struct {
		v *Vector
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(3.0, 2.0, 1.0), args{NewVector(5.0, 6.0, 7.0)}, NewPoint(-2.0, -4.0, -6.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.MinusVector(tt.args.v); !got.Equals(tt.want) {
				t.Errorf("Point.MinusVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_TimesScalar(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(1.0, -2.0, 3.0), args{3.5}, NewPoint(3.5, -7.0, 10.5)},
		{"case2", NewPoint(1.0, -2.0, 3.0), args{0.5}, NewPoint(0.5, -1.0, 1.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.TimesScalar(tt.args.f); !got.Equals(tt.want) {
				t.Errorf("Point.TimesScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_DividedByScalar(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(1.0, -2.0, 3.0), args{2.0}, NewPoint(0.5, -1.0, 1.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.DividedByScalar(tt.args.f); !got.Equals(tt.want) {
				t.Errorf("Point.DividedByScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}
