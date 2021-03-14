package InterpolationCurves

import (
	"Graphics/Primitives"
	"math"
)

func TCBSpline(
	p1 Primitives.Point, p2 Primitives.Point,
	q1 Primitives.Point, q2 Primitives.Point,
	t float64) Primitives.Point {

	x := p1.X + q1.X * t +
		(-3 * p1.X + 3 * p2.X - 2 * q1.X - q2.X) * math.Pow(t, 2) +
		(2 * p1.X - 2 * p2.X + q1.X + q2.X) * math.Pow(t, 3)
	y := p1.Y + q1.Y * t +
		(-3 * p1.Y + 3 * p2.Y - 2 * q1.Y - q2.Y) * math.Pow(t, 2) +
		(2 * p1.Y - 2 * p2.Y + q1.Y + q2.Y) * math.Pow(t, 3)
	return Primitives.Point{X: x, Y: y}
}

func StartVec(
	points []Primitives.Point, t float64, c float64, b float64) Primitives.Point {
	m := ((1 - t) * (1 + b) * (1 + c)) / 2
	k := ((1 - t) * (1 - b) * (1 - c)) / 2
	x := m * (points[1].X - points[0].X) + k * (points[2].X - points[1].X)
	y := m * (points[1].Y - points[0].Y) + k * (points[2].Y - points[1].Y)
	return Primitives.Point{X: x, Y: y}
}

func EndVec(
	points []Primitives.Point, t float64, c float64, b float64) Primitives.Point {
	m := ((1 - t) * (1 + b) * (1 - c)) / 2
	k := ((1 - t) * (1 - b) * (1 + c)) / 2
	x := m * (points[1].X - points[0].X) + k * (points[2].X - points[1].X)
	y := m * (points[1].Y - points[0].Y) + k * (points[2].Y - points[1].Y)
	return Primitives.Point{X: x, Y: y}
}