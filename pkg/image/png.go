package images

import (
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToPNG(img *vips.ImageRef, inputPath, outputPath string) {
	ep := vips.NewPngExportParams()
	ep.Compression = 0
	ep.Quality = 100

	pngBytes, _, err := img.ExportPng(ep)
	if err != nil {
		panic(err)
	}
	if outputPath == "" {
		baseName := filepath.Base(inputPath)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
		outputPath = baseName + ".png"
	}

	err = os.WriteFile(outputPath, pngBytes, 0644)
	if err != nil {
		panic(err)
	}
}
