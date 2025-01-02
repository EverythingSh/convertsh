package converter

type BaseConverter struct {
	SourceFormat ImageFormat
	TargetFormat ImageFormat
	Options      *ConversionOptions
}

func NewBaseConverter(sourceFormat, targetFormat ImageFormat, options *ConversionOptions) *BaseConverter {
	return &BaseConverter{
		SourceFormat: sourceFormat,
		TargetFormat: targetFormat,
		Options:      options,
	}
}
