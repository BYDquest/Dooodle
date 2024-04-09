package avatar

import (
    "math"
    "github.com/BYDquest/RRRDoodlers/utility"
)



func generateMouthShape0(faceContour []point, faceHeight, faceWidth float64) []point {
    mouthRightY := utility.RandomFromInterval(faceHeight/7, faceHeight/3.5)
    mouthLeftY := utility.RandomFromInterval(faceHeight/7, faceHeight/3.5)
    mouthRightX := utility.RandomFromInterval(faceWidth/10, faceWidth/2)
    mouthLeftX := -mouthRightX + utility.RandomFromInterval(-faceWidth/20, faceWidth/20)
    mouthRight := point{mouthRightX, mouthRightY}
    mouthLeft := point{mouthLeftX, mouthLeftY}

    controlPoint0 := point{utility.RandomFromInterval(0, mouthRightX), utility.RandomFromInterval(mouthLeftY+5, faceHeight/1.5)}
    controlPoint1 := point{utility.RandomFromInterval(mouthLeftX, 0), utility.RandomFromInterval(mouthLeftY+5, faceHeight/1.5)}

    mouthPoints := make([]point, 0)
    for i := 0.0; i < 1; i += 0.01 {
        mouthPoints = append(mouthPoints, utility.CubicBezier(mouthLeft, controlPoint1, controlPoint0, mouthRight, i))
    }

    // Additional logic omitted for brevity - adapt from your JavaScript as needed
    return mouthPoints
}

func generateMouthShape1(faceHeight, faceWidth float64) []point {
    mouthRightY := utility.RandomFromInterval(faceHeight/7, faceHeight/4)
    mouthLeftY := utility.RandomFromInterval(faceHeight/7, faceHeight/4)
    mouthRightX := utility.RandomFromInterval(faceWidth/10, faceWidth/2)
    mouthLeftX := -mouthRightX + utility.RandomFromInterval(-faceWidth/20, faceWidth/20)
    mouthRight := point{mouthRightX, mouthRightY}
    mouthLeft := point{mouthLeftX, mouthLeftY}

    controlPoint0 := point{utility.RandomFromInterval(0, mouthRightX), utility.RandomFromInterval(mouthLeftY+5, faceHeight/1.5)}
    controlPoint1 := point{utility.RandomFromInterval(mouthLeftX, 0), utility.RandomFromInterval(mouthLeftY+5, faceHeight/1.5)}

    mouthPoints := make([]point, 0)
    for i := 0.0; i < 1; i += 0.01 {
        mouthPoints = append(mouthPoints, utility.CubicBezier(mouthLeft, controlPoint1, controlPoint0, mouthRight, i))
    }

    // Center, rotate, scale, and adjust the mouth shape
    center := point{(mouthRight.X + mouthLeft.X) / 2, (mouthPoints[25].Y + mouthPoints[75].Y) / 2}
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


func generateMouthShape2(faceContour []point, faceHeight, faceWidth float64) []point {

    
    center := point{utility.RandomFromInterval(-faceWidth/8, faceWidth/8), utility.RandomFromInterval(faceHeight/4, faceHeight/2.5)}

    mouthPoints := getEggShapePoints(utility.RandomFromInterval(faceWidth/4, faceWidth/10), utility.RandomFromInterval(faceHeight/10, faceHeight/20), 0.001, 50)
    randomRotationDegree := utility.RandomFromInterval(-math.Pi/9.5, math.Pi/9.5)

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
