package detect

import (
	"strings"

	"github.com/EverythingSh/convertsh/internal/converter"
)

func ImageType(path string) string {
	ext := strings.ToLower(path[strings.LastIndex(path, ".")+1:])
	for _, format := range converter.RasterFormats {
		if ext == string(format) {
			return string(format)
		}
	}

	for _, format := range converter.VectorFormats {
		if ext == string(format) {
			return string(format)
		}
	}

	return ""
}
