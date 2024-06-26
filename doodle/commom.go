package doodle

import (
	"fmt"

	"math"
	"math/rand"
	"os"
)

func ensureDirectoryExists(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
}

// to check
func convertArrayToFixedFloat(arr [][2]float64) [][2]float64 {
	var converted [][2]float64
	for _, pair := range arr {
		x := fmt.Sprintf("%.2f", pair[0])
		y := fmt.Sprintf("%.2f", pair[1])
		converted = append(converted, [2]float64{parseStringToFloat(x), parseStringToFloat(y)})
	}
	return converted
}

func parseStringToFloat(s string) float64 {
	var f float64
	fmt.Sscanf(s, "%f", &f)
	return f
}

func randomRGB() string {
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

func generateHarmoniousColors() ColorScheme {
	baseHue := rand.Float64() // Random base hue
	const saturation = 0.75
	const lightnessForFace = 0.5

	lightnessForBackground := 0.2
	if lightnessForFace > 0.5 {
		lightnessForBackground = 0.8
	}

	lightnessForHair := 0.3
	if lightnessForFace < 0.5 {
		lightnessForHair = 0.7
	}

	lightnessForMouth := 0.6
	if lightnessForFace > 0.5 {
		lightnessForMouth = 0.4
	}

	complementaryHue := math.Mod(baseHue+0.5, 1)
	triadicHue1 := math.Mod(baseHue+1/3.0, 1)

	faceColor := rgbToString(hslToRgb(baseHue, saturation, lightnessForFace))
	backgroundColor := rgbToString(hslToRgb(complementaryHue, saturation, lightnessForBackground))
	hairColor := rgbToString(hslToRgb(triadicHue1, saturation, lightnessForHair))
	mouthColor := rgbToString(hslToRgb(baseHue, saturation-0.25, lightnessForMouth))

	return ColorScheme{faceColor, backgroundColor, hairColor, mouthColor}
}

func rgbToString(r, g, b int) string {
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

func hslToRgb(h, s, l float64) (int, int, int) {
	var r, g, b float64
	if s == 0 {
		r, g, b = l, l, l // achromatic
	} else {
		var hue2rgb = func(p, q, t float64) float64 {
			switch {
			case t < 0:
				t += 1
			case t > 1:
				t -= 1
			}
			switch {
			case t < 1/6.0:
				return p + (q-p)*6*t
			case t < 1/2.0:
				return q
			case t < 2/3.0:
				return p + (q-p)*(2/3.0-t)*6
			}
			return p
		}

		q := l + s*math.Min(l, 1-l)
		p := 2*l - q
		r = hue2rgb(p, q, h+1/3.0)
		g = hue2rgb(p, q, h)
		b = hue2rgb(p, q, h-1/3.0)
	}

	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

func randomInRange(min, max float64) float64 {
	// Generate random number in the range [0, 1)
	random := rand.Float64()
	// Adjust random number to our range [min, max]
	return min + random*(max-min)
}

func GetEggShapePoints(a, b, k float64, segmentPoints int) []Point {
	result := make([]Point, 0, segmentPoints*4)
	for i := 0; i < segmentPoints; i++ {
		degree := (math.Pi/2/float64(segmentPoints))*float64(i) + randomInRange(-math.Pi/1.1/float64(segmentPoints), math.Pi/1.1/float64(segmentPoints))
		y := math.Sin(degree) * b
		x := math.Sqrt(((1 - (y*y)/(b*b)) / (1 + k*y)) * a * a)
		result = append(result, Point{x, y}, Point{-x, y}, Point{-x, -y}, Point{x, -y})
	}
	return result
}
