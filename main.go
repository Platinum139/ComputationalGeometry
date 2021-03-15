package main

import (
	"Graphics/ConvexHull"
	"Graphics/InterpolationCurves"
	"Graphics/OrientationTests"
	"Graphics/Primitives"
	"github.com/fogleman/gg"
	"image/color"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)

	// white background
	dc.SetColor(color.White)
	dc.Clear()

	// drawBezierCurve(dc)
	// drawRationalBezierCurve(dc)
	// drawComplexBezierCurve(dc)
	// drawCubicBSpline(dc)
	// drawComplexBSpline(dc)
	// drawBSplineBasisFunctions(dc)
	// drawBSpline(dc)
	// drawCubicHermiteCurve(dc)
	// drawTCBSpline(dc)
	// drawGrahamScan(dc)
	// drawGiftWrapping(dc)
	// drawAngleTest(dc)

	dc.SavePNG("storage/orientationTests/angleTest4.png")
}

func drawBezierCurve(dc *gg.Context) {
	// array of points
	points := []Primitives.Point{
		{100, 200},
		{800, 200},
		{800, 700},
		{400, 500},
	}
	drawPolygonalChain(points, dc)

	// draw bezier curve
	dc.SetRGB(200, 150, 20)
	for t := 0.0; t <= 1.0; t += 0.01 {
		bc := InterpolationCurves.BezierCurve(points, t)
		dc.LineTo(bc.X, bc.Y)
	}
	dc.Stroke()
}

func drawRationalBezierCurve(dc *gg.Context) {
	// array of points
	points := []Primitives.Point{
		{100, 200},
		{800, 200},
		{800, 700},
		{400, 500},
	}
	drawPolygonalChain(points, dc)

	// draw rational bezier curve
	dc.SetRGB(200, 20, 200)
	w := []float64{0.5, 0.5, 0.9, 0.8}
	for t := 0.0; t <= 1.0; t += 0.01 {
		bc := InterpolationCurves.RationalBezierCurve(points, w, t)
		dc.LineTo(bc.X, bc.Y)
	}
	dc.Stroke()
}

func drawComplexBezierCurve(dc *gg.Context)  {
	// array of points
	points := []Primitives.Point{
		{400, 100},
		{200, 200},
		{550, 500},
		{600, 500},
		{650, 500},
		{780, 400},
		{810, 400},
		{840, 400},
		{950, 600},
		{980, 600},
	}
	drawPolygonalChain(points, dc)

	colors := [][]float64{
		{0.9, 0.1, 0.1},
		{0.3, 0.8, 0.2},
		{0.3, 0.2, 0.7},
	}
	i := 0
	for k := 0; k < len(points)-1; k += 3 {
		for t := 0.0; t <= 1.0; t += 0.01 {
			bc := InterpolationCurves.BezierCurve(points[k:k+4], t)
			dc.LineTo(bc.X, bc.Y)
		}
		dc.SetRGB(colors[i][0], colors[i][1], colors[i][2]); i++
		dc.SetLineWidth(10)
		dc.Stroke()
	}
}

func drawCubicBSpline(dc *gg.Context) {
	// array of points
	points := []Primitives.Point{
		{100, 500},
		{200, 200},
		{600, 100},
		{800, 800},
	}
	drawPolygonalChain(points, dc)

	// draw b spline curve
	dc.SetRGB(200, 150, 20)
	for t := 0.0; t <= 1.0; t += 0.01 {
		bc := InterpolationCurves.CubicBSpline(points, t)
		dc.LineTo(bc.X, bc.Y)
	}
	dc.Stroke()
}

func drawComplexBSpline(dc *gg.Context) {
	// array of points
	points := []Primitives.Point{
		{100, 600},
		{200, 200},
		{500, 200},
		{600, 600},
		{750, 600},
		{900, 200},
	}
	drawPolygonalChain(points, dc)

	colors := [][]float64{
		{0.9, 0.1, 0.1},
		{0.3, 0.8, 0.2},
		{0.3, 0.2, 0.7},
	}
	for k := 0; k < len(points)-3; k++ {
		for t := 0.0; t <= 1.0; t += 0.01 {
			bc := InterpolationCurves.CubicBSpline(points[k:k+4], t)
			dc.LineTo(bc.X, bc.Y)
		}
		dc.SetRGB(colors[k][0], colors[k][1], colors[k][2])
		dc.SetLineWidth(10)
		dc.Stroke()
	}
}

func drawBSplineBasisFunctions(dc *gg.Context) {
	tVec := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	colors := [][]float64{{0.9, 0.3, 0.0}, {0.0, 0.9, 0.1}, {0.9, 0.9, 0.0}, {0.1, 0.7, 0.8} }
	d := 50.0
	xStart := 150.0
	yStart := 150.0
	drawFourCoordinateSystems(dc, d, xStart, yStart)

	for k := 0; k <= 3; k++ {
		for i := 0; i <= 3; i++ {
			for t := tVec[0]; t < tVec[len(tVec)-1]; t += 0.01 {
				N := InterpolationCurves.FuncN(int(i), k, t, tVec)
				displayX, displayY := transformCoordinatesToDisplay(t, N, xStart, yStart, d)
				l := len(colors)
				dc.SetRGB(colors[i % l][0], colors[i % l][1], colors[i % l][2])
				dc.DrawPoint(displayX, displayY, 3)
				dc.Fill()
			}
		}
		yStart += 250
	}
}

func drawBSpline(dc *gg.Context) {
	// array of points
	points := []Primitives.Point{
		{100, 600},
		{200, 200},
		{300, 600},
		{400, 200},
		{500, 600},
		{600, 200},
		{700, 600},
		{800, 200},
	}
	drawPolygonalChain(points, dc)

	m := len(points) - 1
	k := 3	// cubic B-spline
	tVec := []float64{0, 0.5, 1, 2, 4, 8, 8.5, 9, 9.5, 10, 11, 13}

	for i := k; i <= m; i++ {
		for t := tVec[i]; t <= tVec[i+1]; t += 0.01 {
			x := 0.0
			y := 0.0
			for j := i - k; j <= i; j++ {
				N := InterpolationCurves.FuncN(j, k, t, tVec)
				x += N * points[j].X
				y += N * points[j].Y
			}
			// draw x, y
			dc.SetRGB(1.0, 0.4, 0.2)
			dc.DrawPoint(x, y, 3)
			dc.Fill()
		}
	}
}

func drawCubicHermiteCurve(dc *gg.Context) {
	// array of points
	points := []Primitives.Point{
		{100, 200},
		{800, 200},
	}
	// array of vectors
	vectors := []Primitives.Point{
		{500, 500},
		{600, 500},
	}
	// draw points
	dc.SetColor(color.Black)
	dc.DrawPoint(points[0].X, points[0].Y, 10)
	dc.DrawPoint(points[1].X, points[1].Y, 10)
	dc.Fill()

	// draw vectors
	dc.SetRGB(0.4, 0.8, 0.0)
	dc.SetLineWidth(7)
	dc.DrawLine(points[0].X, points[0].Y, points[0].X + 0.2 * vectors[0].X, points[0].Y + 0.2 * vectors[0].Y)
	dc.DrawLine(points[1].X, points[1].Y, points[1].X - 0.2 * vectors[1].X, points[1].Y - 0.2 * vectors[1].Y)
	dc.Stroke()

	// draw hermite curve
	dc.SetRGB(200, 150, 20)
	dc.SetLineWidth(5)
	for t := 0.0; t <= 1.0; t += 0.01 {
		point := InterpolationCurves.CubicHermiteCurve(points, vectors, t)
		dc.LineTo(point.X, point.Y)
	}
	dc.Stroke()

	// array of points
	points = []Primitives.Point{
		{100, 600},
		{800, 600},
	}
	// array of vectors
	vectors = []Primitives.Point{
		{1000, 1000},
		{600, 500},
	}
	// draw points
	dc.SetColor(color.Black)
	dc.DrawPoint(points[0].X, points[0].Y, 10)
	dc.DrawPoint(points[1].X, points[1].Y, 10)
	dc.Fill()

	// draw vectors
	dc.SetRGB(0.4, 0.8, 0.0)
	dc.SetLineWidth(7)
	dc.DrawLine(points[0].X, points[0].Y, points[0].X + 0.2 * vectors[0].X, points[0].Y + 0.2 * vectors[0].Y)
	dc.DrawLine(points[1].X, points[1].Y, points[1].X - 0.2 * vectors[1].X, points[1].Y - 0.2 * vectors[1].Y)
	dc.Stroke()

	// draw hermite curve
	dc.SetRGB(200, 150, 20)
	dc.SetLineWidth(5)
	for t := 0.0; t <= 1.0; t += 0.01 {
		point := InterpolationCurves.CubicHermiteCurve(points, vectors, t)
		dc.LineTo(point.X, point.Y)
	}
	dc.Stroke()
}

func drawTCBSpline(dc *gg.Context) {
	// array of points
	points := []Primitives.Point{
		{100, 400},
		{300, 250},
		{300, 500},
		{400, 550},
		{500, 450},
		{600, 600},
		{700, 700},
		{450, 650},
	}
	drawPolygonalChain(points, dc)

	colors := [][]float64{
		{0.9, 0.3, 0.0}, {0.0, 0.9, 0.1}, {0.9, 0.9, 0.0},
		{0.1, 0.7, 0.8}, {0.9, 0.1, 0.9}}
	m := len(points) - 1
	t := 0.1
	c := 0.8
	b := 0.2

	for k := 1; k < m-1; k++ {
		p1 := points[k]
		p2 := points[k+1]
		q1 := InterpolationCurves.StartVec(points[k-1:k+2], t, c, b)
		q2 := InterpolationCurves.EndVec(points[k:k+3], t, c, b)

		for i := 0.0; i <= 1.0; i += 0.01 {
			p := InterpolationCurves.TCBSpline(p1, p2, q1, q2, i)
			dc.SetRGB(colors[k-1][0], colors[k-1][1], colors[k-1][2])
			dc.LineTo(p.X, p.Y)
		}
		dc.Stroke()
	}
}

func drawGrahamScan(dc *gg.Context) {
	points := []Primitives.Point{
		{100, 400},
		{800, 600},
		{450, 500},
		{300, 600},
		{350, 750},
		{250, 300},
		{150, 550},
		{250, 400},
	}
	// draw points
	dc.SetColor(color.Black)
	for i := range points {
		dc.DrawPoint(points[i].X, points[i].Y, 10)
		dc.Fill()
	}

	dc.SetRGB(1.0, 0.1, 0.7)
	dc.SetLineWidth(8)
	pointsHull := ConvexHull.GrahamScan(points)

	for i := range pointsHull {
		dc.LineTo(pointsHull[i].X, pointsHull[i].Y)
	}
	dc.Stroke()
}

func drawGiftWrapping(dc *gg.Context) {
	points := []Primitives.Point{
		{100, 400},
		{800, 600},
		{450, 500},
		{300, 600},
		{350, 750},
		{400, 300},
		{150, 550},
		{250, 400},
		{250, 300},
	}
	// draw points
	dc.SetColor(color.Black)
	for i := range points {
		dc.DrawPoint(points[i].X, points[i].Y, 10)
		dc.Fill()
	}

	dc.SetRGB(1.0, 0.7, 0.2)
	dc.SetLineWidth(8)
	pointsHull := ConvexHull.GiftWrapping(points)

	for i := range pointsHull {
		dc.LineTo(pointsHull[i].X, pointsHull[i].Y)
	}
	dc.Stroke()
}

func drawAngleTest(dc *gg.Context) {
	points := []Primitives.Point{
		{100, 400},
		{250, 200},
		{600, 300},
		{350, 600},
		{200, 580},
		{100, 400},
	}
	point := Primitives.Point{X: 500, Y: 420}

	// draw polyline and point
	drawPolygonalChain(points, dc)

	r := OrientationTests.AngleTest(points[:len(points)-1], point)
	switch r {
	// out
	case 0:
		dc.SetRGB(1.0, 0.1, 0.1)
	// in
	case 1:
		dc.SetRGB(0.0, 0.9, 0.1)
	// top
	case 2:
		dc.SetRGB(0.1, 0.5, 1.0)
	// on
	case 3:
		dc.SetRGB(1.0, 0.9, 0.0)
	}
	dc.DrawPoint(point.X, point.Y, 10)
	dc.Fill()
}

func transformCoordinatesToDisplay(x float64, y float64, xs float64, ys float64, d float64) (float64, float64) {
	factor := d
	displayX := xs + x * factor
	displayY := ys - y * factor
	return displayX, displayY
}

func drawFourCoordinateSystems(dc *gg.Context, d float64, xStart float64, yStart float64) {
	dc.SetColor(color.Black)
	dc.SetLineWidth(5)
	for k := 1; k <= 4; k++ {
		dc.DrawLine(xStart-d, yStart, xStart + 550, yStart)
		dc.DrawLine(xStart, yStart-100, xStart, yStart+50)
		for i := 1.0; i <= 10; i++ {
			dc.DrawLine(xStart + i*d, yStart-10, xStart + i*d, yStart+10)
		}
		dc.DrawLine(xStart-10, yStart-d, xStart+10, yStart-d)
		dc.Stroke()
		yStart += 250
	}
}

func drawPolygonalChain(points []Primitives.Point, dc *gg.Context) {
	// join points
	dc.SetRGB(100, 100, 100)
	dc.SetLineWidth(5)
	for i := 0; i < len(points); i++ {
		dc.LineTo(points[i].X, points[i].Y)
	}
	dc.Stroke()

	// draw points
	dc.SetColor(color.Black)
	for i := range points {
		dc.DrawPoint(points[i].X, points[i].Y, 10)
		dc.Fill()
	}
}




