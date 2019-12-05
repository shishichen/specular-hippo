package tracer

import "math"

const epsilon float64 = 0.00001

var (
	black = NewColor(0.0, 0.0, 0.0)
)

func equals(x, y float64) bool {
	return math.Abs(x-y) < epsilon
}
