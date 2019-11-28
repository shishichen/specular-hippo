package tracer

import (
	"testing"
)

func TestNewColor(t *testing.T) {
	type args struct {
		r float64
		g float64
		b float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{-0.5, 0.4, 1.7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewColor(tt.args.r, tt.args.g, tt.args.b)
			if got.R() != tt.args.r {
				t.Errorf("Color.R() = %v, want %v", got.R(), tt.args.r)
			}
			if got.G() != tt.args.g {
				t.Errorf("Color.G() = %v, want %v", got.G(), tt.args.g)
			}
			if got.B() != tt.args.b {
				t.Errorf("Color.B() = %v, want %v", got.B(), tt.args.b)
			}
		})
	}
}

func TestColor_Equals(t *testing.T) {
	type args struct {
		d *Color
	}
	tests := []struct {
		name string
		c    *Color
		args args
		want bool
	}{
		{"case1", NewColor(1.0, 2.0, -3.0), args{NewColor(1.0, 2.0, -3.0)}, true},
		{"case2", NewColor(1.0, 2.0, -3.0), args{NewColor(1.000001, 2.0, -3.0)}, true},
		{"case3", NewColor(1.0, 2.0, -3.0), args{NewColor(1.0, 1.999998, -3.0)}, true},
		{"case4", NewColor(1.0, 2.0, -3.0), args{NewColor(1.0, 2.0, -3.000005)}, true},
		{"case5", NewColor(1.0, 2.0, -3.0), args{NewColor(1.000001, 1.999998, -3.000005)}, true},
		{"case6", NewColor(1.0, 2.0, -3.0), args{NewColor(1.0001, 1.999998, -3.000005)}, false},
		{"case7", NewColor(1.0, 2.0, -3.0), args{NewColor(1.000001, 1.99998, -3.000005)}, false},
		{"case8", NewColor(1.0, 2.0, -3.0), args{NewColor(1.000001, 1.999998, -3.00005)}, false},
		{"case8", NewColor(1.0, 2.0, -3.0), args{NewColor(1.0001, 1.99998, -3.00005)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Equals(tt.args.d); got != tt.want {
				t.Errorf("Color.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_PlusColor(t *testing.T) {
	type args struct {
		d *Color
	}
	tests := []struct {
		name string
		c    *Color
		args args
		want *Color
	}{
		{"case1", NewColor(0.9, 0.6, 0.75), args{NewColor(0.7, 0.1, 0.25)}, NewColor(1.6, 0.7, 1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.PlusColor(tt.args.d); !got.Equals(tt.want) {
				t.Errorf("Color.PlusColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_MinusColor(t *testing.T) {
	type args struct {
		d *Color
	}
	tests := []struct {
		name string
		c    *Color
		args args
		want *Color
	}{
		{"case1", NewColor(0.9, 0.6, 0.75), args{NewColor(0.7, 0.1, 0.25)}, NewColor(0.2, 0.5, 0.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MinusColor(tt.args.d); !got.Equals(tt.want) {
				t.Errorf("Color.MinusColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_TimesScalar(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		c    *Color
		args args
		want *Color
	}{
		{"case1", NewColor(0.2, 0.3, 0.4), args{2.0}, NewColor(0.4, 0.6, 0.8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.TimesScalar(tt.args.f); !got.Equals(tt.want) {
				t.Errorf("Color.TimesScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_TimesColor(t *testing.T) {
	type args struct {
		d *Color
	}
	tests := []struct {
		name string
		c    *Color
		args args
		want *Color
	}{
		{"case1", NewColor(1.0, 0.2, 0.4), args{NewColor(0.9, 1.0, 0.1)}, NewColor(0.9, 0.2, 0.04)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.TimesColor(tt.args.d); !got.Equals(tt.want) {
				t.Errorf("Color.TimesColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_RGBA(t *testing.T) {
	tests := []struct {
		name  string
		c     *Color
		wantR uint32
		wantG uint32
		wantB uint32
		wantA uint32
	}{
		{"case1", NewColor(-0.5, 0.25, 1.0), 0, 63, 255, 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB, gotA := tt.c.RGBA()
			if gotR != tt.wantR {
				t.Errorf("Color.RGBA() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("Color.RGBA() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("Color.RGBA() gotB = %v, want %v", gotB, tt.wantB)
			}
			if gotA != tt.wantA {
				t.Errorf("Color.RGBA() gotA = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
