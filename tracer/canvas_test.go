package tracer

import (
	"testing"
)

func TestNewCanvas(t *testing.T) {
	type args struct {
		w int
		h int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{10, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCanvas(tt.args.w, tt.args.h)
			if got.Width() != tt.args.w {
				t.Errorf("Canvas.Width() = %v, want %v", got.Width(), tt.args.w)
			}
			if got.Height() != tt.args.h {
				t.Errorf("Canvas.Height() = %v, want %v", got.Height(), tt.args.h)
			}
			for i := 0; i < got.Width(); i++ {
				for j := 0; j < got.Height(); j++ {
					if color := got.ColorAt(i, j); !color.Equals(black) {
						t.Errorf("color at (%v, %v) = %v, want [0.0, 0.0, 0.0]", i, j, color)
					}
				}
			}
		})
	}
}

func TestCanvas_SetColor(t *testing.T) {
	type args struct {
		x int
		y int
		d *Color
	}
	tests := []struct {
		name string
		c    *Canvas
		args args
	}{
		{"case1", NewCanvas(10, 20), args{2, 3, NewColor(1.0, 0.0, 0.0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetColor(tt.args.x, tt.args.y, tt.args.d)
			if got := tt.c.ColorAt(tt.args.x, tt.args.y); !got.Equals(tt.args.d) {
				t.Errorf("color at (%v, %v) = %v, want %v", tt.args.x, tt.args.y, got, tt.args.d)
			}
		})
	}
}
