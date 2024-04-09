package doodle

import (
	"math"
	"math/rand"
)

type Point struct {
	X float64
	Y float64
}

// randomFromInterval generates a random float64 between min and max, inclusive
func randomFromInterval(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func cubicBezier(P0, P1, P2, P3 Point, t float64) Point {
	x := math.Pow(1-t, 3)*P0.X + 3*math.Pow(1-t, 2)*t*P1.X + 3*(1-t)*math.Pow(t, 2)*P2.X + math.Pow(t, 3)*P3.X
	y := math.Pow(1-t, 3)*P0.Y + 3*math.Pow(1-t, 2)*t*P1.Y + 3*(1-t)*math.Pow(t, 2)*P2.Y + math.Pow(t, 3)*P3.Y
	return Point{x, y}
}

func GetEggShapePoints(a, b, k float64, segmentPoints int) []Point {
	result := make([]Point, 0, segmentPoints*4)
	for i := 0; i < segmentPoints; i++ {
		degree := (math.Pi/2/float64(segmentPoints))*float64(i) + randomFromInterval(-math.Pi/1.1/float64(segmentPoints), math.Pi/1.1/float64(segmentPoints))
		y := math.Sin(degree) * b
		x := math.Sqrt(((1 - (y*y)/(b*b)) / (1 + k*y)) * a * a)
		result = append(result, Point{x, y}, Point{-x, y}, Point{-x, -y}, Point{x, -y})
	}
	return result
}
