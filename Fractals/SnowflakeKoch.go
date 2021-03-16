package Fractals

import (
	"Graphics/Primitives"
	"math"
)

func SnowflakeKoch(points []Primitives.Point, k int) []Primitives.Point {
	side1 := kochCurve(points[0], points[1], k)
	side2 := kochCurve(points[1], points[2], k)
	side3 := kochCurve(points[2], points[0], k)
	return append(append(side1, side2...), side3...)
}

func kochCurve(startPoint Primitives.Point, endPoint Primitives.Point, k int) []Primitives.Point {
	curvePoints := []Primitives.Point{startPoint, endPoint}
	for i := 1; i <= k; i++ {
		var tmpPoints []Primitives.Point
		for j := 0; j < len(curvePoints) - 1; j++ {
			innerPoints := kochSegment(curvePoints[j], curvePoints[j+1])
			tmpPoints = append(tmpPoints, curvePoints[j])
			tmpPoints = append(tmpPoints, innerPoints...)
		}
		tmpPoints = append(tmpPoints, curvePoints[len(curvePoints)-1])
		curvePoints = tmpPoints
	}
	return curvePoints
}

func kochSegment(startPoint Primitives.Point, endPoint Primitives.Point) []Primitives.Point {
	v := Primitives.Point{
		X: (endPoint.X - startPoint.X) / 3,
		Y: (endPoint.Y - startPoint.Y) / 3,
	}
	p1 := Primitives.Point{
		X: startPoint.X + v.X,
		Y: startPoint.Y + v.Y,
	}
	p3 := Primitives.Point{
		X: endPoint.X - v.X,
		Y: endPoint.Y - v.Y,
	}
	angle := -60.0
	v = rotate(v, angle * math.Pi / 180)
	p2 := Primitives.Point{
		X: p1.X + v.X,
		Y: p1.Y + v.Y,
	}
	return []Primitives.Point{p1, p2, p3}
}

func rotate(point Primitives.Point, angle float64) Primitives.Point {
	return Primitives.Point{
		X: point.X * math.Cos(angle) - point.Y * math.Sin(angle),
		Y: point.X * math.Sin(angle) + point.Y * math.Cos(angle),
	}
}