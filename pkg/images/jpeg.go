package images

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/EverythingSh/convertsh/internal/converter"
)

type JPEGConverter struct {
	con *converter.BaseConverter
}

func NewJPEGConverter(toFormat string, options *converter.ConversionOptions) *JPEGConverter {
	return &JPEGConverter{
		con: converter.NewBaseConverter(converter.JPEG, converter.ImageFormat(toFormat), options),
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

	bounds := img.Bounds()
	j.con.Options.Metadata.Height = bounds.Dy()
	j.con.Options.Metadata.Width = bounds.Dx()
	j.con.Options.Metadata.Format = converter.JPEG

	switch j.con.TargetFormat {
	case converter.PNG:
		err = ToPNG(img, outputFile)
		if err != nil {
			return fmt.Errorf("failed to convert to PNG: %w", err)
		}
	}
	return nil
}

func ToJPEG(img image.Image, outputFile string) error {
	outFile, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		return fmt.Errorf("failed to encode PNG image: %w", err)
	}

	return nil
}
