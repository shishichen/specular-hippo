package tracer

import "math"

// Vector represents a vector.
type Vector [3]float64

// NewVector constructs a new vector.
func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

// X returns the x coordinate.
func (v *Vector) X() float64 {
	return v[0]
}

// Y returns the y coordinate.
func (v *Vector) Y() float64 {
	return v[1]
}

// Z returns the z coordinate.
func (v *Vector) Z() float64 {
	return v[2]
}

// W returns the w coordinate.
func (v *Vector) W() float64 {
	return 0.0
}

// Equals returns whether a vector is approximately equal to this vector.
func (v *Vector) Equals(w *Vector) bool {
	return equals(v.X(), w.X()) && equals(v.Y(), w.Y()) && equals(v.Z(), w.Z())
}

// PlusVector returns a vector representing this vector plus another vector.
func (v *Vector) PlusVector(w *Vector) *Vector {
	return NewVector(v.X()+w.X(), v.Y()+w.Y(), v.Z()+w.Z())
}

// MinusVector returns a vector representing this vector minus another vector.
func (v *Vector) MinusVector(w *Vector) *Vector {
	return NewVector(v.X()-w.X(), v.Y()-w.Y(), v.Z()-w.Z())
}

// Negate returns the negation of this vector.
func (v *Vector) Negate() *Vector {
	return NewVector(0.0-v.X(), 0.0-v.Y(), 0.0-v.Z())
}

// TimesScalar returns this vector multiplied by a scalar.
func (v *Vector) TimesScalar(f float64) *Vector {
	return NewVector(v.X()*f, v.Y()*f, v.Z()*f)
}

// DividedByScalar returns this vector divided by a scalar.
func (v *Vector) DividedByScalar(f float64) *Vector {
	return NewVector(v.X()/f, v.Y()/f, v.Z()/f)
}

// Magnitude returns the length of this vector.
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X()*v.X() + v.Y()*v.Y() + v.Z()*v.Z())
}

// Normalize returns a vector with the same direction as this vector but with length 1.
func (v *Vector) Normalize() *Vector {
	return v.DividedByScalar(v.Magnitude())
}

// DotVector returns the dot product of this vector and another vector.
func (v *Vector) DotVector(w *Vector) float64 {
	return v.X()*w.X() + v.Y()*w.Y() + v.Z()*w.Z()
}

// CrossVector returns the cross product of this vector and another vector.
func (v *Vector) CrossVector(w *Vector) *Vector {
	return NewVector(v.Y()*w.Z()-v.Z()*w.Y(), v.Z()*w.X()-v.X()*w.Z(), v.X()*w.Y()-v.Y()*w.X())
}

// Reflect computes the reflection of a vector around this normal vector.
func (v *Vector) Reflect(w *Vector) *Vector {
	return w.MinusVector(v.TimesScalar(2.0).TimesScalar(w.DotVector(v)))
}
