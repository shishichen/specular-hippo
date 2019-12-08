package tracer

import (
	"testing"
)

func TestCheckerPattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		c    *CheckerPattern
		args args
		want *Color
	}{
		{"case1", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.99, 0.0, 0.0)}, white},
		{"case3", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.01, 0.0, 0.0)}, black},
		{"case4", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.99, 0.0)}, white},
		{"case5", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 1.01, 0.0)}, black},
		{"case6", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 0.99)}, white},
		{"case7", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 1.01)}, black},
		{"case8", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-0.99, 0.0, 0.0)}, black},
		{"case9", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-1.01, 0.0, 0.0)}, white},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("CheckerPattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckerPattern_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
		p *Point
	}
	tests := []struct {
		name string
		c    *CheckerPattern
		args args
		want *Color
	}{
		{"case1", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.5, 0.0, 0.0), NewPoint(0.0, 0.0, 0.0)}, black},
		{"case2", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.5, 0.0, 0.0), NewPoint(0.49, 0.0, 0.0)}, black},
		{"case3", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.5, 0.0, 0.0), NewPoint(0.51, 0.0, 0.0)}, white},
		{"case4", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.5, 0.0, 0.0), NewPoint(1.01, 0.0, 0.0)}, white},
		{"case5", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.5, 0.0, 0.0), NewPoint(1.51, 0.0, 0.0)}, black},
		{"case6", NewCheckerPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewTranslate(0.5, 0.0, 0.0), NewPoint(-0.51, 0.0, 0.0)}, white},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.WithTransform(tt.args.t).ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("CheckerPattern.WithTransform() color = %v, want %v", got, tt.want)
			}
		})
	}
}
