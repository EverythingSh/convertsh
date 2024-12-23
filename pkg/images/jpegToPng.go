package images

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"
)

func ConvertJpegToPng(inputFile, outputFile string) error {
	inFile, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer inFile.Close()

	img, err := jpeg.Decode(inFile)
	if err != nil {
		return fmt.Errorf("failed to decode JPEG image: %w", err)
	}

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
