package utility

import (
	"math"
    "math/rand"
)

type Point struct {
    X float64
    Y float64
}

// randomFromInterval generates a random float64 between min and max, inclusive
func RandomFromInterval(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}

func cubicBezier(P0, P1, P2, P3 Point, t float64) Point {
    x := math.Pow(1-t, 3)*P0.X + 3*math.Pow(1-t, 2)*t*P1.X + 3*(1-t)*math.Pow(t, 2)*P2.X + math.Pow(t, 3)*P3.X
    y := math.Pow(1-t, 3)*P0.Y + 3*math.Pow(1-t, 2)*t*P1.Y + 3*(1-t)*math.Pow(t, 2)*P2.Y + math.Pow(t, 3)*P3.Y
    return Point{x, y}
}