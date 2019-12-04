package tracer

// World represents a world in a scene.
type World struct {
	s Shapes
	l Lights
}

// NewWorld constructs a new world.
func NewWorld(s Shapes, l Lights) *World {
	return &World{s, l}
}

// NewDefaultWorld constructs a new world with some default objects.
func NewDefaultWorld() *World {
	return NewWorld(
		Shapes{NewSphere().WithMaterial(NewMaterial(NewColor(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200.0)),
			NewSphere().WithTransform(NewScale(0.5, 0.5, 0.5))},
		Lights{NewLight(NewPoint(-10.0, 10.0, -10.0), NewColor(1.0, 1.0, 1.0))})
}

// Shapes returns this world's shapes.
func (w *World) Shapes() Shapes {
	return w.s
}

// Lights returns this world's lights.
func (w *World) Lights() Lights {
	return w.l
}

// Intersect returns this world's intersection points with a ray.
func (w *World) Intersect(r *Ray) Intersections {
	i := NewIntersections()
	for _, s := range w.Shapes() {
		i = i.Merge(s.Intersect(r))
	}
	return i
}

// ColorAt returns the color of a ray's intersection with this world.
func (w *World) ColorAt(r *Ray) *Color {
	c := black
	hit := w.Intersect(r).Hit()
	if hit == nil {
		return c
	}
	hit.ComputeMetadata()
	for _, l := range w.Lights() {
		c = c.PlusColor(l.Illuminate(hit.Shape().Material(), hit.Point(), hit.Normal(), hit.Eye()))
	}
	return c
}
