package cmd

import (
	"fmt"
	"os"
	"strings"

	images "github.com/EverythingSh/convertsh/pkg/image"
	"github.com/EverythingSh/convertsh/pkg/tui"
	"github.com/EverythingSh/convertsh/pkg/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "con",
	Short: "Convert any file to any format",
	Long:  `Convert any file to any format. For example:`,
	RunE: func(cmd *cobra.Command, args []string) error {
		model := tui.InitialModel()
		p := tea.NewProgram(model, tea.WithAltScreen())
		finalModel, err := p.Run()
		if err != nil {
			return fmt.Errorf("failed to run program: %w", err)
		}

		selectedFiles, targetFormat := finalModel.(tui.Model).GetSelectedFiles()
		if len(selectedFiles) == 0 {
			fmt.Println("No files selected.")
		}
		if targetFormat == nil {
			fmt.Println("No format selected.")
			return nil
		}

		fmt.Printf("Converting %d files to %s\n", len(selectedFiles), *targetFormat)

		vips.LoggingSettings(nil, vips.LogLevelError)
		vips.Startup(nil)
		defer vips.Shutdown()

		if len(selectedFiles) > 1 {
			for _, file := range selectedFiles {
				if err := convert(file, targetFormat, true); err != nil {
					fmt.Printf("Error converting file %s: %v\n", file.Path, err)
				}
			}
			fmt.Println("Conversion completed.")
			return nil
		}

		if err := convert(selectedFiles[0], targetFormat, false); err != nil {
			fmt.Printf("Error converting file %s: %v\n", selectedFiles[0].Path, err)
			return err
		}

		fmt.Println("Conversion completed.")
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func convert(selectedFile tui.File, targetFormat *types.ImageRasterFormat, isBulk bool) error {
	img, err := vips.NewImageFromFile(selectedFile.Path)
	if err != nil {
		panic(err)
	}

	if img.Format().FileExt()[1:] == toFormat && !(strings.HasSuffix(selectedFile.Name, "dng") && toFormat == "tiff") {
		fmt.Println("already in the desired format")
		return nil
	}
	switch *targetFormat {
	case types.JPEG:
		fallthrough
	case types.JPG:
		fmt.Println("converting to jpeg")
		images.ToJPEG(img, selectedFile.Path, "", isBulk)
	case types.PNG:
		fmt.Println("converting to png")
		images.ToPNG(img, selectedFile.Path, "", isBulk)
	case types.GIF:
		fmt.Println("converting to gif")
		images.ToGIF(img, selectedFile.Path, "", isBulk)
	case types.TIFF:
		fmt.Println("converting to tiff")
		images.ToTIFF(img, selectedFile.Path, "", isBulk)
	case types.BMP:
		fmt.Println("converting to bmp")
		images.ToBMP(img, selectedFile.Path, "", isBulk)
	case types.WEBP:
		fmt.Println("converting to webp")
		images.ToWEBP(img, selectedFile.Path, "", isBulk)
	case types.HEIF:
		fmt.Println("converting to heif")
		images.ToHEIF(img, selectedFile.Path, "", isBulk)

	default:
		fmt.Println("unsupported format")
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.convertsh.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
