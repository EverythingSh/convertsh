package images

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToWEBP(img *vips.ImageRef) {
	ep := vips.NewWebpExportParams()
	webpBytes, _, err := img.ExportWebp(ep)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output_con.webp", webpBytes, 0644)
	if err != nil {
		panic(err)
	}
}
