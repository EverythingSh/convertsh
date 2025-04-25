package images

import (
	"fmt"
	"os"

	"github.com/EverythingSh/convertsh/pkg/util"
	"github.com/davidbyttow/govips/v2/vips"
	"golang.org/x/image/bmp"
)

func ToBMP(img *vips.ImageRef, inputPath, outputPath string, isBulk bool) {
	imgImg, err := img.ToImage(&vips.ExportParams{Format: vips.ImageTypePNG, Quality: 100})
	if err != nil {
		panic(fmt.Errorf("failed to convert to image: %w", err))
	}

	// Create output filename
	if outputPath == "" {
		outputPath = util.CreateOutputPath(inputPath, ".bmp", isBulk)
	}
	// Create output file
	outputFile, err := os.Create(outputPath)
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
