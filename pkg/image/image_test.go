package images

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/EverythingSh/convertsh/assets"
	"github.com/davidbyttow/govips/v2/vips"
)

func TestToBMPConversion(t *testing.T) {
	vips.LoggingSettings(nil, vips.LogLevelError)
	vips.Startup(nil)
	defer vips.Shutdown()

	testImages, err := assets.GetAvailableImages()
	if err != nil {
		t.Fatalf("failed to get test images: %v", err)
	}

	// Create output directory if it doesn't exist
	if err := os.MkdirAll("output_bmp", 0755); err != nil {
		t.Errorf("failed to create output directory: %v", err)
	}

	for _, imgPath := range testImages {
		img, err := vips.NewImageFromFile(imgPath)
		if err != nil {
			t.Errorf("failed to load image %s: %v", imgPath, err)
			continue
		}

		origStat, _ := os.Stat(imgPath)
		origWidth := img.Width()
		origHeight := img.Height()
		origSize := float64(origStat.Size()) / (1024 * 1024)

		// Prepare output filename: output/basename+inputFormat.bmp
		baseName := filepath.Base(imgPath)
		ext := filepath.Ext(baseName)
		baseNameNoExt := baseName[:len(baseName)-len(ext)]
		inputFormat := img.Format().FileExt()[1:]
		outputFilename := filepath.Join("output_bmp", baseNameNoExt+"TO"+inputFormat+".bmp")

		// Call ToBMP with the correct output path logic
		ToBMP(img, imgPath, outputFilename)

		outStat, err := os.Stat(outputFilename)
		if err != nil {
			t.Errorf("output BMP file not found for %s", imgPath)
			img.Close()
			continue
		}

		outFile, err := os.Open(outputFilename)
		if err != nil {
			t.Errorf("failed to open output BMP: %v", err)
			img.Close()
			continue
		}
		bmpImg, _, err := image.DecodeConfig(outFile)
		outFile.Close()
		if err != nil {
			t.Errorf("failed to decode output BMP: %v", err)
			img.Close()
			continue
		}

		outputSize := float64(outStat.Size()) / (1024 * 1024)

		fmt.Printf(
			"Converted %s → %s | Resolution: %dx%d → %dx%d | Size: %.3f mb → %.3f mb\n",
			imgPath, outputFilename,
			origWidth, origHeight, bmpImg.Width, bmpImg.Height,
			origSize, outputSize,
		)

		img.Close()
	}
}
