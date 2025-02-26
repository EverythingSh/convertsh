package images

import (
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToJPEG(img *vips.ImageRef) {
	ep := vips.NewJpegExportParams()
	jpegBytes, _, err := img.ExportJpeg(ep)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("output_con.jpg", jpegBytes, 0644)
	if err != nil {
		panic(err)
	}
}
