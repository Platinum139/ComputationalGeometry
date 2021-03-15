package ConvexHull

import (
	"Graphics/Primitives"
	"math"
	"sort"
)

func GrahamScan(points []Primitives.Point) []Primitives.Point {
	// find extreme point
	ymin := points[0].Y
	for i := range points {
		if points[i].Y < ymin {
			ymin = points[i].Y
			p := points[0]
			points[0] = points[i]
			points[i] = p
		}
	}
	p0 := points[0]
	points = points[1:]

	// sort by angle
	sort.SliceStable(points, func(i, j int) bool {
		v1 := Primitives.Point{
			X: points[i].X - p0.X,
			Y: points[i].Y - p0.Y,
		}
		v2 := Primitives.Point{
			X: points[j].X - p0.X,
			Y: points[j].Y - p0.Y,
		}
		a1 := angle(v1, Primitives.Point{X: points[0].X, Y: 0})
		a2 := angle(v2, Primitives.Point{X: points[0].X, Y: 0})
		if a1 < a2 {
			return true
		}
		if a1 == a2 {
			d1 := distance(points[0], points[i])
			d2 := distance(points[0], points[j])
			if d1 < d2 {
				return true
			}
		}
		return false
	})
	points = append([]Primitives.Point{p0}, points...)

	// create stack for convex hull
	var stack []Primitives.Point
	stack = append(stack, points[0])
	stack = append(stack, points[1])

	for i := 2; i < len(points); i++ {
		for {
			if ccw(stack[len(stack)-2], stack[len(stack)-1], points[i]) {
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, points[i])
	}
	return append(stack, p0)
}

func ccw(a, b, c Primitives.Point) bool {
	return ((b.X - a.X) * (c.Y - b.Y)) - ((b.Y - a.Y) * (c.X - b.X)) >= 0
}

func angle(v1 Primitives.Point, v2 Primitives.Point) float64 {
	dot := (v1.X * v2.X + v1.Y * v2.Y) * 1.0
	d1 := math.Sqrt(v1.X * v1.X + v1.Y * v1.Y)
	d2 := math.Sqrt(v2.X * v2.X + v2.Y * v2.Y)
	return math.Acos(dot / (d1 * d2))
}

func distance(p1 Primitives.Point, p2 Primitives.Point) float64 {
	a := p1.X - p2.X
	b := p1.Y - p2.Y
	return math.Sqrt(a * a + b * b)
}