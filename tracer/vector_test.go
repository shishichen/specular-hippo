package tracer

import (
	"reflect"
	"testing"
)

func TestNewVector(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Vector
	}{
		{"case1", args{4.3, -4.2, 3.1}, &Vector{4.3, -4.2, 3.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVector(tt.args.x, tt.args.y, tt.args.z)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector() = %v, want %v", got, tt.want)
			}
			if got.X() != tt.args.x {
				t.Errorf("x() = %v, want %v", got.X(), tt.args.x)
			}
			if got.Y() != tt.args.y {
				t.Errorf("y() = %v, want %v", got.Y(), tt.args.y)
			}
			if got.Z() != tt.args.z {
				t.Errorf("z() = %v, want %v", got.Z(), tt.args.z)
			}
			if got.W() != 0.0 {
				t.Errorf("w() = %v, want %v", got.W(), 0.0)
			}
		})
	}
}

func TestVector_Equals(t *testing.T) {
	type args struct {
		w *Vector
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want bool
	}{
		{"case1", NewVector(1.0, 2.0, -3.0), args{NewVector(1.0, 2.0, -3.0)}, true},
		{"case2", NewVector(1.0, 2.0, -3.0), args{NewVector(1.000001, 2.0, -3.0)}, true},
		{"case3", NewVector(1.0, 2.0, -3.0), args{NewVector(1.0, 1.999998, -3.0)}, true},
		{"case4", NewVector(1.0, 2.0, -3.0), args{NewVector(1.0, 2.0, -3.000005)}, true},
		{"case5", NewVector(1.0, 2.0, -3.0), args{NewVector(1.000001, 1.999998, -3.000005)}, true},
		{"case6", NewVector(1.0, 2.0, -3.0), args{NewVector(1.0001, 1.999998, -3.000005)}, false},
		{"case7", NewVector(1.0, 2.0, -3.0), args{NewVector(1.000001, 1.99998, -3.000005)}, false},
		{"case8", NewVector(1.0, 2.0, -3.0), args{NewVector(1.000001, 1.999998, -3.00005)}, false},
		{"case8", NewVector(1.0, 2.0, -3.0), args{NewVector(1.0001, 1.99998, -3.00005)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Equals(tt.args.w); got != tt.want {
				t.Errorf("Vector.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_PlusVector(t *testing.T) {
	type args struct {
		w *Vector
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want *Vector
	}{
		{"case1", NewVector(3.0, -2.0, 5.0), args{NewVector(-2.0, 3.0, 1.0)}, NewVector(1.0, 1.0, 6.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.PlusVector(tt.args.w); !got.Equals(tt.want) {
				t.Errorf("Vector.PlusVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_MinusVector(t *testing.T) {
	type args struct {
		w *Vector
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want *Vector
	}{
		{"case1", NewVector(3.0, 2.0, 1.0), args{NewVector(5.0, 6.0, 7.0)}, NewVector(-2.0, -4.0, -6.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.MinusVector(tt.args.w); !got.Equals(tt.want) {
				t.Errorf("Vector.MinusVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Negate(t *testing.T) {
	tests := []struct {
		name string
		v    *Vector
		want *Vector
	}{
		{"arg1", NewVector(1.0, -2.0, 3.0), NewVector(-1.0, 2.0, -3.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Negate(); !got.Equals(tt.want) {
				t.Errorf("Vector.Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_TimesScalar(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want *Vector
	}{
		{"case1", NewVector(1.0, -2.0, 3.0), args{3.5}, NewVector(3.5, -7.0, 10.5)},
		{"case2", NewVector(1.0, -2.0, 3.0), args{0.5}, NewVector(0.5, -1.0, 1.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.TimesScalar(tt.args.f); !got.Equals(tt.want) {
				t.Errorf("Vector.TimesScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_DividedByScalar(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want *Vector
	}{
		{"case1", NewVector(1.0, -2.0, 3.0), args{2.0}, NewVector(0.5, -1.0, 1.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.DividedByScalar(tt.args.f); !got.Equals(tt.want) {
				t.Errorf("Vector.DividedByScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Magnitude(t *testing.T) {
	tests := []struct {
		name string
		v    *Vector
		want float64
	}{
		{"case1", NewVector(1.0, 0.0, 0.0), 1.0},
		{"case2", NewVector(0.0, 1.0, 0.0), 1.0},
		{"case3", NewVector(0.0, 0.0, 1.0), 1.0},
		{"case4", NewVector(1.0, 2.0, 3.0), 3.74166},
		{"case5", NewVector(-1.0, -2.0, -3.0), 3.74166},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Magnitude(); !equals(got, tt.want) {
				t.Errorf("Vector.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Normalize(t *testing.T) {
	tests := []struct {
		name string
		v    *Vector
		want *Vector
	}{
		{"case1", NewVector(4.0, 0.0, 0.0), NewVector(1.0, 0.0, 0.0)},
		{"case2", NewVector(1.0, 2.0, 3.0), NewVector(0.26726, 0.53452, 0.80178)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Normalize(); !(got.Equals(tt.want) && equals(got.Magnitude(), 1.0)) {
				t.Errorf("Vector.Normalize() = %v, want %v; magnitude = %v, want 1.0", got, tt.want, got.Magnitude())
			}
		})
	}
}

func TestVector_DotVector(t *testing.T) {
	type args struct {
		w *Vector
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want float64
	}{
		{"case1", NewVector(1.0, 2.0, 3.0), args{NewVector(2.0, 3.0, 4.0)}, 20.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.DotVector(tt.args.w); !equals(got, tt.want) {
				t.Errorf("Vector.DotVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_CrossVector(t *testing.T) {
	type args struct {
		w *Vector
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want *Vector
	}{
		{"case1", NewVector(1.0, 2.0, 3.0), args{NewVector(2.0, 3.0, 4.0)}, NewVector(-1.0, 2.0, -1.0)},
		{"case2", NewVector(2.0, 3.0, 4.0), args{NewVector(1.0, 2.0, 3.0)}, NewVector(1.0, -2.0, 1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.CrossVector(tt.args.w); !got.Equals(tt.want) {
				t.Errorf("Vector.CrossVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
