package tracer

import (
	"testing"
)

func Test_defaultTransformable(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defaultTransformable()
			if !got.transform.Equals(NewIdentity()) {
				t.Errorf("transform = %v, want identity", got.transform)
			}
			if !got.inverse.Equals(NewIdentity()) {
				t.Errorf("inverse = %v, want identity", got.inverse)
			}
		})
	}
}

func Test_transformable_setTransform(t *testing.T) {
	type args struct {
		m *Matrix4
	}
	tests := []struct {
		name    string
		t       transformable
		args    args
		success bool
	}{
		{"case1", defaultTransformable(), args{NewTranslate(2.0, 3.0, 4.0)}, true},
		{"case2", defaultTransformable(), args{NewScale(0.0, 0.0, 0.0)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.setTransform(tt.args.m); got != tt.success {
				t.Errorf("transformable.setTransform() = %v, should succeed: %v", got, tt.success)
			}
			got := tt.t.transform
			if tt.success && !got.Equals(tt.args.m) {
				t.Errorf("transform = %v, want %v", got, tt.args.m)
			} else if !tt.success && !got.Equals(NewIdentity()) {
				t.Errorf("transform = %v, want identity", got)
			}
		})
	}
}
