package tracer

import (
	"testing"
)

func TestNewWorld(t *testing.T) {
	type args struct {
		s Shapes
		l Lights
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{Shapes{}, Lights{}}},
		{"case2", args{
			Shapes{NewSphere().WithMaterial(NewMaterial(NewSolidPattern(NewColor(0.8, 1.0, 0.6)), 0.1, 0.7, 0.2, 200.0)),
				NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5))},
			Lights{NewLight(NewPoint(-10.0, 10.0, -10.0), white)}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewWorld(tt.args.s, tt.args.l)
			if len(got.Shapes()) != len(tt.args.s) {
				t.Errorf("World.Shapes() length = %v, want %v", len(got.Shapes()), len(tt.args.s))
			}
			for i := range got.Shapes() {
				if got.Shapes()[i] != tt.args.s[i] {
					t.Errorf("World.Shapes() at %v = %v, want %v", i, got.Shapes()[i], tt.args.s[i])
				}
			}
			if len(got.Lights()) != len(tt.args.l) {
				t.Errorf("World.Lights() length = %v, want %v", len(got.Lights()), len(tt.args.l))
			}
			for i := range got.Lights() {
				if got.Lights()[i] != tt.args.l[i] {
					t.Errorf("World.Lights() at %v = %v, want %v", i, got.Lights()[i], tt.args.l[i])
				}
			}
		})
	}
}

func TestWorld_Intersect(t *testing.T) {
	var (
		w = NewDefaultWorld()
		r = NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0))
	)
	type args struct {
		r *Ray
	}
	tests := []struct {
		name string
		w    *World
		args args
		want Intersections
	}{
		{"case1", w, args{r}, NewIntersections(
			NewIntersection(w.Shapes()[0], r, 4.0), NewIntersection(w.Shapes()[1], r, 4.5),
			NewIntersection(w.Shapes()[1], r, 5.5), NewIntersection(w.Shapes()[0], r, 6.0))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Intersect(tt.args.r); !got.Equals(tt.want) {
				t.Errorf("World.Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_IsShadowed(t *testing.T) {
	type args struct {
		p *Point
		l *Light
	}
	tests := []struct {
		name string
		w    *World
		args args
		want bool
	}{
		{"case1", NewDefaultWorld(), args{NewPoint(0.0, 10.0, 0.0), NewLight(NewPoint(-10.0, 10.0, -10.0), white)}, false},
		{"case2", NewDefaultWorld(), args{NewPoint(10.0, -10.0, 10.0), NewLight(NewPoint(-10.0, 10.0, -10.0), white)}, true},
		{"case3", NewDefaultWorld(), args{NewPoint(-20.0, 20.0, -20.0), NewLight(NewPoint(-10.0, 10.0, -10.0), white)}, false},
		{"case4", NewDefaultWorld(), args{NewPoint(-2.0, 2.0, -2.0), NewLight(NewPoint(-10.0, 10.0, -10.0), white)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.IsShadowed(tt.args.p, tt.args.l); got != tt.want {
				t.Errorf("World.IsShadowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorld_ColorAt(t *testing.T) {
	type args struct {
		r *Ray
	}
	tests := []struct {
		name string
		w    *World
		args args
		want *Color
	}{
		{"case1", NewDefaultWorld(), args{NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 1.0, 0.0))}, black},
		{"case2", NewDefaultWorld(), args{NewRay(NewPoint(0.0, 0.0, -5.0), NewVector(0.0, 0.0, 1.0))},
			NewColor(0.38066, 0.47583, 0.2855)},
		{"case3", NewWorld(
			Shapes{NewSphere().WithMaterial(NewMaterial(NewSolidPattern(NewColor(0.8, 1.0, 0.6)), 0.1, 0.7, 0.2, 200.0)),
				NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5))},
			Lights{NewLight(NewPoint(0.0, 0.25, 0.0), white)}),
			args{NewRay(NewPoint(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 1.0))}, NewColor(0.90498, 0.90498, 0.90498)},
		{"case4", NewWorld(
			Shapes{NewSphere().WithMaterial(NewMaterial(NewSolidPattern(NewColor(0.8, 1.0, 0.6)), 1.0, 0.7, 0.2, 200.0)),
				NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5)).
					WithMaterial(NewMaterial(NewSolidPattern(white), 1.0, 0.9, 0.9, 200.0))},
			Lights{NewLight(NewPoint(-10.0, 10.0, -10.0), white)}),
			args{NewRay(NewPoint(0.0, 0.0, 0.75), NewVector(0.0, 0.0, -1.0))}, white},
		{"case5", NewWorld(
			Shapes{NewSphere(), NewSphere().WithTransform(NewTranslate(0.0, 0.0, 10.0))},
			Lights{NewLight(NewPoint(0.0, 0.0, -10.0), white)}),
			args{NewRay(NewPoint(0.0, 0.0, 5.0), NewVector(0.0, 0.0, 1.0))}, NewColor(0.1, 0.1, 0.1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.ColorAt(tt.args.r); !got.Equals(tt.want) {
				t.Errorf("World.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
