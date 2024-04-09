package avatar

import (
    "math"
    "github.com/BYDquest/RRRDoodlers/utility"
)


func getEggShapePoints(a, b, k float64, segmentPoints int) []Point {
    result := make([]Point, 0, segmentPoints*4)
    for i := 0; i < segmentPoints; i++ {
        degree := (math.Pi/2/float64(segmentPoints))*float64(i) + utility.randomFromInterval(-math.Pi/1.1/float64(segmentPoints), math.Pi/1.1/float64(segmentPoints))
        y := math.Sin(degree) * b
        x := math.Sqrt(((1 - (y*y)/(b*b)) / (1 + k*y)) * a * a)
        result = append(result, Point{x, y}, Point{-x, y}, Point{-x, -y}, Point{x, -y})
    }
    return result
}

func generateMouthShape0(faceContour []Point, faceHeight, faceWidth float64) []Point {
    mouthRightY := utility.randomFromInterval(faceHeight/7, faceHeight/3.5)
    mouthLeftY := utility.randomFromInterval(faceHeight/7, faceHeight/3.5)
    mouthRightX := utility.randomFromInterval(faceWidth/10, faceWidth/2)
    mouthLeftX := -mouthRightX + utility.randomFromInterval(-faceWidth/20, faceWidth/20)
    mouthRight := Point{mouthRightX, mouthRightY}
    mouthLeft := Point{mouthLeftX, mouthLeftY}

    controlPoint0 := Point{utility.randomFromInterval(0, mouthRightX), utility.randomFromInterval(mouthLeftY+5, faceHeight/1.5)}
    controlPoint1 := Point{utility.randomFromInterval(mouthLeftX, 0), utility.randomFromInterval(mouthLeftY+5, faceHeight/1.5)}

    mouthPoints := make([]Point, 0)
    for i := 0.0; i < 1; i += 0.01 {
        mouthPoints = append(mouthPoints, cubicBezier(mouthLeft, controlPoint1, controlPoint0, mouthRight, i))
    }

    // Additional logic omitted for brevity - adapt from your JavaScript as needed
    return mouthPoints
}

func generateMouthShape1(faceHeight, faceWidth float64) []Point {
    mouthRightY := utility.randomFromInterval(faceHeight/7, faceHeight/4)
    mouthLeftY := utility.randomFromInterval(faceHeight/7, faceHeight/4)
    mouthRightX := utility.randomFromInterval(faceWidth/10, faceWidth/2)
    mouthLeftX := -mouthRightX + utility.randomFromInterval(-faceWidth/20, faceWidth/20)
    mouthRight := Point{mouthRightX, mouthRightY}
    mouthLeft := Point{mouthLeftX, mouthLeftY}

    controlPoint0 := Point{utility.randomFromInterval(0, mouthRightX), utility.randomFromInterval(mouthLeftY+5, faceHeight/1.5)}
    controlPoint1 := Point{utility.randomFromInterval(mouthLeftX, 0), utility.randomFromInterval(mouthLeftY+5, faceHeight/1.5)}

    mouthPoints := make([]Point, 0)
    for i := 0.0; i < 1; i += 0.01 {
        mouthPoints = append(mouthPoints, cubicBezier(mouthLeft, controlPoint1, controlPoint0, mouthRight, i))
    }

    // Center, rotate, scale, and adjust the mouth shape
    center := Point{(mouthRight.X + mouthLeft.X) / 2, (mouthPoints[25].Y + mouthPoints[75].Y) / 2}
    for i := range mouthPoints {
        // Translate to center
        mouthPoints[i].X -= center.X
        mouthPoints[i].Y -= center.Y
        // Rotate 180 degrees (optional, based on your logic)
        mouthPoints[i].Y = -mouthPoints[i].Y
        // Scale smaller
        mouthPoints[i].X *= 0.6
        mouthPoints[i].Y *= 0.6
        // Translate back
        mouthPoints[i].X += center.X
        mouthPoints[i].Y += center.Y * 0.8
    }

    return mouthPoints
}


func generateMouthShape2(faceContour []Point, faceHeight, faceWidth float64) []Point {
    center := Point{utility.randomFromInterval(-faceWidth/8, faceWidth/8), utility.randomFromInterval(faceHeight/4, faceHeight/2.5)}

    mouthPoints := getEggShapePoints(utility.randomFromInterval(faceWidth/4, faceWidth/10), utility.randomFromInterval(faceHeight/10, faceHeight/20), 0.001, 50)
    randomRotationDegree := utility.randomFromInterval(-math.Pi/9.5, math.Pi/9.5)

    for i := range mouthPoints {
        x, y := mouthPoints[i].X, mouthPoints[i].Y
        mouthPoints[i].X = x*math.Cos(randomRotationDegree) - y*math.Sin(randomRotationDegree) + center.X
        mouthPoints[i].Y = x*math.Sin(randomRotationDegree) + y*math.Cos(randomRotationDegree) + center.Y
    }

    return mouthPoints
}

// func main() {
//     // Example usage
//     faceContour := []Point{} // Define your face contour points
//     faceHeight, faceWidth := 200.0, 100.0 // Example face dimensions

//     mouthShape := generateMouthShape0(faceContour, faceHeight, faceWidth)
//     for _, point := range mouthShape {
//         println(point.X, point.Y)
//     }

//     // Call other functions similarly
// }
