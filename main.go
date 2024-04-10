package main

import (
	"fmt"
	"math"
	"math/rand"
)

type ColorScheme struct {
	FaceColor        string
	BackgroundColor  string
	HairColor        string
	MouthColor       string
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

func main() {
	colors := generateHarmoniousColors()
	fmt.Println("BackgroundColor:", colors.BackgroundColor)
}
