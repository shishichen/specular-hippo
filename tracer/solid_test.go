package tracer

import (
	"testing"
)

func TestSolidPattern_ColorAt(t *testing.T) {
	type args struct {
		p *Point
	}
	tests := []struct {
		name string
		s    *SolidPattern
		args args
		want *Color
	}{
		{"case1", NewSolidPattern(white), args{NewPoint(0.0, 5.0, 3.0)}, white},
		{"case2", NewSolidPattern(white), args{NewPoint(85.0, -15.0, 45.0)}, white},
		{"case3", NewSolidPattern(black), args{NewPoint(0.0, 0.0, 0.0)}, black},
		{"case4", NewSolidPattern(black), args{NewPoint(-10.0, 8.0, -2.0)}, black},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("SolidPattern.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
