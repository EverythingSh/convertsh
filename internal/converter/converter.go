package converter

type BaseConverter struct {
	sourceFormat ImageFormat
	targetFormat ImageFormat
	options      *ConversionOptions
}

func NewBaseConverter(sourceFormat, targetFormat ImageFormat, options *ConversionOptions) *BaseConverter {
	return &BaseConverter{
		sourceFormat: sourceFormat,
		targetFormat: targetFormat,
		options:      options,
	}
}
