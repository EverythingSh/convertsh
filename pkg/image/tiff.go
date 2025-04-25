package images

import (
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToTIFF(img *vips.ImageRef, inputPath, outputPath string) {
	ep := vips.NewTiffExportParams()
	tiffBytes, _, err := img.ExportTiff(ep)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		baseName := filepath.Base(inputPath)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
		outputPath = baseName + ".tiff"
	}
	err = os.WriteFile(outputPath, tiffBytes, 0644)
	if err != nil {
		panic(err)
	}
}
