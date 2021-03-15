package OrientationTests

import (
	"Graphics/Primitives"
	"math"
)

func AngleTest(polyline []Primitives.Point, point Primitives.Point) int {
	angleSum := 0.0
	for i := 0; i < len(polyline) - 1; i++ {
		v := Primitives.Point{
			X: polyline[i].X - point.X,
			Y: polyline[i].Y - point.Y,
		}
		w := Primitives.Point{
			X: polyline[i+1].X - point.X,
			Y: polyline[i+1].Y - point.Y,
		}
		if isZeroVector(v) || isZeroVector(w) {
			return 2
		}
		fi := angle(v, w) * 180 / math.Pi
		if math.Abs(fi - 180) < 0.1 {
			return 3
		}
		angleSum += fi
	}
	if angleSum <= 180 {
		return 0
	}
	return 1
}

func isZeroVector(v Primitives.Point) bool {
	if v.X == 0 && v.Y == 0 {
		return true
	}
	return false
}

func angle(v1 Primitives.Point, v2 Primitives.Point) float64 {
	dot := (v1.X * v2.X + v1.Y * v2.Y) * 1.0
	d1 := math.Sqrt(v1.X * v1.X + v1.Y * v1.Y)
	d2 := math.Sqrt(v2.X * v2.X + v2.Y * v2.Y)
	sign := 1.0
	if v1.X * v2.Y - v1.Y * v2.X < 0 {
		sign = -1.0
	}
	return sign * math.Acos(dot / (d1 * d2))
}