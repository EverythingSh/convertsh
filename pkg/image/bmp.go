package images

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
)

func ToBMP(inputPath string) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		panic(fmt.Errorf("failed to open input file: %w", err))
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		panic(fmt.Errorf("failed to decode image: %w", err))
	}

	// Create output filename
	baseName := filepath.Base(inputPath)
	baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
	outputFilename := baseName + ".bmp"

	// Create output file
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		panic(fmt.Errorf("failed to create output file: %w", err))
	}
	defer outputFile.Close()

	// Encode as BMP
	err = bmp.Encode(outputFile, img)
	if err != nil {
		panic(fmt.Errorf("failed to encode as BMP: %w", err))
	}
}
