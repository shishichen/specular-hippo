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
		{"case1", args{NewPoint(0.0, 0.0, 0.0), NewColor(1.0, 1.0, 1.0)}},
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

func TestLight_IsShadowed(t *testing.T) {
	type args struct {
		p *Point
		s Shapes
	}
	tests := []struct {
		name string
		l    *Light
		args args
		want bool
	}{
		{"case1", NewLight(NewPoint(-10.0, 10.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(0.0, 10.0, 0.0),
				Shapes{NewSphere().WithMaterial(NewMaterial(NewColor(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0)),
					NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5))}}, false},
		{"case2", NewLight(NewPoint(-10.0, 10.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(10.0, -10.0, 10.0),
				Shapes{NewSphere().WithMaterial(NewMaterial(NewColor(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0)),
					NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5))}}, true},
		{"case3", NewLight(NewPoint(-10.0, 10.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(-20.0, 20.0, -20.0),
				Shapes{NewSphere().WithMaterial(NewMaterial(NewColor(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0)),
					NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5))}}, false},
		{"case4", NewLight(NewPoint(-10.0, 10.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(-2.0, 2.0, -2.0),
				Shapes{NewSphere().WithMaterial(NewMaterial(NewColor(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0)),
					NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5))}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.IsShadowed(tt.args.p, tt.args.s); got != tt.want {
				t.Errorf("Light.IsShadowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLight_Illuminate(t *testing.T) {
	type args struct {
		p      *Point
		m      *Material
		normal *Vector
		eye    *Vector
		s      Shapes
	}
	tests := []struct {
		name string
		l    *Light
		args args
		want *Color
	}{
		{"case1", NewLight(NewPoint(0.0, 0.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(0.0, 0.0, 0.0), NewDefaultMaterial(), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), Shapes{}},
			NewColor(1.9, 1.9, 1.9)},
		{"case2", NewLight(NewPoint(0.0, 0.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(0.0, 0.0, 0.0), NewDefaultMaterial(), NewVector(0.0, 0.0, -1.0),
				NewVector(0.0, math.Sqrt(2.0)/2.0, math.Sqrt(2.0)/-2.0), Shapes{}},
			NewColor(1.0, 1.0, 1.0)},
		{"case3", NewLight(NewPoint(0.0, 10.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(0.0, 0.0, 0.0), NewDefaultMaterial(), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), Shapes{}},
			NewColor(0.7364, 0.7364, 0.7364)},
		{"case4", NewLight(NewPoint(0.0, 10.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(0.0, 0.0, 0.0), NewDefaultMaterial(), NewVector(0.0, 0.0, -1.0),
				NewVector(0.0, math.Sqrt(2.0)/-2.0, math.Sqrt(2.0)/-2.0), Shapes{}},
			NewColor(1.6364, 1.6364, 1.6364)},
		{"case5", NewLight(NewPoint(0.0, 0.0, 10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(0.0, 0.0, 0.0), NewDefaultMaterial(), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0), Shapes{}},
			NewColor(0.1, 0.1, 0.1)},
		{"case6", NewLight(NewPoint(0.0, 0.0, -10.0), NewColor(1.0, 1.0, 1.0)),
			args{NewPoint(0.0, 0.0, 0.0), NewDefaultMaterial(), NewVector(0.0, 0.0, -1.0), NewVector(0.0, 0.0, -1.0),
				Shapes{NewSphere()}},
			NewColor(0.1, 0.1, 0.1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Illuminate(tt.args.m, tt.args.p, tt.args.normal, tt.args.eye, tt.args.s); !got.Equals(tt.want) {
				t.Errorf("Light.Illuminate() = %v, want %v", got, tt.want)
			}
		})
	}
}
