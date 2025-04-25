package images

import (
	"os"

	"github.com/EverythingSh/convertsh/pkg/util"
	"github.com/davidbyttow/govips/v2/vips"
)

func ToPNG(img *vips.ImageRef, inputPath, outputPath string, isBulk bool) {
	ep := vips.NewPngExportParams()
	ep.Compression = 0
	ep.Quality = 100

	pngBytes, _, err := img.ExportPng(ep)
	if err != nil {
		panic(err)
	}
	if outputPath == "" {
		outputPath = util.CreateOutputPath(inputPath, ".png", isBulk)
	}

	err = os.WriteFile(outputPath, pngBytes, 0644)
	if err != nil {
		panic(err)
	}
}
