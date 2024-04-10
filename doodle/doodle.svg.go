package doodle

import (
	"fmt"
	"strings"
)

func doodlesvg() string {
	// Assuming these variables are defined somewhere in your Go code.
	var (
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

	// Constructing the SVG string.
	svgString := fmt.Sprintf(`<svg viewBox="-100 -100 200 200" xmlns="http://www.w3.org/2000/svg" width="500" height="500" id="face-svg">
    <defs>
      <clipPath id="leftEyeClipPath">
        <polyline points="%s" />
      </clipPath>
      <clipPath id="rightEyeClipPath">
        <polyline points="%s" />
      </clipPath>
      <filter id="fuzzy">
        <feTurbulence id="turbulence" baseFrequency="0.05" numOctaves="3" type="noise" result="noise" />
        <feDisplacementMap in="SourceGraphic" in2="noise" scale="2" />
      </filter>
      <linearGradient id="rainbowGradient" x1="0%%" y1="0%%" x2="100%%" y2="0%%">
        <stop offset="0%%" style="stop-color: %s; stop-opacity: 1" />
        <stop offset="%s" style="stop-color: %s; stop-opacity: 1" />
        <stop offset="100%%" style="stop-color: %s; stop-opacity: 1" />
      </linearGradient>
    </defs>
    <rect x="-100" y="-100" width="100%%" height="100%%" fill="%s" />
    <polyline id="faceContour" points="%s" fill="%s" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
    <g transform="translate(%f %f)">
      <polyline id="rightCountour" points="%s" fill="white" stroke="white" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
    </g>
    <g transform="translate(%f %f)">
      <polyline id="leftCountour" points="%s" fill="white" stroke="white" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
    </g>
    <g transform="translate(%f %f)">
      <polyline id="rightUpper" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
      <polyline id="rightLower" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
      %s
      </g>
    <g transform="translate(%f %f)">
      <polyline id="leftUpper" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
      <polyline id="leftLower" points="%s" fill="none" stroke="black" stroke-width="%f" stroke-linejoin="round" filter="url(#fuzzy)" />
      %s
    </g>
    <g id="hairs">
      %s
    </g>
    %s
    <g id="mouth">
      <polyline points="%s" fill="%s" stroke="black" stroke-width="3" stroke-linejoin="round" filter="url(#fuzzy)" />
    </g>
  </svg>`, joinWithSpace(eyeLeftContour), joinWithSpace(eyeRightContour),
		randomRGB(), dyeColorOffset, randomRGB(), randomRGB(),
		backgroundColor,
		joinWithSpace([]string{}), // Place where computedFacePoints.join(' ') should be used.
		faceColor, 3.0/faceScale,
		center[0]+distanceBetweenEyes+rightEyeOffsetX, -(-center[1] + eyeHeightOffset + rightEyeOffsetY), joinWithSpace(eyeRightContour), 0.0/faceScale,
		-(center[0] + distanceBetweenEyes + leftEyeOffsetX), -(-center[1] + eyeHeightOffset + leftEyeOffsetY), joinWithSpace(eyeLeftContour), 0.0/faceScale,
		center[0]+distanceBetweenEyes+rightEyeOffsetX, -(-center[1] + eyeHeightOffset + rightEyeOffsetY), joinWithSpace(eyeRightUpper), 3.0/faceScale, joinWithSpace(eyeRightLower), 4.0/faceScale, rightEyePupilMarkup,
		-(center[0] + distanceBetweenEyes + leftEyeOffsetX), -(-center[1] + eyeHeightOffset + leftEyeOffsetY), joinWithSpace(eyeLeftUpper), 4.0/faceScale, joinWithSpace(eyeLeftLower), 4.0/faceScale, leftEyePupilMarkup,
		hairsMarkup,
		noseMarkup,
		joinWithSpace(mouthPoints), mouthColor,
	)
	return svgString

}
