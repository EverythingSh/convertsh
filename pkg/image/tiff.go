package images

import (
	"os"

	"github.com/EverythingSh/convertsh/pkg/util"
	"github.com/davidbyttow/govips/v2/vips"
)

func ToTIFF(img *vips.ImageRef, inputPath, outputPath string, isBulk bool) {
	ep := vips.NewTiffExportParams()
	tiffBytes, _, err := img.ExportTiff(ep)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		outputPath = util.CreateOutputPath(inputPath, ".tiff", isBulk)
	}
	err = os.WriteFile(outputPath, tiffBytes, 0644)
	if err != nil {
		panic(err)
	}
}
