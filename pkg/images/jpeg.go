package images

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/EverythingSh/convertsh/internal/converter"
)

type JPEGConverter struct {
	*converter.BaseConverter
}

func NewJPEGConverter(options *converter.ConversionOptions) *JPEGConverter {
	return &JPEGConverter{
		BaseConverter: converter.NewBaseConverter(converter.JPEG, converter.PNG, options),
	}
}

func (j *JPEGConverter) Convert(inputFile, outputFile string) error {
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
