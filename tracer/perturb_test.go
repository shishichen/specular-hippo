package tracer

import (
	"math"
	"testing"
)

func TestPerturbedPattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		e    *PerturbedPattern
		args args
		want *Color
	}{
		{"case1", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewPoint(-0.4, 0.0, 0.0)}, black},
		{"case3", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewPoint(0.0, -0.4, 0.0)}, black},
		{"case4", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewPoint(0.0, 0.0, -0.4)}, white},
		{"case5", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewPoint(0.0, 0.0, 3.1)}, black},
		{"case6", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewPoint(0.0, -0.8, 0.0)}, black},
		{"case7", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewPoint(0.0, 0.0, -1.2)}, white},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("PerturbedPattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerturbedPattern_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
		p *Point
	}
	tests := []struct {
		name string
		e    *PerturbedPattern
		args args
		want *Color
	}{
		{"case1", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewRotateX(math.Pi / 2.0), NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewRotateX(math.Pi / 2.0), NewPoint(-0.4, 0.0, 0.0)}, black},
		{"case3", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewRotateX(math.Pi / 2.0), NewPoint(0.0, -0.4, 0.0)}, black},
		{"case4", NewPerturbedPattern(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), false),
			args{NewRotateX(math.Pi / 2.0), NewPoint(0.0, 0.0, -0.4)}, black},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.WithTransform(tt.args.t).ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("PerturbedPattern.WithTransform() color = %v, want %v", got, tt.want)
			}
		})
	}
}
