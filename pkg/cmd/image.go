package cmd

import (
	"fmt"
	"strings"

	images "github.com/EverythingSh/convertsh/pkg/image"
	"github.com/EverythingSh/convertsh/pkg/types"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/spf13/cobra"
)

var toFormat string

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:     "image",
	Aliases: []string{"img"},
	Args:    cobra.ExactArgs(1),
	Short:   "Convert any images to a different format",
	Long:    `This command can convert any images to a different format`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if toFormat == "" {
			return fmt.Errorf("please provide the format to convert to")
		}

		toFormat = strings.ToLower(toFormat)

		var isConvertible bool
		for _, format := range types.RasterFormats {
			if toFormat == string(format) {
				isConvertible = true
				break
			}
		}

		if !isConvertible {
			return fmt.Errorf("unsupported format")
		}

		vips.LoggingSettings(nil, vips.LogLevelError)
		vips.Startup(nil)
		defer vips.Shutdown()
		img, err := vips.NewImageFromFile(args[0])
		if err != nil {
			panic(err)
		}

		if img.Format().FileExt()[1:] == toFormat {
			fmt.Println("already in the desired format")
			return nil
		}

		switch toFormat {
		case "jpeg":
			fallthrough
		case "jpg":
			fmt.Println("converting to jpeg")
			images.ToJPEG(img)
		case "png":
			fmt.Println("coverting to png")
			images.ToPNG(img)
		case "gif":
			fmt.Println("converting to gif")
			images.ToGIF(img)
		case "tiff":
			fmt.Println("converting to tiff")
			images.ToTIFF(img)
		case "bmp":
			fmt.Println("converting to bmp")
			images.ToBMP(args[0])
		case "webp":
			fmt.Println("converting to webp")
			images.ToWEBP(img)
		case "heif":
			fmt.Println("converting to heif")
			images.ToHEIF(img)
		case "heic":
			fmt.Println("converting to heic")
			// images.ToHEIC(img)
		default:
			fmt.Println("unsupported format")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)
	imgCmd.Flags().StringVarP(&toFormat, "to", "t", "", "The format to convert to")
}
