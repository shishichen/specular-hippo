package tracer

import (
	"math"
	"testing"
)

func TestNewMatrix2(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
		d float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{-3.0, 5.0, 1.0, -2.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatrix2(tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			if value := got.At(0, 0); !equals(value, -3.0) {
				t.Errorf("value at (0, 0) is not -3.0: %v", value)
			}
			if value := got.At(0, 1); !equals(value, 5.0) {
				t.Errorf("value at (0, 1) is not 5.0: %v", value)
			}
			if value := got.At(1, 0); !equals(value, 1.0) {
				t.Errorf("value at (1, 0) is not 1.0: %v", value)
			}
			if value := got.At(1, 1); !equals(value, -2.0) {
				t.Errorf("value at (1, 1) is not -2.0: %v", value)
			}
		})
	}
}

func TestMatrix2_Equals(t *testing.T) {
	type args struct {
		n *Matrix2
	}
	tests := []struct {
		name string
		m    *Matrix2
		args args
		want bool
	}{
		{"case1", NewMatrix2(1.0, 2.0, 3.0, 4.0), args{NewMatrix2(1.0, 2.0, 3.0, 4.0)}, true},
		{"case2", NewMatrix2(1.0, 2.0, 3.0, 4.0), args{NewMatrix2(2.0, 3.0, 4.0, 5.0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.args.n); got != tt.want {
				t.Errorf("Matrix2.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix2_determinant(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix2
		want float64
	}{
		{"case1", NewMatrix2(1.0, 5.0, -3.0, 2.0), 17.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.determinant(); !equals(got, tt.want) {
				t.Errorf("Matrix2.determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix3(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
		d float64
		e float64
		f float64
		g float64
		h float64
		i float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{-3.0, 5.0, 0.0, 1.0, -2.0, -7.0, 0.0, 1.0, 1.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatrix3(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.e,
				tt.args.f, tt.args.g, tt.args.h, tt.args.i)
			if value := got.At(0, 0); !equals(value, -3.0) {
				t.Errorf("value at (0, 0) is not -3.0: %v", value)
			}
			if value := got.At(1, 1); !equals(value, -2.0) {
				t.Errorf("value at (1, 1) is not -2.0: %v", value)
			}
			if value := got.At(2, 2); !equals(value, 1.0) {
				t.Errorf("value at (2, 2) is not 1.0: %v", value)
			}
		})
	}
}

func TestMatrix3_Equals(t *testing.T) {
	type args struct {
		n *Matrix3
	}
	tests := []struct {
		name string
		m    *Matrix3
		args args
		want bool
	}{
		{"case1", NewMatrix3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0),
			args{NewMatrix3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0)}, true},
		{"case2", NewMatrix3(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0),
			args{NewMatrix3(2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 8.0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.args.n); got != tt.want {
				t.Errorf("Matrix3.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_submatrix(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		m    *Matrix3
		args args
		want *Matrix2
	}{
		{"case1", NewMatrix3(1.0, 5.0, 0.0, -3.0, 2.0, 7.0, 0.0, 6.0, -3.0),
			args{0, 2}, NewMatrix2(-3.0, 2.0, 0.0, 6.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.submatrix(tt.args.x, tt.args.y); !got.Equals(tt.want) {
				t.Errorf("Matrix3.submatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_cofactor(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		m    *Matrix3
		args args
		want float64
	}{
		{"case1", NewMatrix3(3.0, 5.0, 0.0, 2.0, -1.0, -7.0, 6.0, -1.0, 5.0), args{0, 0}, -12.0},
		{"case2", NewMatrix3(3.0, 5.0, 0.0, 2.0, -1.0, -7.0, 6.0, -1.0, 5.0), args{1, 0}, -25.0},
		{"case3", NewMatrix3(1.0, 2.0, 6.0, -5.0, 8.0, -4.0, 2.0, 6.0, 4.0), args{0, 0}, 56.0},
		{"case4", NewMatrix3(1.0, 2.0, 6.0, -5.0, 8.0, -4.0, 2.0, 6.0, 4.0), args{0, 1}, 12.0},
		{"case5", NewMatrix3(1.0, 2.0, 6.0, -5.0, 8.0, -4.0, 2.0, 6.0, 4.0), args{0, 2}, -46.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.cofactor(tt.args.x, tt.args.y); !equals(got, tt.want) {
				t.Errorf("Matrix3.cofactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix3_determinant(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix3
		want float64
	}{
		{"case1", NewMatrix3(1.0, 2.0, 6.0, -5.0, 8.0, -4.0, 2.0, 6.0, 4.0), -196.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.determinant(); !equals(got, tt.want) {
				t.Errorf("Matrix3.determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrix4(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
		d float64
		e float64
		f float64
		g float64
		h float64
		i float64
		j float64
		k float64
		l float64
		m float64
		n float64
		o float64
		p float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{1.0, 2.0, 3.0, 4.0, 5.5, 6.5, 7.5, 8.5, 9.0, 10.0, 11.0, 12.0, 13.5, 14.5, 15.5, 16.5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMatrix4(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.e, tt.args.f, tt.args.g, tt.args.h,
				tt.args.i, tt.args.j, tt.args.k, tt.args.l, tt.args.m, tt.args.n, tt.args.o, tt.args.p)
			if value := got.At(0, 0); !equals(value, 1.0) {
				t.Errorf("value at (0, 0) is not 1.0: %v", value)
			}
			if value := got.At(0, 3); !equals(value, 4.0) {
				t.Errorf("value at (0, 3) is not 4.0: %v", value)
			}
			if value := got.At(1, 0); !equals(value, 5.5) {
				t.Errorf("value at (1, 0) is not 5.5: %v", value)
			}
			if value := got.At(1, 2); !equals(value, 7.5) {
				t.Errorf("value at (1, 2) is not 7.5: %v", value)
			}
			if value := got.At(2, 2); !equals(value, 11.0) {
				t.Errorf("value at (2, 2) is not 11.0: %v", value)
			}
			if value := got.At(3, 0); !equals(value, 13.5) {
				t.Errorf("value at (3, 0) is not 13.5: %v", value)
			}
			if value := got.At(3, 2); !equals(value, 15.5) {
				t.Errorf("value at (3, 2) is not 15.5: %v", value)
			}
		})
	}
}

func TestMatrix4_Equals(t *testing.T) {
	type args struct {
		n *Matrix4
	}
	tests := []struct {
		name string
		m    *Matrix4
		args args
		want bool
	}{
		{"case1", NewMatrix4(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 8.0, 7.0, 6.0, 5.0, 4.0, 3.0, 2.0),
			args{NewMatrix4(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 8.0, 7.0, 6.0, 5.0, 4.0, 3.0, 2.0)}, true},
		{"case2", NewMatrix4(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 8.0, 7.0, 6.0, 5.0, 4.0, 3.0, 2.0),
			args{NewMatrix4(2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 8.0, 7.0, 6.0, 5.0, 4.0, 3.0, 2.0, 1.0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.args.n); got != tt.want {
				t.Errorf("Matrix4.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_timesMatrix(t *testing.T) {
	type args struct {
		n *Matrix4
	}
	tests := []struct {
		name string
		m    *Matrix4
		args args
		want *Matrix4
	}{
		{"case1", NewMatrix4(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 8.0, 7.0, 6.0, 5.0, 4.0, 3.0, 2.0),
			args{NewMatrix4(-2.0, 1.0, 2.0, 3.0, 3.0, 2.0, 1.0, -1.0, 4.0, 3.0, 6.0, 5.0, 1.0, 2.0, 7.0, 8.0)},
			NewMatrix4(20.0, 22.0, 50.0, 48.0, 44.0, 54.0, 114.0, 108.0, 40.0, 58.0, 110.0, 102.0, 16.0, 26.0, 46.0, 42.0)},
		{"case2", NewMatrix4(0.0, 1.0, 2.0, 4.0, 1.0, 2.0, 4.0, 8.0, 2.0, 4.0, 8.0, 16.0, 4.0, 8.0, 16.0, 32.0),
			args{NewIdentity()},
			NewMatrix4(0.0, 1.0, 2.0, 4.0, 1.0, 2.0, 4.0, 8.0, 2.0, 4.0, 8.0, 16.0, 4.0, 8.0, 16.0, 32.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.timesMatrix(tt.args.n); !got.Equals(tt.want) {
				t.Errorf("Matrix4.timesMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_TimesPoint(t *testing.T) {
	type args struct {
		v *Point
	}
	tests := []struct {
		name string
		m    *Matrix4
		args args
		want *Point
	}{
		{"case1", NewMatrix4(1.0, 2.0, 3.0, 4.0, 2.0, 4.0, 4.0, 2.0, 8.0, 6.0, 4.0, 1.0, 0.0, 0.0, 0.0, 1.0),
			args{NewPoint(1.0, 2.0, 3.0)}, NewPoint(18.0, 24.0, 33.0)},
		{"case2", NewIdentity(), args{NewPoint(1.0, 2.0, 3.0)}, NewPoint(1.0, 2.0, 3.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.TimesPoint(tt.args.v); !got.Equals(tt.want) {
				t.Errorf("Matrix4.TimesPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_TimesVector(t *testing.T) {
	type args struct {
		v *Vector
	}
	tests := []struct {
		name string
		m    *Matrix4
		args args
		want *Vector
	}{
		{"case1", NewMatrix4(1.0, 2.0, 3.0, 4.0, 2.0, 4.0, 4.0, 2.0, 8.0, 6.0, 4.0, 1.0, 0.0, 0.0, 0.0, 1.0),
			args{NewVector(1.0, 2.0, 3.0)}, NewVector(14.0, 22.0, 32.0)},
		{"case2", NewIdentity(), args{NewVector(1.0, 2.0, 3.0)}, NewVector(1.0, 2.0, 3.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.TimesVector(tt.args.v); !got.Equals(tt.want) {
				t.Errorf("Matrix4.TimesVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_Transpose(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix4
		want *Matrix4
	}{
		{"case1", NewMatrix4(0.0, 9.0, 3.0, 0.0, 9.0, 8.0, 0.0, 8.0, 1.0, 8.0, 5.0, 3.0, 0.0, 0.0, 5.0, 8.0),
			NewMatrix4(0.0, 9.0, 1.0, 0.0, 9.0, 8.0, 8.0, 0.0, 3.0, 0.0, 5.0, 5.0, 0.0, 8.0, 3.0, 8.0)},
		{"case2", NewIdentity(), NewIdentity()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Transpose(); !got.Equals(tt.want) {
				t.Errorf("Matrix4.Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_submatrix(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		m    *Matrix4
		args args
		want *Matrix3
	}{
		{"case1", NewMatrix4(-6.0, 1.0, 1.0, 6.0, -8.0, 5.0, 8.0, 6.0, -1.0, 0.0, 8.0, 2.0, -7.0, 1.0, -1.0, 1.0),
			args{2, 1}, NewMatrix3(-6.0, 1.0, 6.0, -8.0, 8.0, 6.0, -7.0, -1.0, 1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.submatrix(tt.args.x, tt.args.y); !got.Equals(tt.want) {
				t.Errorf("Matrix4.submatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_cofactor(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		m    *Matrix4
		args args
		want float64
	}{
		{"case1", NewMatrix4(-2.0, -8.0, 3.0, 5.0, -3.0, 1.0, 7.0, 3.0, 1.0, 2.0, -9.0, 6.0, -6.0, 7.0, 7.0, -9.0),
			args{0, 0}, 690.0},
		{"case2", NewMatrix4(-2.0, -8.0, 3.0, 5.0, -3.0, 1.0, 7.0, 3.0, 1.0, 2.0, -9.0, 6.0, -6.0, 7.0, 7.0, -9.0),
			args{0, 1}, 447.0},
		{"case3", NewMatrix4(-2.0, -8.0, 3.0, 5.0, -3.0, 1.0, 7.0, 3.0, 1.0, 2.0, -9.0, 6.0, -6.0, 7.0, 7.0, -9.0),
			args{0, 2}, 210.0},
		{"case4", NewMatrix4(-2.0, -8.0, 3.0, 5.0, -3.0, 1.0, 7.0, 3.0, 1.0, 2.0, -9.0, 6.0, -6.0, 7.0, 7.0, -9.0),
			args{0, 3}, 51.0},
		{"case5", NewMatrix4(-5.0, 2.0, 6.0, -8.0, 1.0, -5.0, 1.0, 8.0, 7.0, 7.0, -6.0, -7.0, 1.0, -3.0, 7.0, 4.0),
			args{2, 3}, -160.0},
		{"case6", NewMatrix4(-5.0, 2.0, 6.0, -8.0, 1.0, -5.0, 1.0, 8.0, 7.0, 7.0, -6.0, -7.0, 1.0, -3.0, 7.0, 4.0),
			args{3, 2}, 105.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.cofactor(tt.args.x, tt.args.y); !equals(got, tt.want) {
				t.Errorf("Matrix4.cofactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_determinant(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix4
		want float64
	}{
		{"case1", NewMatrix4(-2.0, -8.0, 3.0, 5.0, -3.0, 1.0, 7.0, 3.0, 1.0, 2.0, -9.0, 6.0, -6.0, 7.0, 7.0, -9.0), -4071.0},
		{"case2", NewMatrix4(6.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 6.0, 4.0, -9.0, 3.0, -7.0, 9.0, 1.0, 7.0, -6.0), -2120.0},
		{"case3", NewMatrix4(-4.0, 2.0, -2.0, -3.0, 9.0, 6.0, 2.0, 6.0, 0.0, -5.0, 1.0, -5.0, 0.0, 0.0, 0.0, 0.0), 0.0},
		{"case4", NewMatrix4(-5.0, 2.0, 6.0, -8.0, 1.0, -5.0, 1.0, 8.0, 7.0, 7.0, -6.0, -7.0, 1.0, -3.0, 7.0, 4.0), 532.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.determinant(); !equals(got, tt.want) {
				t.Errorf("Matrix4.determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_HasInverse(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix4
		want bool
	}{
		{"case1", NewMatrix4(6.0, 4.0, 4.0, 4.0, 5.0, 5.0, 7.0, 6.0, 4.0, -9.0, 3.0, -7.0, 9.0, 1.0, 7.0, -6.0), true},
		{"case2", NewMatrix4(-4.0, 2.0, -2.0, -3.0, 9.0, 6.0, 2.0, 6.0, 0.0, -5.0, 1.0, -5.0, 0.0, 0.0, 0.0, 0.0), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.HasInverse(); got != tt.want {
				t.Errorf("Matrix4.HasInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_Inverse(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix4
		want *Matrix4
	}{
		{"case1", NewMatrix4(-5.0, 2.0, 6.0, -8.0, 1.0, -5.0, 1.0, 8.0, 7.0, 7.0, -6.0, -7.0, 1.0, -3.0, 7.0, 4.0),
			NewMatrix4(0.21805, 0.45113, 0.24060, -0.04511, -0.80827, -1.45677, -0.44361, 0.52068,
				-0.07895, -0.22368, -0.05263, 0.19737, -0.52256, -0.81391, -0.30075, 0.30639)},
		{"case2", NewMatrix4(8.0, -5.0, 9.0, 2.0, 7.0, 5.0, 6.0, 1.0, -6.0, 0.0, 9.0, 6.0, -3.0, 0.0, -9.0, -4.0),
			NewMatrix4(-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308, 0.02564, 0.03077,
				0.35897, 0.35897, 0.43590, 0.92308, -0.69231, -0.69231, -0.76923, -1.92308)},
		{"case3", NewMatrix4(9.0, 3.0, 0.0, 9.0, -5.0, -2.0, -6.0, -3.0, -4.0, 9.0, 6.0, 4.0, -7.0, 6.0, 6.0, 2.0),
			NewMatrix4(-0.04074, -0.07778, 0.14444, -0.22222, -0.07778, 0.03333, 0.36667, -0.33333,
				-0.02901, -0.14630, -0.10926, 0.12963, 0.17778, 0.06667, -0.26667, 0.33333)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Inverse(); !got.Equals(tt.want) {
				t.Errorf("Matrix4.Inverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestMatrix4_TimesInverse(t *testing.T) {
	tests := []struct {
		name string
		m    *Matrix4
		n    *Matrix4
	}{
		{"case1", NewMatrix4(3.0, -9.0, 7.0, 3.0, 3.0, -8.0, 2.0, -9.0, -4.0, 4.0, 4.0, 1.0, -6.0, 5.0, -1.0, 1.0),
			NewMatrix4(8.0, 2.0, 2.0, 2.0, 3.0, -1.0, 7.0, 0.0, 7.0, 0.0, 5.0, 4.0, 6.0, -2.0, 0.0, 5.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.timesMatrix(tt.n).timesMatrix(tt.n.Inverse()); !got.Equals(tt.m) {
				t.Errorf("Matrix4.TimesInverse() = %v, original %v", got, tt.m)
			}
		})
	}
}

func TestMatrix4_PointTranslation(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(-3.0, 4.0, 5.0), args{NewTranslation(5.0, -3.0, 2.0)}, NewPoint(2.0, 1.0, 7.0)},
		{"case2", NewPoint(-3.0, 4.0, 5.0), args{NewTranslation(5.0, -3.0, 2.0).Inverse()}, NewPoint(-8.0, 7.0, 3.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.TimesPoint(tt.p); !got.Equals(tt.want) {
				t.Errorf("PointTranslation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_VectorTranslation(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want *Vector
	}{
		{"case1", NewVector(-3.0, 4.0, 5.0), args{NewTranslation(5.0, -3.0, 2.0)}, NewVector(-3.0, 4.0, 5.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.TimesVector(tt.v); !got.Equals(tt.want) {
				t.Errorf("VectorTranslation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_PointScaling(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(-4.0, 6.0, 8.0), args{NewScaling(2.0, 3.0, 4.0)}, NewPoint(-8.0, 18.0, 32.0)},
		{"case2", NewPoint(-4.0, 6.0, 8.0), args{NewScaling(2.0, 3.0, 4.0).Inverse()}, NewPoint(-2.0, 2.0, 2.0)},
		{"case3", NewPoint(2.0, 3.0, 4.0), args{NewScaling(-1.0, 1.0, 1.0)}, NewPoint(-2.0, 3.0, 4.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.TimesPoint(tt.p); !got.Equals(tt.want) {
				t.Errorf("PointScaling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_VectorScaling(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		v    *Vector
		args args
		want *Vector
	}{
		{"case1", NewVector(-4.0, 6.0, 8.0), args{NewScaling(2.0, 3.0, 4.0)}, NewVector(-8.0, 18.0, 32.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.TimesVector(tt.v); !got.Equals(tt.want) {
				t.Errorf("VectorScaling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_PointRotation(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(0.0, 1.0, 0.0), args{NewRotationX(math.Pi / 4.0)},
			NewPoint(0.0, math.Sqrt(2)/2.0, math.Sqrt(2)/2.0)},
		{"case2", NewPoint(0.0, 1.0, 0.0), args{NewRotationX(math.Pi / 2.0)},
			NewPoint(0.0, 0.0, 1.0)},
		{"case3", NewPoint(0.0, 1.0, 0.0), args{NewRotationX(math.Pi / 4.0).Inverse()},
			NewPoint(0.0, math.Sqrt(2)/2.0, math.Sqrt(2)/-2.0)},
		{"case4", NewPoint(0.0, 0.0, 1.0), args{NewRotationY(math.Pi / 4.0)},
			NewPoint(math.Sqrt(2)/2.0, 0.0, math.Sqrt(2)/2.0)},
		{"case5", NewPoint(0.0, 0.0, 1.0), args{NewRotationY(math.Pi / 2.0)},
			NewPoint(1.0, 0.0, 0.0)},
		{"case6", NewPoint(0.0, 1.0, 0.0), args{NewRotationZ(math.Pi / 4.0)},
			NewPoint(math.Sqrt(2)/-2.0, math.Sqrt(2)/2.0, 0.0)},
		{"case7", NewPoint(0.0, 1.0, 0.0), args{NewRotationZ(math.Pi / 2.0)},
			NewPoint(-1.0, 0.0, 0.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.TimesPoint(tt.p); !got.Equals(tt.want) {
				t.Errorf("PointRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_PointShearing(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want *Point
	}{
		{"case1", NewPoint(2.0, 3.0, 4.0), args{NewShearing(1.0, 0.0, 0.0, 0.0, 0.0, 0.0)}, NewPoint(5.0, 3.0, 4.0)},
		{"case2", NewPoint(2.0, 3.0, 4.0), args{NewShearing(0.0, 1.0, 0.0, 0.0, 0.0, 0.0)}, NewPoint(6.0, 3.0, 4.0)},
		{"case3", NewPoint(2.0, 3.0, 4.0), args{NewShearing(0.0, 0.0, 1.0, 0.0, 0.0, 0.0)}, NewPoint(2.0, 5.0, 4.0)},
		{"case4", NewPoint(2.0, 3.0, 4.0), args{NewShearing(0.0, 0.0, 0.0, 1.0, 0.0, 0.0)}, NewPoint(2.0, 7.0, 4.0)},
		{"case5", NewPoint(2.0, 3.0, 4.0), args{NewShearing(0.0, 0.0, 0.0, 0.0, 1.0, 0.0)}, NewPoint(2.0, 3.0, 6.0)},
		{"case6", NewPoint(2.0, 3.0, 4.0), args{NewShearing(0.0, 0.0, 0.0, 0.0, 0.0, 1.0)}, NewPoint(2.0, 3.0, 7.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.t.TimesPoint(tt.p); !got.Equals(tt.want) {
				t.Errorf("PointShearing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix4_PointSequence(t *testing.T) {
	type args struct {
		a *Matrix4
		b *Matrix4
		c *Matrix4
	}
	type results struct {
		a *Point
		b *Point
		c *Point
	}
	tests := []struct {
		name string
		p    *Point
		args args
		want results
	}{
		{"case1", NewPoint(1.0, 0.0, 1.0),
			args{NewRotationX(math.Pi / 2.0), NewScaling(5.0, 5.0, 5.0), NewTranslation(10.0, 5.0, 7.0)},
			results{NewPoint(1.0, -1.0, 0.0), NewPoint(5.0, -5.0, 0.0), NewPoint(15.0, 0.0, 7.0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.a.TimesPoint(tt.p)
			if !got.Equals(tt.want.a) {
				t.Errorf("PointSequence() = %v, want %v", got, tt.want.a)
			}
			got = tt.args.b.TimesPoint(got)
			if !got.Equals(tt.want.b) {
				t.Errorf("PointSequence() = %v, want %v", got, tt.want.b)
			}
			got = tt.args.c.TimesPoint(got)
			if !got.Equals(tt.want.c) {
				t.Errorf("PointSequence() = %v, want %v", got, tt.want.c)
			}
			combined := tt.args.a.Concatenate(tt.args.b).Concatenate(tt.args.c)
			if got := combined.TimesPoint(tt.p); !got.Equals(tt.want.c) {
				t.Errorf("PointSequence() = %v, want %v", got, tt.want.c)
			}
		})
	}
}
