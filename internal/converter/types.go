package converter

type ImageFormat string

const (
	JPEG ImageFormat = "jpeg"
	PNG  ImageFormat = "png"
)

type Converter interface {
	Convert(inputFile, outputFile string) error
}

type ImageMetadata struct {
	Width  int
	Height int
	Format ImageFormat
}

type ConversionOptions struct {
	Quality          int
	CompressionLevel int
	Metadata         *ImageMetadata
}
