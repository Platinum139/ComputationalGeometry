package InterpolationCurves

import (
	"Graphics/Primitives"
	"math"
)

func BezierCurve(points []Primitives.Point, t float64) Primitives.Point {
	x := 0.0
	y := 0.0
	for i := 0; i < len(points); i++ {
		b := bernsteinPolynomial(i, len(points)-1, t)
		x += b * points[i].X
		y += b * points[i].Y
	}
	return Primitives.Point{X: x, Y: y}
}

func RationalBezierCurve(points []Primitives.Point, w []float64, t float64) Primitives.Point {
	x := 0.0
	y := 0.0
	sum := 0.0
	for i := 0; i < len(points); i++ {
		wb := w[i] * bernsteinPolynomial(i, len(points)-1, t)
		sum += wb
		x += wb * points[i].X
		y += wb * points[i].Y
	}
	return Primitives.Point{X: x / sum, Y: y / sum}
}

func bernsteinPolynomial(k int, n int, t float64) float64 {
	c := float64(combinations(k, n))
	return c * math.Pow(t, float64(k)) * math.Pow(1.0 - t, float64(n - k))
}

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func combinations(k int, n int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}


