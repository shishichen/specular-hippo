package tracer

import (
	"math"
	"sync"
)

// Camera represents a camera in a scene.
type Camera struct {
	transformable
	w           int
	h           int
	fieldOfView float64
	threads     int
}

// NewCamera constructs a new camera.
func NewCamera(w, h int, fieldOfView float64) *Camera {
	return &Camera{defaultTransformable(), w, h, fieldOfView, 1}
}

// Width returns the horizontal size in pixels of this camera's canvas.
func (c *Camera) Width() int {
	return c.w
}

// Height returns the vertical size in pixels of this camera's canvas.
func (c *Camera) Height() int {
	return c.h
}

// FieldOfView returns this camera's field of view.
func (c *Camera) FieldOfView() float64 {
	return c.fieldOfView
}

// WithTransform sets this camera's view transform.
// May return nil if the transform is invalid.
func (c *Camera) WithTransform(t *Matrix4) *Camera {
	if !c.setTransform(t) {
		return nil
	}
	return c
}

// WithTransformFromParameters sets this camera's view transform from human understandable parameters.
func (c *Camera) WithTransformFromParameters(from *Point, to *Point, up *Vector) *Camera {
	forward := to.MinusPoint(from).Normalize()
	left := forward.CrossVector(up.Normalize())
	up = left.CrossVector(forward)
	orientation := NewMatrix4(left.X(), left.Y(), left.Z(), 0.0, up.X(), up.Y(), up.Z(), 0.0,
		-forward.X(), -forward.Y(), -forward.Z(), 0.0, 0.0, 0.0, 0.0, 1.0)
	transform := orientation.TimesMatrix(NewTranslate(-from.X(), -from.Y(), -from.Z()))
	return c.WithTransform(transform)
}

// WithThreads sets the number of threads this camera uses to render the image.
func (c *Camera) WithThreads(threads int) *Camera {
	c.threads = threads
	return c
}

// PixelSize returns the size in world space units of a pixel on this canvas.
func (c *Camera) PixelSize() float64 {
	view := 2.0 * math.Tan(c.FieldOfView()/2.0)
	return math.Min(view/float64(c.Width()), view/float64(c.Height()))
}

// RayForPixel returns a ray starting from this camera and passing through the pixel on the canvas at (x, y).
func (c *Camera) RayForPixel(x, y int) *Ray {
	view := 2.0 * math.Tan(c.FieldOfView()/2.0)
	aspect := float64(c.Width()) / float64(c.Height())
	width, height := view, view
	if aspect >= 1.0 {
		height /= aspect
	} else {
		width *= aspect
	}
	xWorld := width * (0.5 - (float64(x)+0.5)/float64(c.Width()))
	yWorld := height * (0.5 - (float64(y)+0.5)/float64(c.Height()))
	pixel := c.toLocalPoint(NewPoint(xWorld, yWorld, -1.0))
	origin := c.toLocalPoint(NewPoint(0.0, 0.0, 0.0))
	direction := pixel.MinusPoint(origin).Normalize()
	return NewRay(origin, direction)
}

// Render returns an image with this camera's rendering of the world.
func (c *Camera) Render(w *World) *Canvas {
	image := NewCanvas(c.Width(), c.Height())

	var wg sync.WaitGroup
	wg.Add(c.threads)
	for t := 0; t < c.threads; t++ {
		go func(minWidth int, maxWidth int) {
			defer wg.Done()
			for i := minWidth; i < maxWidth; i++ {
				for j := 0; j < c.Height(); j++ {
					r := c.RayForPixel(i, j)
					image.SetColor(i, j, w.ColorAt(r))
				}
			}
		}(int(float64(t)/float64(c.threads)*float64(c.Width())+0.5),
			int(float64(t+1)/float64(c.threads)*float64(c.Width())+0.5))
	}
	wg.Wait()

	return image
}
