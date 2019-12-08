package tracer

import (
	"math"
	"testing"
)

func TestRingPattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		r    *RingPattern
		args args
		want *Color
	}{
		{"case1", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(1.0, 0.0, 0.0)}, black},
		{"case3", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, 1.0)}, black},
		{"case4", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.708, 0.0, 0.708)}, black},
		{"case5", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-1.0, 0.0, 0.0)}, black},
		{"case6", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(0.0, 0.0, -1.0)}, black},
		{"case7", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)), args{NewPoint(-0.708, 0.0, -0.708)}, black},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("RingPattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRingPattern_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
		p *Point
	}
	tests := []struct {
		name string
		r    *RingPattern
		args args
		want *Color
	}{
		{"case1", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)),
			args{NewRotateZ(math.Pi / 2.0), NewPoint(0.0, 0.0, 0.0)}, white},
		{"case2", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)),
			args{NewRotateZ(math.Pi / 2.0), NewPoint(0.0, 1.0, 0.0)}, black},
		{"case3", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)),
			args{NewRotateZ(math.Pi / 2.0), NewPoint(0.0, 0.0, 1.0)}, black},
		{"case4", NewRingPattern(NewSolidPattern(white), NewSolidPattern(black)),
			args{NewRotateZ(math.Pi / 2.0), NewPoint(0.0, 0.708, 0.708)}, black},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.WithTransform(tt.args.t).ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("RingPattern.WithTransform() color = %v, want %v", got, tt.want)
			}
		})
	}
}
