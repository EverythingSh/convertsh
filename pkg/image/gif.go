package images

import (
	"fmt"
	"os"

	"github.com/davidbyttow/govips/v2/vips"
)

func ToGIF(img *vips.ImageRef) {
	ep := vips.NewGifExportParams()
	gifBytes, _, err := img.ExportGIF(ep)
	if err != nil {
		fmt.Println("error in gif export")
		panic(err)
	}

	err = os.WriteFile("output_con.gif", gifBytes, 0644)
	if err != nil {
		fmt.Println("error in writing gif file")
		panic(err)
	}
}
