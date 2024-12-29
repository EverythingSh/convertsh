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
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		jpegImage := args[0]
		var pngImage string
		if strings.HasSuffix(jpegImage, ".jpeg") {
			pngImage = strings.TrimSuffix(jpegImage, ".jpeg") + ".png"
		} else {
			pngImage = strings.TrimSuffix(jpegImage, ".jpg") + ".png"
		}

		err := images.ConvertJpegToPng(jpegImage, pngImage)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
