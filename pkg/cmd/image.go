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
		fmt.Printf("converting %s to %s\n", args[0], toFormat)
		img, err := vips.NewImageFromFile(args[0])
		if err != nil {
			panic(err)
		}

		if img.Format().FileExt()[1:] == toFormat && !(strings.HasSuffix(args[0], "dng") && toFormat == "tiff") {
			fmt.Println("already in the desired format")
			return nil
		}

		switch toFormat {
		case "jpeg":
			fallthrough
		case "jpg":
			fmt.Println("converting to jpeg")
			images.ToJPEG(img, args[0], "")
		case "png":
			fmt.Println("coverting to png")
			images.ToPNG(img, args[0], "")
		case "gif":
			fmt.Println("converting to gif")
			images.ToGIF(img, args[0], "")
		case "tiff":
			fmt.Println("converting to tiff")
			images.ToTIFF(img, args[0], "")
		case "bmp":
			fmt.Println("converting to bmp")
			images.ToBMP(img, args[0], "")
		case "webp":
			fmt.Println("converting to webp")
			images.ToWEBP(img, args[0], "")
		case "heif":
			fmt.Println("converting to heif")
			images.ToHEIF(img, args[0], "")
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
