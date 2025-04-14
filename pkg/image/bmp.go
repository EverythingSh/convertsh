package images

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
	"golang.org/x/image/bmp"
)

func ToBMP(img *vips.ImageRef, inputPath string) {
	imgImg, err := img.ToImage(&vips.ExportParams{Format: vips.ImageTypePNG, Quality: 100})
	if err != nil {
		panic(fmt.Errorf("failed to convert to image: %w", err))
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
	err = bmp.Encode(outputFile, imgImg)
	if err != nil {
		panic(fmt.Errorf("failed to encode as BMP: %w", err))
	}
}
