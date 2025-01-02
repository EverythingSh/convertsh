package images

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func ToPNG(img image.Image, outputFile string) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG image: %w", err)
	}

	return nil
}
