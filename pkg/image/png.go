package images

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToPNG(img *vips.ImageRef) {
	ep := vips.NewPngExportParams()
	ep.Compression = 0
	ep.Quality = 100

	pngBytes, _, err := img.ExportPng(ep)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output_con.png", pngBytes, 0644)
	if err != nil {
		panic(err)
	}
}
