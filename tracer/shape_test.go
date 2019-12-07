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

			if !got.Transform().Equals(NewIdentity()) {
				t.Errorf("internalShape.Transform() = %v, want identity", got.Transform())
			}
			if !got.Material().Equals(NewDefaultMaterial()) {
				t.Errorf("internalShape.Material() = %v, want default", got.Material())
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
		{"case1", defaultInternalShape(), args{NewMaterial(NewColor(1.0, 1.0, 1.0), 1.0, 0.9, 0.9, 200.0)}},
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

func Test_internalShape_setTransform(t *testing.T) {
	type args struct {
		t *Matrix4
	}
	tests := []struct {
		name    string
		s       internalShape
		args    args
		success bool
	}{
		{"case1", defaultInternalShape(), args{NewTranslate(2.0, 3.0, 4.0)}, true},
		{"case2", defaultInternalShape(), args{NewScale(0.0, 0.0, 0.0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.setTransform(tt.args.t); got != tt.success {
				t.Errorf("internalShape.setTransform() = %v, should succeed: %v", got, tt.success)
			}
			got := tt.s.Transform()
			if tt.success && !got.Equals(tt.args.t) {
				t.Errorf("transform = %v, want %v", got, tt.args.t)
			} else if !tt.success && !got.Equals(NewIdentity()) {
				t.Errorf("transform = %v, want identity", got)
			}
		})
	}
}
