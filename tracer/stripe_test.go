package tracer

import (
	"math"
	"testing"
)

func TestStripePattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		s    *StripePattern
		args args
		want *Color
	}{
		{"case1", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 1.0, 0.0)}, white},
		{"case3", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 2.0, 0.0)}, white},
		{"case4", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 1.0)}, white},
		{"case5", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 2.0)}, white},
		{"case6", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.9, 0.0, 0.0)}, white},
		{"case7", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.0, 0.0, 0.0)}, black},
		{"case8", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-0.1, 0.0, 0.0)}, black},
		{"case9", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-1.0, 0.0, 0.0)}, black},
		{"case10", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-1.1, 0.0, 0.0)}, white},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("StripePattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripePattern_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
		p *Point
	}
	tests := []struct {
		name string
		s    *StripePattern
		args args
		want *Color
	}{
		{"case1", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewRotateZ(math.Pi / 2.0),
			NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewRotateZ(math.Pi / 2.0),
			NewPoint(0.0, 1.1, 0.0)}, black},
		{"case3", NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewRotateZ(math.Pi / 2.0),
			NewPoint(0.0, 2.1, 0.0)}, white},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.WithTransform(tt.args.t).ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("StripePattern.WithTransform() color = %v, want %v", got, tt.want)
			}
		})
	}
}
