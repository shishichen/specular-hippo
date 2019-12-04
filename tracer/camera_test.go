package tracer

import (
	"math"
	"testing"
)

func TestNewCamera(t *testing.T) {
	type args struct {
		w           int
		h           int
		fieldOfView float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{160, 120, math.Pi / 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCamera(tt.args.w, tt.args.h, tt.args.fieldOfView)
			if got.Width() != tt.args.w {
				t.Errorf("Camera.Width() = %v, want %v", got.Width(), tt.args.w)
			}
			if got.Height() != tt.args.h {
				t.Errorf("Camera.Height() = %v, want %v", got.Height(), tt.args.h)
			}
			if !equals(got.FieldOfView(), tt.args.fieldOfView) {
				t.Errorf("Camera.FieldOfView = %v, want %v", got.FieldOfView(), tt.args.fieldOfView)
			}
			if !got.Transform().Equals(NewIdentity()) {
				t.Errorf("Camera.Transform() = %v, want identity", got.Transform())
			}
		})
	}
}

func TestCamera_WithTransform(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name string
		c    *Camera
		args args
	}{
		{"case1", NewCamera(100, 100, math.Pi/2.0), args{NewTranslate(5.0, 3.0, 2.0).Scale(4.0, 1.0, 2.0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.WithTransform(tt.args.t)
			if !tt.c.Transform().Equals(tt.args.t) {
				t.Errorf("transform = %v, want %v", tt.c.Transform(), tt.args.t)
			}
		})
	}
}

func TestCamera_WithTransformFromParameters(t *testing.T) {
	type args struct {
		from *Point
		to   *Point
		up   *Vector
	}
	tests := []struct {
		name string
		c    *Camera
		args args
		want *Matrix4
	}{
		{"case1", NewCamera(100, 100, math.Pi/2.0),
			args{NewPoint(0.0, 0.0, 0.0), NewPoint(0.0, 0.0, -1.0), NewVector(0.0, 1.0, 0.0)}, NewIdentity()},
		{"case2", NewCamera(100, 100, math.Pi/2.0),
			args{NewPoint(0.0, 0.0, 0.0), NewPoint(0.0, 0.0, 1.0), NewVector(0.0, 1.0, 0.0)}, NewScale(-1.0, 1.0, -1.0)},
		{"case3", NewCamera(100, 100, math.Pi/2.0),
			args{NewPoint(0.0, 0.0, 8.0), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 1.0, 0.0)}, NewTranslate(0.0, 0.0, -8.0)},
		{"case4", NewCamera(100, 100, math.Pi/2.0),
			args{NewPoint(1.0, 3.0, 2.0), NewPoint(4.0, -2.0, 8.0), NewVector(1.0, 1.0, 0.0)},
			NewMatrix4(-0.50709, 0.50709, 0.67612, -2.36643, 0.76772, 0.60609, 0.12122, -2.82843,
				-0.35857, 0.59761, -0.71714, 0.0, 0.0, 0.0, 0.0, 1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.WithTransformFromParameters(tt.args.from, tt.args.to, tt.args.up)
			if !tt.c.Transform().Equals(tt.want) {
				t.Errorf("transform = %v, want %v", tt.c.Transform(), tt.want)
			}
		})
	}
}

func TestCamera_PixelSize(t *testing.T) {
	tests := []struct {
		name string
		c    *Camera
		want float64
	}{
		{"case1", NewCamera(200, 125, math.Pi/2), 0.01},
		{"case2", NewCamera(125, 200, math.Pi/2), 0.01},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.PixelSize(); !equals(got, tt.want) {
				t.Errorf("Camera.PixelSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamera_RayForPixel(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		c    *Camera
		args args
		want *Ray
	}{
		{"case1", NewCamera(201, 101, math.Pi/2), args{100, 50}, NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, -1.0))},
		{"case2", NewCamera(201, 101, math.Pi/2), args{0, 0},
			NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.66519, 0.33259, -0.66851))},
		{"case3", NewCamera(201, 101, math.Pi/2).WithTransform(NewTranslate(0.0, -2.0, 5.0).RotateY(math.Pi / 4.0)),
			args{100, 50}, NewRay(NewPoint(0.0, 2.0, -5.0), NewVector(math.Sqrt(2.0)/2.0, 0.0, math.Sqrt(2.0)/-2.0))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.RayForPixel(tt.args.x, tt.args.y); !got.Equals(tt.want) {
				t.Errorf("Camera.RayForPixel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamera_Render(t *testing.T) {
	type args struct {
		w *World
		x int
		y int
	}
	tests := []struct {
		name string
		c    *Camera
		args args
		want *Color
	}{
		{"case1", NewCamera(11, 11, math.Pi/2.0).WithTransformFromParameters(
			NewPoint(0.0, 0.0, -5.0), NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 1.0, 0.0)),
			args{NewDefaultWorld(), 5, 5}, NewColor(0.38066, 0.47583, 0.2855)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Render(tt.args.w); !got.ColorAt(tt.args.x, tt.args.y).Equals(tt.want) {
				t.Errorf("Camera.Render() = %v, want %v", got.ColorAt(tt.args.x, tt.args.y), tt.want)
			}
		})
	}
}
