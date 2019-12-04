package tracer

import "math"

var (
	black = NewColor(0.0, 0.0, 0.0)
)

func equals(x, y float64) bool {
	const epsilon float64 = 0.00001
	return math.Abs(x-y) < epsilon
}
