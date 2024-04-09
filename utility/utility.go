package utility

import (
    "math/rand"
)


// randomFromInterval generates a random float64 between min and max, inclusive
func randomFromInterval(min, max float64) float64 {
    return min + rand.Float64()*(max-min)
}