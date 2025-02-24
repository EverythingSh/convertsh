package cmd

import (
	"fmt"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/spf13/cobra"
)

var toFormat string

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:     "image",
	Aliases: []string{"img"},
	Args:    cobra.ExactArgs(1),
	Short:   "Convert JPEG images to PNG format",
	Long:    `Convert JPEG images to PNG format using the image command.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if toFormat == "" {
			return fmt.Errorf("please provide the format to convert to")
		}

		vips.LoggingSettings(nil, vips.LogLevelError)
		vips.Startup(nil)
		defer vips.Shutdown()
		img, err := vips.NewImageFromFile(args[0])
		if err != nil {
			panic(err)
		}

		formatDetect := img.Format().FileExt()
		switch strings.TrimPrefix(formatDetect, ".") {
		case "jpeg":
			fallthrough
		case "jpg":
			fmt.Println("jpeg detected")
		case "png":
			fmt.Println("png detected")
		case "gif":
			fmt.Println("gif detected")
		default:
			fmt.Println("Unsupported format")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)
	imgCmd.Flags().StringVarP(&toFormat, "to", "t", "", "The format to convert to")
}
