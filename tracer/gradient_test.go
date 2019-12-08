package tracer

import (
	"math"
	"testing"
)

func TestGradientPattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		g    *GradientPattern
		args args
		want *Color
	}{
		{"case1", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.25, 0.0, 0.0)},
			NewColor(0.75, 0.75, 0.75)},
		{"case3", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.5, 0.0, 0.0)},
			NewColor(0.5, 0.5, 0.5)},
		{"case4", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.75, 0.0, 0.0)},
			NewColor(0.25, 0.25, 0.25)},
		{"case5", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.25, 0.0, 0.0)},
			NewColor(0.25, 0.25, 0.25)},
		{"case6", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.5, 0.0, 0.0)},
			NewColor(0.5, 0.5, 0.5)},
		{"case7", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.75, 0.0, 0.0)},
			NewColor(0.75, 0.75, 0.75)},
		{"case8", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-0.25, 0.0, 0.0)},
			NewColor(0.75, 0.75, 0.75)},
		{"case9", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-0.5, 0.0, 0.0)},
			NewColor(0.5, 0.5, 0.5)},
		{"case10", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-0.75, 0.0, 0.0)},
			NewColor(0.25, 0.25, 0.25)},
		{"case11", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-2.25, 0.0, 0.0)},
			NewColor(0.75, 0.75, 0.75)},
		{"case12", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-2.5, 0.0, 0.0)},
			NewColor(0.5, 0.5, 0.5)},
		{"case13", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-2.75, 0.0, 0.0)},
			NewColor(0.25, 0.25, 0.25)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("GradientPattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGradientPattern_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
		p *Point
	}
	tests := []struct {
		name string
		g    *GradientPattern
		args args
		want *Color
	}{
		{"case1", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewRotateZ(math.Pi / 2.0),
			NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewRotateZ(math.Pi / 2.0),
			NewPoint(0.0, 0.25, 0.0)}, NewColor(0.75, 0.75, 0.75)},
		{"case3", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewRotateZ(math.Pi / 2.0),
			NewPoint(0.0, 1.25, 0.0)}, NewColor(0.25, 0.25, 0.25)},
		{"case4", NewGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewRotateZ(math.Pi / 2.0),
			NewPoint(0.0, -2.25, 0.0)}, NewColor(0.75, 0.75, 0.75)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.WithTransform(tt.args.t).ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("GradientPattern.WithTransform() color = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadialGradientPattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		r    *RadialGradientPattern
		args args
		want *Color
	}{
		{"case1", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.25, 0.0, 0.0)},
			NewColor(0.75, 0.75, 0.75)},
		{"case3", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.5, 0.0, 0.0)},
			NewColor(0.5, 0.5, 0.5)},
		{"case4", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.75, 0.0, 0.0)},
			NewColor(0.25, 0.25, 0.25)},
		{"case5", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.25, 0.0, 0.0)},
			NewColor(0.25, 0.25, 0.25)},
		{"case6", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.5, 0.0, 0.0)},
			NewColor(0.5, 0.5, 0.5)},
		{"case7", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.75, 0.0, 0.0)},
			NewColor(0.75, 0.75, 0.75)},
		{"case8", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, -0.25)},
			NewColor(0.75, 0.75, 0.75)},
		{"case9", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, -0.5)},
			NewColor(0.5, 0.5, 0.5)},
		{"case10", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, -0.75)},
			NewColor(0.25, 0.25, 0.25)},
		{"case11", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, -2.25)},
			NewColor(0.75, 0.75, 0.75)},
		{"case12", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, -2.5)},
			NewColor(0.5, 0.5, 0.5)},
		{"case13", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, -2.75)},
			NewColor(0.25, 0.25, 0.25)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("RadialGradientPattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRadialGradientPattern_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
		p *Point
	}
	tests := []struct {
		name string
		r    *RadialGradientPattern
		args args
		want *Color
	}{
		{"case1", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.0, 0.0, 0.5),
			NewPoint(0.0, 0.0, 0.0)}, NewColor(0.5, 0.5, 0.5)},
		{"case2", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.0, 0.0, 0.5),
			NewPoint(0.0, 0.0, 0.25)}, NewColor(0.75, 0.75, 0.75)},
		{"case3", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.0, 0.0, 0.5),
			NewPoint(0.0, 0.0, 1.25)}, NewColor(0.25, 0.25, 0.25)},
		{"case4", NewRadialGradientPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.0, 0.0, 0.5),
			NewPoint(0.0, 0.0, -2.25)}, NewColor(0.25, 0.25, 0.25)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.WithTransform(tt.args.t).ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("RadialGradientPattern.WithTransform() color = %v, want %v", got, tt.want)
			}
		})
	}
}
