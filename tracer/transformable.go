package tracer

type transformable struct {
	transform *Matrix4
	inverse   *Matrix4
}

func defaultTransformable() transformable {
	return transformable{NewIdentity(), NewIdentity()}
}

func (t *transformable) Transform() *Matrix4 {
	return t.transform
}

func (t *transformable) setTransform(m *Matrix4) bool {
	if !m.HasInverse() {
		return false
	}
	t.transform = m
	t.inverse = m.Inverse()
	return true
}

func (t *transformable) toLocalRay(r *Ray) *Ray {
	return t.inverse.TimesRay(r)
}

func (t *transformable) toLocalPoint(p *Point) *Point {
	return t.inverse.TimesPoint(p)
}

func (t *transformable) toWorldVector(v *Vector) *Vector {
	return t.inverse.Transpose().TimesVector(v).Normalize()
}
