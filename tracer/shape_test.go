package tracer

import (
	"testing"
)

func Test_defaultInternalShape(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defaultInternalShape()
			if !got.Material().Equals(NewDefaultMaterial()) {
				t.Errorf("internalShape.Material() = %v, want default", got.Material())
			}
		})
	}
}

func Test_internalShape_ColorAt(t *testing.T) {
	type fields struct {
		m *Material
		t *Matrix4
	}
	type args struct {
		p *Point
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{"case1",
			fields{NewMaterial(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)), 0.2, 0.9, 0.9, 200.0),
				NewScale(2.0, 2.0, 2.0)},
			args{NewPoint(1.5, 0.0, 0.0)}, white},
		{"case2",
			fields{NewMaterial(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).
				WithTransform(NewScale(2.0, 2.0, 2.0)), 0.2, 0.9, 0.9, 200.0),
				NewIdentity()},
			args{NewPoint(1.5, 0.0, 0.0)}, white},
		{"case3",
			fields{NewMaterial(NewStripePattern(NewSolidPattern(white), NewSolidPattern(black)).
				WithTransform(NewTranslate(0.5, 0.0, 0.0)), 0.2, 0.9, 0.9, 200.0),
				NewScale(2.0, 2.0, 2.0)},
			args{NewPoint(1.5, 0.0, 0.0)}, white},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := defaultInternalShape()
			s.setMaterial(tt.fields.m)
			s.setTransform(tt.fields.t)
			if got := s.ColorAt(tt.args.p); !got.Equals(tt.want) {
				t.Errorf("internalShape.ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_internalShape_setMaterial(t *testing.T) {
	type args struct {
		m *Material
	}
	tests := []struct {
		name string
		s    internalShape
		args args
	}{
		{"case1", defaultInternalShape(), args{NewMaterial(NewSolidPattern(white), 1.0, 0.9, 0.9, 200.0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.setMaterial(tt.args.m)
			if got := tt.s.Material(); !got.Equals(tt.args.m) {
				t.Errorf("material = %v, want %v", got, tt.args.m)
			}
		})
	}
}
