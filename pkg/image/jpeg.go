package images

import (
	"os"

	"github.com/EverythingSh/convertsh/pkg/util"
	"github.com/davidbyttow/govips/v2/vips"
)

func ToJPEG(img *vips.ImageRef, inputPath, outputPath string, isBulk bool) {
	ep := vips.NewJpegExportParams()
	jpegBytes, _, err := img.ExportJpeg(ep)
	if err != nil {
		panic(err)
	}
	if outputPath == "" {
		outputPath = util.CreateOutputPath(inputPath, ".jpeg", isBulk)
	}

	err = os.WriteFile(outputPath, jpegBytes, 0644)
	if err != nil {
		panic(err)
	}
}
