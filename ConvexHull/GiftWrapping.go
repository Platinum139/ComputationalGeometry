package ConvexHull

import (
	"Graphics/Primitives"
	"math"
)

func GiftWrapping(points []Primitives.Point) []Primitives.Point {
	// find extreme point
	idx := 0
	for i := range points {
		if points[i].Y < points[idx].Y || (points[i].Y == points[idx].Y && points[i].X < points[idx].X) {
			idx = i
		}
	}
	points[idx], points[0] = points[0], points[idx]
	points = append(points, points[0])

	v := Primitives.Point{X: 10, Y: 0}
	k := 0
	for {
		cMax := -100.0
		idx := -1

		if k != 0 {
			v = Primitives.Point{
				X: points[k].X - points[k-1].X,
				Y: points[k].Y - points[k-1].Y,
			}
		}
		for i := k + 1; i < len(points); i++ {
			w := Primitives.Point{
				X: points[i].X - points[k].X,
				Y: points[i].Y - points[k].Y,
			}
			if !(w.X == 0 && w.Y == 0) {
				c := cosFi(v, w)
				if c >= cMax {
					cMax = c
					idx = i
				}
			}
		}
		k++
		points[k], points[idx] = points[idx], points[k]
		if points[k].X == points[0].X && points[k].Y == points[0].Y {
			break
		}
	}
	return points[:k+1]
}

func cosFi(v1 Primitives.Point, v2 Primitives.Point) float64 {
	dot := (v1.X * v2.X + v1.Y * v2.Y) * 1.0
	d1 := math.Sqrt(v1.X * v1.X + v1.Y * v1.Y)
	d2 := math.Sqrt(v2.X * v2.X + v2.Y * v2.Y)
	return dot / (d1 * d2)
}
