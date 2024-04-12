package lib

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
	"sync"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

var (
	// Updated to match both rgb() and hex color formats
	rectRegex = regexp.MustCompile(`<rect[^>]+fill="(?:rgb\((\d+),\s*(\d+),\s*(\d+)\)|#([0-9A-Fa-f]{6}))"[^>]*>`)
)

func convertSVGtoPNG(inputPath, outputPath string, wg *sync.WaitGroup, errChan chan error) {
	defer wg.Done()

	svgContent, err := os.ReadFile(inputPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to read SVG file %s: %w", inputPath, err)
		return
	}

	// Extracts the fill color from the SVG content
	matches := rectRegex.FindStringSubmatch(string(svgContent))
	var backgroundColor color.RGBA
	if len(matches) > 1 {
		if matches[1] != "" && matches[2] != "" && matches[3] != "" {
			// RGB format
			r, _ := strconv.Atoi(matches[1])
			g, _ := strconv.Atoi(matches[2])
			b, _ := strconv.Atoi(matches[3])
			backgroundColor = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		} else if matches[4] != "" {
			// Hex format
			hexColor := matches[4]
			r, _ := strconv.ParseUint(hexColor[0:2], 16, 8)
			g, _ := strconv.ParseUint(hexColor[2:4], 16, 8)
			b, _ := strconv.ParseUint(hexColor[4:6], 16, 8)
			backgroundColor = color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		} else {
			backgroundColor = color.RGBA{255, 255, 255, 255} // Default to white if not found
		}
	} else {
		backgroundColor = color.RGBA{255, 255, 255, 255} // Default to white
	}

	icon, err := oksvg.ReadIconStream(strings.NewReader(string(svgContent)))
	if err != nil {
		errChan <- fmt.Errorf("failed to parse SVG file %s: %w", inputPath, err)
		return
	}

	if icon.ViewBox.W <= 0 || icon.ViewBox.H <= 0 {
		errChan <- fmt.Errorf("invalid or missing viewBox for %s; skipping file", inputPath)
		return
	}

	icon.SetTarget(0, 0, float64(icon.ViewBox.W), float64(icon.ViewBox.H))
	img := image.NewRGBA(image.Rect(0, 0, int(icon.ViewBox.W), int(icon.ViewBox.H)))

	for y := 0; y < int(icon.ViewBox.H); y++ {
		for x := 0; x < int(icon.ViewBox.W); x++ {
			img.Set(x, y, backgroundColor)
		}
	}

	scannerGV := rasterx.NewScannerGV(int(icon.ViewBox.W), int(icon.ViewBox.H), img, img.Bounds())
	rasterizer := rasterx.NewDasher(int(icon.ViewBox.W), int(icon.ViewBox.H), scannerGV)

	icon.Draw(rasterizer, 1.0)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to create PNG file %s: %w", outputPath, err)
		return
	}
	defer outputFile.Close()

	if err := png.Encode(outputFile, img); err != nil {
		errChan <- fmt.Errorf("failed to encode PNG for %s: %w", outputPath, err)
		return
	}

	fmt.Println("Converted:", inputPath, "to PNG")
}

func Svg2png() {

	inputFolder := "./avatar"
	outputFolder := "./avatar-png"
	files, err := os.ReadDir(inputFolder)
	if err != nil {
		fmt.Println("Error reading input folder:", err)
		return
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(files))
	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".svg") {
			continue
		}

		inputPath := filepath.Join(inputFolder, file.Name())
		outputPath := filepath.Join(outputFolder, strings.TrimSuffix(file.Name(), ".svg")+".png")

		wg.Add(1)
		go convertSVGtoPNG(inputPath, outputPath, &wg, errChan)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
