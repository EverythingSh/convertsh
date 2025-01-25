package converter

type ImageRasterFormat string
type ImageVectorFormat string

const (
	JPEG ImageRasterFormat = "jpeg"
	PNG  ImageRasterFormat = "png"
	JPG  ImageRasterFormat = "jpg"
	GIF  ImageRasterFormat = "gif"
	TIFF ImageRasterFormat = "tiff"
	TIF  ImageRasterFormat = "tif"
	BMP  ImageRasterFormat = "bmp"
	SVG  ImageVectorFormat = "svg"
	AI   ImageVectorFormat = "ai"
	EPS  ImageVectorFormat = "eps"
	WEBP ImageRasterFormat = "webp"
	HEIF ImageRasterFormat = "heif"
	HEIC ImageRasterFormat = "heic"
	AVIF ImageRasterFormat = "avif"
	RAW  ImageRasterFormat = "raw"
	APNG ImageRasterFormat = "apng"
)

var RasterFormats = []ImageRasterFormat{
	JPEG, PNG, JPG, GIF, TIFF, TIF, BMP, WEBP, HEIF, HEIC, AVIF, RAW, APNG,
}

var VectorFormats = []ImageVectorFormat{
	SVG, AI, EPS,
}

type Converter interface {
	Convert(inputFile, outputFile string) error
}

type ImageMetadata struct {
	Width  int
	Height int
	// Format ImageFormat
}

type ConversionOptions struct {
	Quality          int
	CompressionLevel int
	Metadata         *ImageMetadata
}
