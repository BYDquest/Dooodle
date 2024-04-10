package main

import (
	"fmt"
	"math/rand"
	"math"
)

func generateHarmoniousColors() map[string]string {
	baseHue := rand.Float64() // Random base hue
	saturation := 0.75
	lightnessForFace := 0.5
	var lightnessForBackground, lightnessForHair, lightnessForMouth float64

	if lightnessForFace > 0.5 {
		lightnessForBackground = 0.2
	} else {
		lightnessForBackground = 0.8
	}

	if lightnessForFace < 0.5 {
		lightnessForHair = 0.7
	} else {
		lightnessForHair = 0.3
	}

	if lightnessForFace > 0.5 {
		lightnessForMouth = 0.4
	} else {
		lightnessForMouth = 0.6
	}

	complementaryHue := math.Mod(baseHue+0.5, 1)
	triadicHue1 := math.Mod(baseHue+1/3, 1)

	r, g, b := hslToRgb(baseHue, saturation, lightnessForFace)
	faceColor := rgbToString(r, g, b)

	r, g, b = hslToRgb(complementaryHue, saturation, lightnessForBackground)
	backgroundColor := rgbToString(r, g, b)

	r, g, b = hslToRgb(triadicHue1, saturation, lightnessForHair)
	hairColor := rgbToString(r, g, b)

	r, g, b = hslToRgb(baseHue, saturation-0.25, lightnessForMouth)
	mouthColor := rgbToString(r, g, b)

	colors := map[string]string{
		"faceColor":        faceColor,
		"backgroundColor":  backgroundColor,
		"hairColor":        hairColor,
		"mouthColor":       mouthColor,
	}

	return colors
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
			if t < 0 {
				t += 1
			}
			if t > 1 {
				t -= 1
			}
			if t < 1/6 {
				return p + (q-p)*6*t
			}
			if t < 1/2 {
				return q
			}
			if t < 2/3 {
				return p + (q-p)*(2/3-t)*6
			}
			return p
		}

		q := l * (1 + s)
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hue2rgb(p, q, h+1/3)
		g = hue2rgb(p, q, h)
		b = hue2rgb(p, q, h-1/3)
	}

	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}


func main(){

	va:=generateHarmoniousColors()

	fmt.Println(va["backgroundColor"])
}