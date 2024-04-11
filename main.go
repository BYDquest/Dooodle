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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func drawBackground(img *image.RGBA, bgColor color.Color) {
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, bgColor)
		}
	}
}

func convertSVGtoPNG(inputFolder, outputFolder string) error {
	files, err := os.ReadDir(inputFolder)
	if err != nil {
		return fmt.Errorf("error reading input folder: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		fileName := file.Name()
		if !strings.HasSuffix(fileName, ".svg") {
			continue // Skip non-SVG files
		}

		inputPath := filepath.Join(inputFolder, fileName)
		svgContent, err := ioutil.ReadFile(inputPath)
		if err != nil {
			return fmt.Errorf("failed to read SVG file %s: %w", fileName, err)
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

		// Set the background color here. Change the color as needed.
		backgroundColor := color.RGBA{255, 255, 255, 255} // White background
		drawBackground(img, backgroundColor)

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
