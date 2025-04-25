package images

import (
	"os"

	"github.com/EverythingSh/convertsh/pkg/util"
	"github.com/davidbyttow/govips/v2/vips"
)

func ToWEBP(img *vips.ImageRef, inputPath, outputPath string, isBulk bool) {
	ep := vips.NewWebpExportParams()
	webpBytes, _, err := img.ExportWebp(ep)
	if err != nil {
		panic(err)
	}

	if outputPath == "" {
		outputPath = util.CreateOutputPath(inputPath, ".webp", isBulk)
	}

	err = os.WriteFile(outputPath, webpBytes, 0644)
	if err != nil {
		panic(err)
	}
}
