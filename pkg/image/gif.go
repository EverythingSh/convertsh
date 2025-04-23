package images

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToGIF(img *vips.ImageRef, inputPath, outputPath string) {
	ep := vips.NewGifExportParams()
	gifBytes, _, err := img.ExportGIF(ep)
	if err != nil {
		fmt.Println("error in gif export")
		panic(err)
	}

	// Create output filename
	if outputPath == "" {
		baseName := filepath.Base(inputPath)
		baseName = baseName[:len(baseName)-len(filepath.Ext(baseName))]
		outputPath = baseName + ".bmp"
	}

	err = os.WriteFile(outputPath, gifBytes, 0644)
	if err != nil {
		fmt.Println("error in writing gif file")
		panic(err)
	}
}
