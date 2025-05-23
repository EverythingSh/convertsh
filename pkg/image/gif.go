package images

import (
	"fmt"
	"os"

	"github.com/EverythingSh/convertsh/pkg/util"
	"github.com/davidbyttow/govips/v2/vips"
)

func ToGIF(img *vips.ImageRef, inputPath, outputPath string, isBulk bool) {
	ep := vips.NewGifExportParams()
	gifBytes, _, err := img.ExportGIF(ep)
	if err != nil {
		fmt.Println("error in gif export")
		panic(err)
	}

	// Create output filename
	if outputPath == "" {
		outputPath = util.CreateOutputPath(inputPath, ".gif", isBulk)
	}

	err = os.WriteFile(outputPath, gifBytes, 0644)
	if err != nil {
		fmt.Println("error in writing gif file")
		panic(err)
	}
}
