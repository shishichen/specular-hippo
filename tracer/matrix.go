package tracer

import (
	"math"
)

// Matrix2 represents a 2x2 matrix.
type Matrix2 [2][2]float64

// Matrix3 represents a 3x3 matrix.
type Matrix3 [3][3]float64

// Matrix4 represents a 4x4 matrix.
type Matrix4 [4][4]float64

// NewMatrix2 constructs a new matrix of size 2.
func NewMatrix2(a, b, c, d float64) *Matrix2 {
	return &Matrix2{{a, b}, {c, d}}
}

// At returns the value at (x, y).
func (m *Matrix2) At(x, y int) float64 {
	return m[x][y]
}

// Equals returns whether a matrix is approximately equal to this matrix.
func (m *Matrix2) Equals(n *Matrix2) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if !equals(m.At(i, j), n.At(i, j)) {
				return false
			}
		}
	}
	return true
}

func (m *Matrix2) determinant() float64 {
	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
}

// NewMatrix3 constructs a new matrix of size 3.
func NewMatrix3(a, b, c, d, e, f, g, h, i float64) *Matrix3 {
	return &Matrix3{{a, b, c}, {d, e, f}, {g, h, i}}
}

// At returns the value at (x, y).
func (m *Matrix3) At(x, y int) float64 {
	return m[x][y]
}

// Equals returns whether a matrix is approximately equal to this matrix.
func (m *Matrix3) Equals(n *Matrix3) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !equals(m.At(i, j), n.At(i, j)) {
				return false
			}
		}
	}
	return true
}

func (m *Matrix3) submatrix(x, y int) *Matrix2 {
	r := Matrix2{}
	for ir, im := 0, 0; ir < 2; ir, im = ir+1, im+1 {
		if im == x {
			im++
		}
		for jr, jm := 0, 0; jr < 2; jr, jm = jr+1, jm+1 {
			if jm == y {
				jm++
			}
			r[ir][jr] = m[im][jm]
		}
	}
	return &r
}

func (m *Matrix3) cofactor(x, y int) float64 {
	r := m.submatrix(x, y).determinant()
	if equals(math.Mod(float64(x+y), 2), 0.0) {
		return r
	}
	return -1.0 * r
}

func (m *Matrix3) determinant() float64 {
	r := 0.0
	for i := 0; i < 3; i++ {
		r += m[0][i] * m.cofactor(0, i)
	}
	return r
}

// NewMatrix4 constructs a new matrix of size 4.
func NewMatrix4(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p float64) *Matrix4 {
	return &Matrix4{{a, b, c, d}, {e, f, g, h}, {i, j, k, l}, {m, n, o, p}}
}

// NewIdentity4 returns a 4x4 identity matrix.
func NewIdentity4() *Matrix4 {
	return NewMatrix4(1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0, 0.0, 0.0, 0.0, 1.0)
}

// At returns the value at (x, y).
func (m *Matrix4) At(x, y int) float64 {
	return m[x][y]
}

// Equals returns whether a matrix is approximately equal to this matrix.
func (m *Matrix4) Equals(n *Matrix4) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if !equals(m.At(i, j), n.At(i, j)) {
				return false
			}
		}
	}
	return true
}

// TimesMatrix returns this matrix multiplied by a matrix.
func (m *Matrix4) TimesMatrix(n *Matrix4) *Matrix4 {
	r := Matrix4{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				r[i][j] += m[i][k] * n[k][j]
			}
		}
	}
	return &r
}

// TimesVector returns a vector representing this matrix multiplied by a vector.
func (m *Matrix4) TimesVector(v *Vector) *Vector {
	r := Vector{}
	for i := 0; i < 3; i++ {
		r[i] = m[i][0]*v.X() + m[i][1]*v.Y() + m[i][2]*v.Z()
	}
	return &r
}

// Transpose returns the tranpose of this matrix.
func (m *Matrix4) Transpose() *Matrix4 {
	r := Matrix4{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r[i][j] = m[j][i]
		}
	}
	return &r
}

func (m *Matrix4) submatrix(x, y int) *Matrix3 {
	r := Matrix3{}
	for ir, im := 0, 0; ir < 3; ir, im = ir+1, im+1 {
		if im == x {
			im++
		}
		for jr, jm := 0, 0; jr < 3; jr, jm = jr+1, jm+1 {
			if jm == y {
				jm++
			}
			r[ir][jr] = m[im][jm]
		}
	}
	return &r
}

func (m *Matrix4) cofactor(x, y int) float64 {
	r := m.submatrix(x, y).determinant()
	if equals(math.Mod(float64(x+y), 2), 0.0) {
		return r
	}
	return -1.0 * r
}

func (m *Matrix4) determinant() float64 {
	r := 0.0
	for i := 0; i < 4; i++ {
		r += m[0][i] * m.cofactor(0, i)
	}
	return r
}

// HasInverse returns whether this matrix is invertible.
func (m *Matrix4) HasInverse() bool {
	return !equals(m.determinant(), 0.0)
}

// Inverse returns the inverse of this matrix.
func (m *Matrix4) Inverse() *Matrix4 {
	if !m.HasInverse() {
		return nil
	}
	d := m.determinant()
	r := Matrix4{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			r[j][i] = m.cofactor(i, j) / d
		}
	}
	return &r
}
