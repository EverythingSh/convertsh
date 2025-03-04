package images

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToTIFF(img *vips.ImageRef) {
	ep := vips.NewTiffExportParams()
	tiffBytes, _, err := img.ExportTiff(ep)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output_con.tiff", tiffBytes, 0644)
	if err != nil {
		panic(err)
	}
}
