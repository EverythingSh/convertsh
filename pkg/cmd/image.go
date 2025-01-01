package cmd

import (
	"strings"

	"github.com/EverythingSh/convertsh/pkg/images"
	"github.com/spf13/cobra"
)

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:     "image",
	Aliases: []string{"img"},
	Args:    cobra.ExactArgs(1),
	Short:   "Convert JPEG images to PNG format",
	Long:    `Convert JPEG images to PNG format using the image command.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		jpegImage := args[0]
		var pngImage string
		if strings.HasSuffix(jpegImage, ".jpeg") {
			pngImage = strings.TrimSuffix(jpegImage, ".jpeg") + ".png"
		} else {
			pngImage = strings.TrimSuffix(jpegImage, ".jpg") + ".png"
		}

		jpegConverter := images.NewJPEGConverter(nil)
		err := jpegConverter.Convert(jpegImage, pngImage)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)
}
