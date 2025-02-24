package images

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func Test() {
	vips.Startup(nil)
	defer vips.Shutdown()
	img, err := vips.NewImageFromFile("assets/Cat.jpg")
	if err != nil {
		panic(err)
	}

	err = img.AutoRotate()
	if err != nil {
		panic(err)
	}

	ep := vips.NewPngExportParams()
	pngBytes, _, err := img.ExportPng(ep)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output.png", pngBytes, 0644)
	if err != nil {
		panic(err)
	}
}
