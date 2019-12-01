package tracer

import (
	"testing"
)

func TestNewMaterial(t *testing.T) {
	type args struct {
		c         *Color
		ambient   float64
		diffuse   float64
		specular  float64
		shininess float64
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{NewColor(1.0, 3.0, 1.0), 0.1, 0.9, 0.5, 200.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMaterial(tt.args.c, tt.args.ambient, tt.args.diffuse, tt.args.specular, tt.args.shininess)
			if !got.Color().Equals(tt.args.c) {
				t.Errorf("Material.Color() = %v, want %v", got.Color(), tt.args.c)
			}
			if !equals(got.Ambient(), tt.args.ambient) {
				t.Errorf("Material.Ambient() = %v, want %v", got.Ambient(), tt.args.ambient)
			}
			if !equals(got.Diffuse(), tt.args.diffuse) {
				t.Errorf("Material.Diffusee() = %v, want %v", got.Diffuse(), tt.args.diffuse)
			}
			if !equals(got.Specular(), tt.args.specular) {
				t.Errorf("Material.specular() = %v, want %v", got.Specular(), tt.args.specular)
			}
			if !equals(got.Shininess(), tt.args.shininess) {
				t.Errorf("Material.Shininess() = %v, want %v", got.Shininess(), tt.args.shininess)
			}
		})
	}
}

func TestMaterial_Equals(t *testing.T) {
	type args struct {
		n *Material
	}
	tests := []struct {
		name string
		m    *Material
		args args
		want bool
	}{
		{"case1", NewMaterial(NewColor(1.0, 1.0, 1.0), 0.5, 0.7, 0.9, 100.0),
			args{NewMaterial(NewColor(1.0, 1.0, 1.0), 0.5, 0.7, 0.9, 100.0)}, true},
		{"case2", NewDefaultMaterial(), args{NewMaterial(NewColor(1.0, 1.0, 1.0), 0.1, 0.9, 0.9, 200.0)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.args.n); got != tt.want {
				t.Errorf("Material.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
