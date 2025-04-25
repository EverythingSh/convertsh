package images

import (
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToHEIF(img *vips.ImageRef, inputPath, outputPath string) {
	ep := vips.NewHeifExportParams()
	heifBytes, _, err := img.ExportHeif(ep)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		baseName := filepath.Base(inputPath)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
		outputPath = baseName + ".heif"
	}

	err = os.WriteFile(outputPath, heifBytes, 0644)
	if err != nil {
		panic(err)
	}
}
