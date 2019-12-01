package tracer

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
)

// Canvas represents a canvas.
type Canvas struct {
	colors        [][]Color
	width, height int
}

// NewCanvas constructs a new canvas.
func NewCanvas(w, h int) *Canvas {
	black := NewColor(0.0, 0.0, 0.0)
	c := &Canvas{make([][]Color, w), w, h}
	for i := 0; i < w; i++ {
		c.colors[i] = make([]Color, h)
		for j := 0; j < h; j++ {
			c.colors[i][j] = *black
		}
	}
	return c
}

// Width returns the width of the canvas.
func (c *Canvas) Width() int {
	return c.width
}

// Height returns the height of the canvas.
func (c *Canvas) Height() int {
	return c.height
}

func (c *Canvas) contains(x, y int) bool {
	return x >= 0 && x < c.Width() && y >= 0 && y < c.Height()
}

// ColorAt returns the color of the pixel at (x, y).
func (c *Canvas) ColorAt(x, y int) *Color {
	if !c.contains(x, y) {
		return nil
	}
	return &c.colors[x][y]
}

// SetColor sets the color of the pixel at (x, y) to a color.
func (c *Canvas) SetColor(x, y int, d *Color) {
	if !c.contains(x, y) {
		return
	}
	c.colors[x][y] = *d
}

// ToFile converts this canvas to an image file.
func (c *Canvas) ToFile() {
	img := image.NewRGBA(image.Rectangle{image.Point{}, image.Point{c.Width(), c.Height()}})
	for i := 0; i < c.Width(); i++ {
		for j := 0; j < c.Height(); j++ {
			img.Set(i, j, c.ColorAt(i, j))
		}
	}

	f, err := os.Create(fmt.Sprintf("img/i%v.png", time.Now().Unix()))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}
}
