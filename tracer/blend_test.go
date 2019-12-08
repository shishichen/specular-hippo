package tracer

import (
	"math"
	"testing"
)

func TestBlendedPattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		l    *BlendedPattern
		args args
		want *Color
	}{
		{"case1", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(0.5, 0.0, 0.5)}, NewColor(0.5, 0.5, 0.5)},
		{"case2", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(1.5, 0.0, 0.5)}, black},
		{"case3", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(0.5, 0.0, 1.5)}, white},
		{"case4", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(1.5, 0.0, 1.5)}, NewColor(0.5, 0.5, 0.5)},
		{"case5", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(0.5, 0.0, -0.5)}, white},
		{"case6", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(1.5, 0.0, -0.5)}, NewColor(0.5, 0.5, 0.5)},
		{"case7", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(0.5, 0.0, -1.5)}, NewColor(0.5, 0.5, 0.5)},
		{"case8", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewPoint(1.5, 0.0, -1.5)}, black},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("BlendedPattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBlendedPattern_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
		p *Point
	}
	tests := []struct {
		name string
		l    *BlendedPattern
		args args
		want *Color
	}{
		{"case1", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewRotateZ(math.Pi / 2), NewPoint(0.0, 0.5, 0.5)}, NewColor(0.5, 0.5, 0.5)},
		{"case2", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewRotateZ(math.Pi / 2), NewPoint(0.0, 1.5, 0.5)}, black},
		{"case3", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewRotateZ(math.Pi / 2), NewPoint(0.0, 0.5, 1.5)}, white},
		{"case4", NewBlendedPattern(
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)),
			NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).WithTransform(NewRotateY(math.Pi/2.0))),
			args{NewRotateZ(math.Pi / 2), NewPoint(0.0, 1.5, 1.5)}, NewColor(0.5, 0.5, 0.5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.WithTransform(tt.args.t).ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("BlendedPattern.WithTransform() color = %v, want %v", got, tt.want)
			}
		})
	}
}
