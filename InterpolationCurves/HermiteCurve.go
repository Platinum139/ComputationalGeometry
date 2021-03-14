package InterpolationCurves

import "Graphics/Primitives"

func CubicHermiteCurve(points []Primitives.Point, vectors []Primitives.Point, t float64) Primitives.Point {
	h0 := (1 + 2*t) * (1 - t) * (1 - t)
	h1 := (3 - 2*t) * t * t
	h2 := (1 - t) * (1 - t) * t
	h3 := (t - 1) * t * t
	x := h0 * points[0].X + h1 * points[1].X + h2 * vectors[0].X + h3 * vectors[1].X
	y := h0 * points[0].Y + h1 * points[1].Y + h2 * vectors[0].Y + h3 * vectors[1].Y
	return Primitives.Point{X: x, Y: y}
}
