package images

import (
	"os"

	"github.com/EverythingSh/convertsh/pkg/util"
	"github.com/davidbyttow/govips/v2/vips"
)

func ToHEIF(img *vips.ImageRef, inputPath, outputPath string, isBulk bool) {
	ep := vips.NewHeifExportParams()
	heifBytes, _, err := img.ExportHeif(ep)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		outputPath = util.CreateOutputPath(inputPath, ".heif", isBulk)
	}

	err = os.WriteFile(outputPath, heifBytes, 0644)
	if err != nil {
		panic(err)
	}
}
