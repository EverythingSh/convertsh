package images

import (
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToJPEG(img *vips.ImageRef, inputPath, outputPath string) {
	ep := vips.NewJpegExportParams()
	jpegBytes, _, err := img.ExportJpeg(ep)
	if err != nil {
		panic(err)
	}
	if outputPath == "" {
		baseName := filepath.Base(inputPath)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
		outputPath = baseName + ".jpeg"
	}

	err = os.WriteFile(outputPath, jpegBytes, 0644)
	if err != nil {
		panic(err)
	}
}
