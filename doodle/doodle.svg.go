package doodle

import (
	"fmt"
	"strconv"
	"strings"
)

func doodleSvg() string {
	// Assuming these variables are defined somewhere in your Go code.
	var (
		computedFacePoints  []string
		eyeLeftContour      []string // This should be of type []string, not []string.
		eyeRightContour     []string
		eyeRightUpper       []string
		eyeRightLower       []string
		rightEyePupilMarkup string // Assuming this is a string containing valid SVG markup.
		eyeLeftUpper        []string
		eyeLeftLower        []string
		leftEyePupilMarkup  string
		hairs               [][]string // This would be a slice of slices of string for each hair polyline points.
		noseMarkup          string
		mouthPoints         []string
		backgroundColor     string
		faceColor           string
		faceScale           float64
		center              [2]float64
		distanceBetweenEyes float64
		rightEyeOffsetX     float64
		rightEyeOffsetY     float64
		leftEyeOffsetX      float64
		leftEyeOffsetY      float64
		eyeHeightOffset     float64
		hairColor           string
		mouthColor          string
		dyeColorOffset      string // Assuming this is a string.
	)

	// Helper function to join slices of strings with a space, mimicking JavaScript's Array.join(' ').
	joinWithSpace := func(slice []string) string {
		return strings.Join(slice, " ")
	}

	// RandomRGB and other dynamic functions need to be defined in Go.
	randomRGB := func() string {
		// Implement the randomRGB function or replace it with your implementation.
		return "#000000" // Placeholder return.
	}

	// For the hairs part, you'll need to iterate over the hairs slice and construct each part.
	hairsMarkup := ""
	for _, hair := range hairs {
		hairsMarkup += fmt.Sprintf(`<polyline points="%s" fill="none" stroke="%s" stroke-width="2" stroke-linejoin="round" filter="url(#fuzzy)" />`, joinWithSpace(hair), hairColor)
	}
	svgString := `<svg viewBox="-100 -100 200 200" xmlns="http://www.w3.org/2000/svg" width="500" height="500" id="face-svg">
    <defs>
      <clipPath id="leftEyeClipPath">
        <polyline points="` + strings.Join(eyeLeftContour, " ") + `" />
      </clipPath>
      <clipPath id="rightEyeClipPath">
        <polyline points="` + strings.Join(eyeRightContour, " ") + `" />
      </clipPath>
      <filter id="fuzzy">
        <feTurbulence id="turbulence" baseFrequency="0.05" numOctaves="3" type="noise" result="noise" />
        <feDisplacementMap in="SourceGraphic" in2="noise" scale="2" />
      </filter>
      <linearGradient id="rainbowGradient" x1="0%" y1="0%" x2="100%" y2="0%">
        <stop offset="0%" style="stop-color: ` + randomRGB() + `; stop-opacity: 1" />
        <stop offset="` + dyeColorOffset + `" style="stop-color: ` + randomRGB() + `; stop-opacity: 1" />
        <stop offset="100%" style="stop-color: ` + randomRGB() + `; stop-opacity: 1" />
      </linearGradient>
    </defs>
    <rect x="-100" y="-100" width="100%" height="100%" fill="` + backgroundColor + `" />
    <polyline id="faceContour" points="` + strings.Join(computedFacePoints, " ") + `" fill="` + faceColor + `" stroke="black" stroke-width="` + fmt.Sprintf("%.2f", 3.0/faceScale) + `" stroke-linejoin="round" filter="url(#fuzzy)" />`

	// Right eye group
	svgString += `
    <g transform="translate(` + fmt.Sprintf("%f %f", center[0]+distanceBetweenEyes+rightEyeOffsetX, -(-center[1]+eyeHeightOffset+rightEyeOffsetY)) + `)">
      <polyline id="rightCountour" points="` + strings.Join(eyeRightContour, " ") + `" fill="white" stroke="white" stroke-width="` + fmt.Sprintf("%.2f", 0.0/faceScale) + `" stroke-linejoin="round" filter="url(#fuzzy)" />
      <polyline id="rightUpper" points="` + strings.Join(eyeRightUpper, " ") + `" fill="none" stroke="black" stroke-width="` + fmt.Sprintf("%.2f", 3.0/faceScale) + `" stroke-linejoin="round" filter="url(#fuzzy)" />
      <polyline id="rightLower" points="` + strings.Join(eyeRightLower, " ") + `" fill="none" stroke="black" stroke-width="` + fmt.Sprintf("%.2f", 4.0/faceScale) + `" stroke-linejoin="round" filter="url(#fuzzy)" />
      ` + rightEyePupilMarkup + `
    </g>`

	// Left eye group
	svgString += `
    <g transform="translate(` + fmt.Sprintf("%f %f", -(center[0]+distanceBetweenEyes+leftEyeOffsetX), -(-center[1]+eyeHeightOffset+leftEyeOffsetY)) + `)">
      <polyline id="leftCountour" points="` + strings.Join(eyeLeftContour, " ") + `" fill="white" stroke="white" stroke-width="` + fmt.Sprintf("%.2f", 0.0/faceScale) + `" stroke-linejoin="round" filter="url(#fuzzy)" />
      <polyline id="leftUpper" points="` + strings.Join(eyeLeftUpper, " ") + `" fill="none" stroke="black" stroke-width="` + fmt.Sprintf("%.2f", 3.0/faceScale) + `" stroke-linejoin="round" filter="url(#fuzzy)" />
      <polyline id="leftLower" points="` + strings.Join(eyeLeftLower, " ") + `" fill="none" stroke="black" stroke-width="` + fmt.Sprintf("%.2f", 4.0/faceScale) + `" stroke-linejoin="round" filter="url(#fuzzy)" />
      ` + leftEyePupilMarkup + `
    </g>`

	// Hairs
	svgString += `
    <g id="hairs">` + func() string {
		hairStrings := make([]string, len(hairs))
		for i, hair := range hairs {
			hairStrings[i] = `<polyline points="` + strings.Join(hair, " ") + `" fill="none" stroke="` + hairColor + `" stroke-width="2" stroke-linejoin="round" filter="url(#fuzzy)" />`
		}
		return strings.Join(hairStrings, "")
	}() + `
    </g>`

	// Nose and mouth
	svgString += noseMarkup + `
    <g id="mouth">
      <polyline points="` + strings.Join(mouthPoints, " ") + `" fill="` + mouthColor + `" stroke="black" stroke-width="3" stroke-linejoin="round" filter="url(#fuzzy)" />
    </g>
  </svg>`
	return svgString

}



func svgS() {
	// Sample data (should be properly initialized in your actual code)
	center := []float64{50.0, 50.0}
	distanceBetweenEyes := 20.0
	rightEyeOffsetX := 10.0
	rightEyeOffsetY := 5.0
	leftEyeOffsetX := 10.0
	leftEyeOffsetY := 5.0
	eyeHeightOffset := 10.0
	eyeRightContour := []string{"10 10", "20 20", "30 30"}
	eyeRightUpper := []string{"11 11", "21 21", "31 31"}
	eyeRightLower := []string{"12 12", "22 22", "32 32"}
	eyeLeftContour := []string{"10 10", "20 20", "30 30"}
	eyeLeftUpper := []string{"11 11", "21 21", "31 31"}
	eyeLeftLower := []string{"12 12", "22 22", "32 32"}
	computedFacePoints := []string{"40 40", "50 50", "60 60"}
	faceScale := 1.0
	hairs := [][]string{{"100 100", "120 120"}, {"130 130", "140 140"}}
	faceColor := "#ececec"
	backgroundColor := "#fff"
	hairColor := "#000"
	mouthColor := "#f00"
	mouthPoints := []string{"150 150", "160 160", "170 170"}
	rightEyePupilMarkup := `<circle cx="15" cy="15" r="3" fill="black"/>`
	leftEyePupilMarkup := `<circle cx="15" cy="15" r="3" fill="black"/>`
	noseMarkup := `<polyline points="130 120, 140 130, 150 120" fill="none" stroke="black"/>`
	dyeColorOffset := "50%"

	var sb strings.Builder
	sb.WriteString(`<svg viewBox="-100 -100 200 200" xmlns="http://www.w3.org/2000/svg" width="500" height="500" id="face-svg">`)
	sb.WriteString(`
		<defs>
			<clipPath id="leftEyeClipPath">
				<polyline points="` + strings.Join(eyeLeftContour, " ") + `" />
			</clipPath>
			<clipPath id="rightEyeClipPath">
				<polyline points="` + strings.Join(eyeRightContour, " ") + `" />
			</clipPath>
			<filter id="fuzzy">
				<feTurbulence id="turbulence" baseFrequency="0.05" numOctaves="3" type="noise" result="noise" />
				<feDisplacementMap in="SourceGraphic" in2="noise" scale="2" />
			</filter>
			<linearGradient id="rainbowGradient" x1="0%" y1="0%" x2="100%" y2="0%">
				<stop offset="0%" style="stop-color: ` + randomRGB() + `; stop-opacity: 1" />
				<stop offset="` + dyeColorOffset + `" style="stop-color: ` + randomRGB() + `; stop-opacity: 1" />
				<stop offset="100%" style="stop-color: ` + randomRGB() + `; stop-opacity: 1" />
			</linearGradient>
		</defs>
		<rect x="-100" y="-100" width="100%" height="100%" fill="` + backgroundColor + `" />
		<polyline id="faceContour" points="` + strings.Join(computedFacePoints, " ") + `" fill="` + faceColor + `" stroke="black" stroke-width="` + strconv.FormatFloat(3.0/faceScale, 'f', 2, 64) + `" stroke-linejoin="round" filter="url(#fuzzy)" />
	`)
	// Right eye group
	sb.WriteString(fmt.Sprintf(`
		<g transform="translate(%f %f)">
			<polyline id="rightCountour" points="%s" fill="white" stroke="white" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
			<polyline id="rightUpper" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
			<polyline id="rightLower" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
			%s
		</g>`, center[0]+distanceBetweenEyes+rightEyeOffsetX, -(-center[1]+eyeHeightOffset+rightEyeOffsetY), strings.Join(eyeRightContour, " "), 0.0/faceScale, strings.Join(eyeRightUpper, " "), 3.0/faceScale, strings.Join(eyeRightLower, " "), 4.0/faceScale, rightEyePupilMarkup))

	// Left eye group
	sb.WriteString(fmt.Sprintf(`
		<g transform="translate(%f %f)">
			<polyline id="leftCountour" points="%s" fill="white" stroke="white" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
			<polyline id="leftUpper" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
			<polyline id="leftLower" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
			%s
		</g>`, -(center[0]+distanceBetweenEyes+leftEyeOffsetX), -(-center[1]+eyeHeightOffset+leftEyeOffsetY), strings.Join(eyeLeftContour, " "), 0.0/faceScale, strings.Join(eyeLeftUpper, " "), 3.0/faceScale, strings.Join(eyeLeftLower, " "), 4.0/faceScale, leftEyePupilMarkup))

	// Hairs
	sb.WriteString(`
		<g id="hairs">`)
	for _, hair := range hairs {
		sb.WriteString(fmt.Sprintf(`<polyline points="%s" fill="none" stroke="%s" stroke-width="2" stroke-linejoin="round" filter="url(#fuzzy)" />`, strings.Join(hair, " "), hairColor))
	}
	sb.WriteString(`
		</g>`)

	// Nose and mouth
	sb.WriteString(noseMarkup)
	sb.WriteString(fmt.Sprintf(`
		<g id="mouth">
			<polyline points="%s" fill="%s" stroke="black" stroke-width="3" stroke-linejoin="round" filter="url(#fuzzy)" />
		</g>
	</svg>`, strings.Join(mouthPoints, " "), mouthColor))

	fmt.Println(sb.String())
}
