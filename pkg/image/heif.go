package images

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToHEIF(img *vips.ImageRef) {
	ep := vips.NewHeifExportParams()
	heifBytes, _, err := img.ExportHeif(ep)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output_con.heif", heifBytes, 0644)
	if err != nil {
		panic(err)
	}
}
