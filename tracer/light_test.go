package tracer

import (
	"math"
	"testing"
)

func TestNewLight(t *testing.T) {
	type args struct {
		p *Point
		i *Color
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{NewPoint(0.0, 0.0, 0.0), white}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLight(tt.args.p, tt.args.i)
			if !got.Position().Equals(tt.args.p) {
				t.Errorf("Light.Position() = %v, want %v", got.Position(), tt.args.p)
			}
			if !got.Intensity().Equals(tt.args.i) {
				t.Errorf("Light.Intensity() = %v, want %v", got.Intensity(), tt.args.i)
			}
		})
	}
}

func TestLight_Illuminate(t *testing.T) {
	type args struct {
		s          Shape
		p          *Point
		normal     *Vector
		eye        *Vector
		isShadowed bool
	}
	tests := []struct {
		name string
		l    *Light
		args args
		want *Color
	}{
		{"case1", NewLight(NewPoint(0.0, 0.0, -10.0), white),
			args{NewSphere(), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), false},
			NewColor(1.9, 1.9, 1.9)},
		{"case2", NewLight(NewPoint(0.0, 0.0, -10.0), white),
			args{NewSphere(), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, -1.0),
				NewVector(0.0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/-2.0), false},
			white},
		{"case3", NewLight(NewPoint(0.0, 10.0, -10.0), white),
			args{NewSphere(), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), false},
			NewColor(0.7364, 0.7364, 0.7364)},
		{"case4", NewLight(NewPoint(0.0, 10.0, -10.0), white),
			args{NewSphere(), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, -1.0),
				NewVector(0.0, math.Sqrt(2.0)/-2.0, math.Sqrt(2.0)/-2.0), false},
			NewColor(1.6364, 1.6364, 1.6364)},
		{"case5", NewLight(NewPoint(0.0, 0.0, 10.0), white),
			args{NewSphere(), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), false},
			NewColor(0.1, 0.1, 0.1)},
		{"case6", NewLight(NewPoint(0.0, 0.0, -10.0), white),
			args{NewSphere(), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), true},
			NewColor(0.1, 0.1, 0.1)},
		{"case7", NewLight(NewPoint(0.0, 0.0, -10.0), white),
			args{NewSphere().WithMaterial(
				NewMaterial(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), 1.0, 0.0, 0.0, 200.0)),
				NewPoint(0.9, 0.0, 0.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), false},
			white},
		{"case8", NewLight(NewPoint(0.0, 0.0, -10.0), white),
			args{NewSphere().WithMaterial(
				NewMaterial(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), 1.0, 0.0, 0.0, 200.0)),
				NewPoint(1.1, 0.0, 0.0), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), false},
			black}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Illuminate(tt.args.s, tt.args.p, tt.args.normal, tt.args.eye, tt.args.isShadowed); !got.Equals(tt.want) {
				t.Errorf("Light.Illuminate() = %v, want %v", got, tt.want)
			}
		})
	}
}
