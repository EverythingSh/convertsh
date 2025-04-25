package images

import (
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToWEBP(img *vips.ImageRef, inputPath, outputPath string) {
	ep := vips.NewWebpExportParams()
	webpBytes, _, err := img.ExportWebp(ep)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		baseName := filepath.Base(inputPath)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
		outputPath = baseName + ".webp"
	}

	err = os.WriteFile(outputPath, webpBytes, 0644)
	if err != nil {
		panic(err)
	}
}
