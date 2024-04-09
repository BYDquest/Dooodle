package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "io/ioutil"
    "os"
    "path/filepath"
    "sync"

    "github.com/srwiley/oksvg"
    "github.com/srwiley/rasterx"
)

// Function to process a batch of SVG images
func processBatch(batchFiles []string, imageDir, outputDir string, batchNumber, gridWidth, outputImageWidth, outputImageHeight int, wg *sync.WaitGroup) {
    defer wg.Done()

    combinedImageWidth := gridWidth * outputImageWidth
    combinedImageHeight := (len(batchFiles) / gridWidth) * outputImageHeight
    if len(batchFiles)%gridWidth != 0 {
        combinedImageHeight += outputImageHeight
    }

    canvas := image.NewRGBA(image.Rect(0, 0, combinedImageWidth, combinedImageHeight))
    draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

    for index, file := range batchFiles {
        x := (index % gridWidth) * outputImageWidth
        y := (index / gridWidth) * outputImageHeight

        if file != "" {
            imgPath := filepath.Join(imageDir, file)
            svgFile, err := os.Open(imgPath)
            if err != nil {
                fmt.Printf("Error opening SVG file %s: %v\n", imgPath, err)
                continue
            }
            defer svgFile.Close()

            icon, err := oksvg.ReadIconStream(svgFile)
            if err != nil {
                fmt.Printf("Error reading SVG image %s: %v\n", imgPath, err)
                continue
            }

            // Adjust the viewbox to fit the desired output dimensions
            // This ensures the SVG scales correctly to fill the space

            icon.SetTarget(0, 0, float64(outputImageWidth), float64(outputImageHeight))

            // Prepare the rasterizer to draw the scaled SVG
            rgba := image.NewRGBA(image.Rect(0, 0, outputImageWidth, outputImageHeight))
            sc := rasterx.NewScannerGV(outputImageWidth, outputImageHeight, rgba, rgba.Bounds())
            rast := rasterx.NewDasher(outputImageWidth, outputImageHeight, sc)
            icon.Draw(rast, 1)

            rect := image.Rect(x, y, x+outputImageWidth, y+outputImageHeight)
            draw.Draw(canvas, rect, rgba, image.Point{}, draw.Over)
        }
    }

    outputPath := filepath.Join(outputDir, fmt.Sprintf("%d.png", batchNumber))
    file, err := os.Create(outputPath)
    if err != nil {
        fmt.Printf("Error creating output file %s: %v\n", outputPath, err)
        return
    }
    defer file.Close()

    err = png.Encode(file, canvas)
    if err != nil {
        fmt.Printf("Error encoding PNG image %s: %v\n", outputPath, err)
        return
    }

    fmt.Printf("Combined image %d created in %s directory.\n", batchNumber, outputDir)
}

// Function to combine SVG images with concurrency control
func combineSVGImages(imageDir, outputDir string, gridWidth, outputImageWidth, outputImageHeight, concurrencyLimit int) error {
    err := os.MkdirAll(outputDir, os.ModePerm)
    if err != nil {
        return fmt.Errorf("error creating output directory: %v", err)
    }

    imageFiles, err := ioutil.ReadDir(imageDir)
    if err != nil {
        return fmt.Errorf("error reading image directory: %v", err)
    }

    var svgFiles []string
    for _, file := range imageFiles {
        if filepath.Ext(file.Name()) == ".svg" {
            svgFiles = append(svgFiles, file.Name())
        }
    }

    imagesPerCombined := gridWidth * gridWidth
    batches := make([][]string, (len(svgFiles)+imagesPerCombined-1)/imagesPerCombined)

    for i := 0; i < len(svgFiles); i += imagesPerCombined {
        end := i + imagesPerCombined
        if end > len(svgFiles) {
            end = len(svgFiles)
        }
        batches[i/imagesPerCombined] = svgFiles[i:end]
    }

    var wg sync.WaitGroup
    semaphore := make(chan struct{}, concurrencyLimit)

    for i, batchFiles := range batches {
        wg.Add(1)
        semaphore <- struct{}{}
        go func(batchFiles []string, batchNumber int) {
            processBatch(batchFiles, imageDir, outputDir, batchNumber, gridWidth, outputImageWidth, outputImageHeight, &wg)
            <-semaphore
        }(batchFiles, i)
    }

    wg.Wait()

    return nil
}

func main() {
    err := combineSVGImages("avatar3", "combined-images", 10, 100, 100, 4)
    if err != nil {
        fmt.Printf("Error combining SVG images: %v\n", err)
    }
}
