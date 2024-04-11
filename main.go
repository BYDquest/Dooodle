// package main

// import (
// 	"fmt"
// 	"github.com/BYDquest/Dooodle/doodle"
// )

// func main() {

	
// 	face := doodle.GenerateFaceContourPoints(100)
// 	fmt.Println(face.Center)


// 	leftEye, rightEye := doodle.GenerateBothEyes(50)
// 	fmt.Println("Left Eye Upper Points:", leftEye.Upper)
// 	fmt.Println("Right Eye Upper Points:", rightEye.Upper)
// }






package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func convertSVGtoPNG(inputFolder, outputFolder string) error {
	files, err := os.ReadDir(inputFolder)
	if err != nil {
		return fmt.Errorf("error reading input folder: %w", err)
	}

	// Regular expression to find the rect element and extract its fill attribute
	rectRegex := regexp.MustCompile(`<rect[^>]+fill="rgb\((\d+),\s*(\d+),\s*(\d+)\)"[^>]*>`)

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".svg") {
			continue // Skip non-SVG files
		}

		inputPath := filepath.Join(inputFolder, fileName)
		svgContent, err := os.ReadFile(inputPath)
		if err != nil {
			return fmt.Errorf("failed to read SVG file %s: %w", fileName, err)
		}

		// Attempt to find the rect element and extract its fill color
		matches := rectRegex.FindStringSubmatch(string(svgContent))
		var backgroundColor color.RGBA
		if len(matches) == 4 {
			r, _ := strconv.Atoi(matches[1])
			g, _ := strconv.Atoi(matches[2])
			b, _ := strconv.Atoi(matches[3])
			backgroundColor = color.RGBA{uint8(r), uint8(g), uint8(b), 255} // Assuming full opacity
		} else {
			// Default background color if no rect element is found
			backgroundColor = color.RGBA{255, 255, 255, 255} // White
		}

		icon, err := oksvg.ReadIconStream(strings.NewReader(string(svgContent)))
		if err != nil {
			return fmt.Errorf("failed to parse SVG file %s: %w", fileName, err)
		}

		if icon.ViewBox.W <= 0 || icon.ViewBox.H <= 0 {
			fmt.Printf("Invalid or missing viewBox for %s; skipping file.\n", fileName)
			continue
		}

		icon.SetTarget(0, 0, float64(icon.ViewBox.W), float64(icon.ViewBox.H))
		img := image.NewRGBA(image.Rect(0, 0, int(icon.ViewBox.W), int(icon.ViewBox.H)))

		// Fill the image with the extracted background color
		for y := 0; y < int(icon.ViewBox.H); y++ {
			for x := 0; x < int(icon.ViewBox.W); x++ {
				img.Set(x, y, backgroundColor)
			}
		}

		scannerGV := rasterx.NewScannerGV(int(icon.ViewBox.W), int(icon.ViewBox.H), img, img.Bounds())
		rasterizer := rasterx.NewDasher(int(icon.ViewBox.W), int(icon.ViewBox.H), scannerGV)

		icon.Draw(rasterizer, 1.0)

		outputPath := filepath.Join(outputFolder, strings.TrimSuffix(fileName, ".svg")+".png")
		outputFile, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("failed to create PNG file %s: %w", outputPath, err)
		}
		defer outputFile.Close()

		if err := png.Encode(outputFile, img); err != nil {
			return fmt.Errorf("failed to encode PNG for %s: %w", outputPath, err)
		}

		fmt.Println("Converted:", fileName, "to PNG")
	}

	return nil
}

func main() {
	inputFolder := "./avatar"
	outputFolder := "./avatar-png"

	if err := convertSVGtoPNG(inputFolder, outputFolder); err != nil {
		fmt.Println("Error converting SVG files:", err)
	}
}



