package tracer

import "math"

func equals(x, y float64) bool {
	const epsilon float64 = 0.00001
	return math.Abs(x-y) < epsilon
}
