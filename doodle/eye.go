package doodle

import (
	"math"
	"math/rand"
)

type EyeParameters struct {
	// Eyelid heights
	HeightUpper float64
	HeightLower float64

	// Random adjustments for initial points of the upper eyelid
	P0UpperRandX float64
	P3UpperRandX float64
	P0UpperRandY float64
	P3UpperRandY float64

	// Random adjustments for control points' vertical positions
	OffsetUpperLeftRandY  float64
	OffsetUpperRightRandY float64

	// Calculated width of the eye, considering random adjustments
	EyeTrueWidth float64

	// Horizontal offsets for control points to create the curvature of the eyelids
	OffsetUpperLeftX  float64
	OffsetUpperRightX float64
	OffsetLowerLeftX  float64
	OffsetLowerRightX float64

	// Vertical offsets for control points to shape the eyelids
	OffsetUpperLeftY  float64
	OffsetUpperRightY float64
	OffsetLowerLeftY  float64
	OffsetLowerRightY float64

	// Parameters to control the convergence of the bezier control points
	LeftConverge0  float64
	RightConverge0 float64
	LeftConverge1  float64
	RightConverge1 float64
}

type EyePoints struct {
	Upper, Lower, Center [][]float64
}


func cubicBezier(P0, P1, P2, P3 [2]float64, t float64) [2]float64 {
	x := math.Pow(1-t, 3)*P0[0] + 3*math.Pow(1-t, 2)*t*P1[0] + 3*(1-t)*math.Pow(t, 2)*P2[0] + math.Pow(t, 3)*P3[0]
	y := math.Pow(1-t, 3)*P0[1] + 3*math.Pow(1-t, 2)*t*P1[1] + 3*(1-t)*math.Pow(t, 2)*P2[1] + math.Pow(t, 3)*P3[1]
	return [2]float64{x, y}
}

func generateEyeParameters(width float64) EyeParameters {
	heightUpper := rand.Float64() * width / 1.2
	heightLower := rand.Float64() * width / 1.2
	P0UpperRandX := rand.Float64()*0.4 - 0.2
	P3UpperRandX := rand.Float64()*0.4 - 0.2
	P0UpperRandY := rand.Float64()*0.4 - 0.2
	P3UpperRandY := rand.Float64()*0.4 - 0.2
	offsetUpperLeftRandY := rand.Float64()
	offsetUpperRightRandY := rand.Float64()

	eyeTrueWidth := width + (P3UpperRandX-P0UpperRandX)*width/16

	offsetUpperLeftX := randomInRange(-eyeTrueWidth/10.0, eyeTrueWidth/2.3)
	offsetUpperRightX := randomInRange(-eyeTrueWidth/10.0, eyeTrueWidth/2.3)
	offsetUpperLeftY := offsetUpperLeftRandY * heightUpper
	offsetUpperRightY := offsetUpperRightRandY * heightUpper
	offsetLowerLeftX := randomInRange(offsetUpperLeftX, eyeTrueWidth/2.1)
	offsetLowerRightX := randomInRange(offsetUpperRightX, eyeTrueWidth/2.1)
	offsetLowerLeftY := randomInRange(-offsetUpperLeftY+5, heightLower)
	offsetLowerRightY := randomInRange(-offsetUpperRightY+5, heightLower)

	return EyeParameters{
		HeightUpper: heightUpper, HeightLower: heightLower,
		P0UpperRandX: P0UpperRandX, P3UpperRandX: P3UpperRandX, P0UpperRandY: P0UpperRandY, P3UpperRandY: P3UpperRandY,
		OffsetUpperLeftRandY: offsetUpperLeftRandY, OffsetUpperRightRandY: offsetUpperRightRandY,
		EyeTrueWidth:     eyeTrueWidth,
		OffsetUpperLeftX: offsetUpperLeftX, OffsetUpperRightX: offsetUpperRightX,
		OffsetUpperLeftY: offsetUpperLeftY, OffsetUpperRightY: offsetUpperRightY,
		OffsetLowerLeftX: offsetLowerLeftX, OffsetLowerRightX: offsetLowerRightX,
		OffsetLowerLeftY: offsetLowerLeftY, OffsetLowerRightY: offsetLowerRightY,
		LeftConverge0: rand.Float64(), RightConverge0: rand.Float64(),
		LeftConverge1: rand.Float64(), RightConverge1: rand.Float64(),
	}
}

func generateBezierCurve(P0, P1, P2, P3 [2]float64, convergence [2]float64, left bool, weightFunction func(int) float64) [][]float64 {
	points := make([][]float64, 100)
	for t := 0; t < 100; t++ {
		points[t] = make([]float64, 2)
		bt := cubicBezier(P0, P1, P2, P3, float64(t)/99)
		points[t][0] = bt[0]
		points[t][1] = bt[1]
	}

	controlPoints := make([][]float64, 100)

	for t := 0; t < 100; t++ {
		var point [2]float64
		if left {
			point = cubicBezier([2]float64{convergence[0], convergence[1]}, P0, P1, P2, float64(t)/99)
		} else {
			point = cubicBezier(P1, P2, P3, [2]float64{convergence[0], convergence[1]}, float64(t)/99)
		}
		// Convert [2]float64 array to a slice []float64 and assign it to controlPoints[t]
		controlPoints[t] = []float64{point[0], point[1]}
	}

	for i := 25; i < 100; i++ {
		weight := weightFunction(i)
		if left && i < 75 {
			points[i][0] = points[i][0]*(1-weight) + controlPoints[i][0]*weight
			points[i][1] = points[i][1]*(1-weight) + controlPoints[i][1]*weight
		} else if !left && i >= 25 {
			j := i - 25
			if j >= 75 {
				break
			}
			points[j][0] = points[j][0]*weight + controlPoints[i][0]*(1-weight)
			points[j][1] = points[j][1]*weight + controlPoints[i][1]*(1-weight)
		}
	}

	return points
}

func generateEyePoints(rands EyeParameters, width float64) EyePoints {
	P0Upper := [2]float64{-width/2 + rands.P0UpperRandX*width/16, rands.P0UpperRandY * rands.HeightUpper / 16}
	P3Upper := [2]float64{width/2 + rands.P3UpperRandX*width/16, rands.P3UpperRandY * rands.HeightUpper / 16}
	P1Upper := [2]float64{P0Upper[0] + rands.OffsetUpperLeftX, P0Upper[1] + rands.OffsetUpperLeftY}
	P2Upper := [2]float64{P3Upper[0] - rands.OffsetUpperRightX, P3Upper[1] + rands.OffsetUpperRightY}

	P0Lower := P0Upper
	P3Lower := P3Upper
	P1Lower := [2]float64{P0Lower[0] + rands.OffsetLowerLeftX, P0Lower[1] - rands.OffsetLowerLeftY}
	P2Lower := [2]float64{P3Lower[0] - rands.OffsetLowerRightX, P3Lower[1] - rands.OffsetLowerRightY}

	upperLeftControlPoint := [2]float64{P0Upper[0]*(1-rands.LeftConverge0) + P1Lower[0]*rands.LeftConverge0, P0Upper[1]*(1-rands.LeftConverge0) + P1Lower[1]*rands.LeftConverge0}
	upperRightControlPoint := [2]float64{P3Upper[0]*(1-rands.RightConverge0) + P2Lower[0]*rands.RightConverge0, P3Upper[1]*(1-rands.RightConverge0) + P2Lower[1]*rands.RightConverge0}
	lowerLeftControlPoint := [2]float64{P0Lower[0]*(1-rands.LeftConverge1) + P1Upper[0]*rands.LeftConverge1, P0Lower[1]*(1-rands.LeftConverge1) + P1Upper[1]*rands.LeftConverge1}
	lowerRightControlPoint := [2]float64{P3Lower[0]*(1-rands.RightConverge1) + P2Upper[0]*rands.RightConverge1, P3Lower[1]*(1-rands.RightConverge1) + P2Upper[1]*rands.RightConverge1}

	upperEyelidPoints := generateBezierCurve(P0Upper, P1Upper, P2Upper, P3Upper, upperLeftControlPoint, true, func(i int) float64 {
		return math.Pow(float64(75-i)/75, 2)
	})
	upperEyelidPoints = append(upperEyelidPoints, generateBezierCurve(P0Upper, P1Upper, P2Upper, P3Upper, upperRightControlPoint, false, func(i int) float64 {
		return math.Pow(float64(i-25)/75, 2)
	})...)

	lowerEyelidPoints := generateBezierCurve(P0Lower, P1Lower, P2Lower, P3Lower, lowerLeftControlPoint, true, func(i int) float64 {
		return math.Pow(float64(75-i)/75, 2)
	})
	lowerEyelidPoints = append(lowerEyelidPoints, generateBezierCurve(P0Lower, P1Lower, P2Lower, P3Lower, lowerRightControlPoint, false, func(i int) float64 {
		return math.Pow(float64(i-25)/75, 2)
	})...)

	eyeCenter := [2]float64{(upperEyelidPoints[50][0] + lowerEyelidPoints[50][0]) / 2, (upperEyelidPoints[50][1] + lowerEyelidPoints[50][1]) / 2}

	for i := range upperEyelidPoints {
		upperEyelidPoints[i][0] -= eyeCenter[0]
		upperEyelidPoints[i][1] -= eyeCenter[1]
		lowerEyelidPoints[i][0] -= eyeCenter[0]
		lowerEyelidPoints[i][1] -= eyeCenter[1]
	}

	return EyePoints{Upper: upperEyelidPoints, Lower: lowerEyelidPoints, Center: [][]float64{{0, 0}}}
}

func GenerateBothEyes(width float64) (EyePoints, EyePoints) {
	randsLeft := generateEyeParameters(width)
	randsRight := randsLeft

	randsRight.HeightUpper += randomInRange(-randsRight.HeightUpper/2, randsRight.HeightUpper/2)
	// Repeat the above line for other fields of randsRight as needed

	leftEye := generateEyePoints(randsLeft, width)
	rightEye := generateEyePoints(randsRight, width)

	return leftEye, rightEye
}
