package InterpolationCurves

import (
	"Graphics/Primitives"
	"math"
)

func CubicBSpline(points []Primitives.Point, t float64) Primitives.Point {
	n := [4]float64{
		math.Pow(1.0 - t, 3) / 6,
		(3 * math.Pow(t, 3) - 6 * math.Pow(t, 2) + 4 ) / 6,
		(-3 * math.Pow(t, 3) + 3 * math.Pow(t, 2) + 3 * t + 1) / 6,
		math.Pow(t, 3) / 6,
	}
	x := 0.0
	y := 0.0
	for i := range points {
		x += n[i] * points[i].X
		y += n[i] * points[i].Y
	}
	return Primitives.Point{X: x, Y: y}
}

func NonuniformBSpline(k int, t float64) {

}

func FuncN(i int, k int, t float64, tVec []float64) float64 {
	if k == 0 {
		if t < tVec[i] || t >= tVec[i+1] {
			return 0
		}
		return 1
	}
	part1 := (t - tVec[i]) / (tVec[i+k] - tVec[i]) * FuncN(i, k - 1, t, tVec)
	part2 := (tVec[i+k+1] - t) / (tVec[i+k+1] - tVec[i+1]) * FuncN(i+1, k-1, t, tVec)
	return part1 + part2
}
