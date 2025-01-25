package cmd

import (
	"fmt"

	"github.com/EverythingSh/convertsh/internal/converter/detect"
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
		format := detect.ImageType(args[0])

		if format == "" {
			return fmt.Errorf("could not detect image type")
		}

		fmt.Println("TYPE:", format)
		// jpegImage := args[0]
		// var pngImage string
		// if strings.HasSuffix(jpegImage, ".jpeg") {
		// 	pngImage = strings.TrimSuffix(jpegImage, ".jpeg") + ".png"
		// } else {
		// 	pngImage = strings.TrimSuffix(jpegImage, ".jpg") + ".png"
		// }

		// if toFormat == "" {
		// 	toFormat = "png"
		// }

		// jpegConverter := images.NewJPEGConverter(toFormat, &converter.ConversionOptions{
		// 	Quality:          100,
		// 	CompressionLevel: 0,
		// 	Metadata:         &converter.ImageMetadata{},
		// })

		// err := jpegConverter.Convert(jpegImage, pngImage)
		// if err != nil {
		// 	return err
		// }

		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)
	imgCmd.Flags().StringVarP(&toFormat, "to", "t", "", "The format to convert to")
}
