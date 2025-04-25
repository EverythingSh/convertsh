package images

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/EverythingSh/convertsh/assets"
	"github.com/EverythingSh/convertsh/pkg/types"
	"github.com/davidbyttow/govips/v2/vips"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

type ImageStat struct {
	Width  int
	Height int
	Size   float64
}

// TestAllImageConversions tests conversion to all supported formats
func TestAllImageConversions(t *testing.T) {
	vips.LoggingSettings(nil, vips.LogLevelError)
	vips.Startup(nil)
	defer vips.Shutdown()

	formats := []struct {
		format     types.ImageRasterFormat
		extension  string
		skipSuffix string // Skip files that already have this suffix
		converter  func(*vips.ImageRef, string, string, bool)
	}{
		{types.PNG, ".png", "png", ToPNG},
		{types.JPEG, ".jpg", "jpg", ToJPEG},
		{types.GIF, ".gif", "gif", ToGIF},
		{types.TIFF, ".tiff", "tiff", ToTIFF},
		{types.BMP, ".bmp", "bmp", ToBMP},
		{types.WEBP, ".webp", "webp", ToWEBP},
		{types.HEIF, ".heif", "heif", ToHEIF},
	}

	testImages, err := assets.GetAvailableImages()
	if err != nil {
		t.Fatalf("failed to get test images: %v", err)
	}

	outputBaseDir := "test_output"
	if err := os.MkdirAll(outputBaseDir, 0755); err != nil {
		t.Fatalf("failed to create output directory: %v", err)
	}

	for _, format := range formats {
		t.Run(fmt.Sprintf("To%s", strings.ToUpper(string(format.format))), func(t *testing.T) {
			// Create format-specific output directory
			outputDir := filepath.Join(outputBaseDir, string(format.format))
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				t.Fatalf("failed to create output directory for %s: %v", format.format, err)
			}

			for _, imgPath := range testImages {
				// Skip if the input file is already in target format
				if strings.HasSuffix(strings.ToLower(imgPath), format.skipSuffix) {
					continue
				}

				img, err := vips.NewImageFromFile(imgPath)
				if err != nil {
					t.Logf("skipping image %s: %v", imgPath, err)
					continue
				}

				origStat, err := imageStat(imgPath)
				if err != nil {
					t.Errorf("failed to get image stat for %s: %v", imgPath, err)
					img.Close()
					continue
				}

				// Create output filename
				baseName := filepath.Base(imgPath)
				ext := filepath.Ext(baseName)
				baseNameNoExt := baseName[:len(baseName)-len(ext)]
				outputFilename := filepath.Join(outputDir, baseNameNoExt+format.extension)

				// Convert the image
				format.converter(img, imgPath, outputFilename, true)

				// Verify the output exists
				_, err = os.Stat(outputFilename)
				if err != nil {
					t.Errorf("output file not found for %s to %s conversion: %v",
						imgPath, format.format, err)
					img.Close()
					continue
				}

				// Get output image stats
				outStat, err := getImageStats(outputFilename)
				if err != nil {
					t.Errorf("failed to get stats for output file %s: %v", outputFilename, err)
					img.Close()
					continue
				}

				t.Logf(
					"Converted %s → %s \n Resolution: %dx%d → %dx%d \n Size: %.3f MB → %.3f MB\n",
					imgPath, outputFilename,
					origStat.Width, origStat.Height, outStat.Width, outStat.Height,
					origStat.Size, outStat.Size,
				)

				img.Close()
			}
		})
	}
}

// getImageStats gets image dimensions and file size
func getImageStats(imgPath string) (ImageStat, error) {
	// For file size
	stat, err := os.Stat(imgPath)
	if err != nil {
		return ImageStat{}, fmt.Errorf("failed to get file stat: %v", err)
	}

	// Try to get dimensions using vips first
	img, err := vips.NewImageFromFile(imgPath)
	if err == nil {
		defer img.Close()
		return ImageStat{
			Width:  img.Width(),
			Height: img.Height(),
			Size:   float64(stat.Size()) / (1024 * 1024),
		}, nil
	}

	// Fall back to Go's image package if vips fails
	file, err := os.Open(imgPath)
	if err != nil {
		return ImageStat{}, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return ImageStat{}, fmt.Errorf("failed to decode image: %v", err)
	}

	return ImageStat{
		Width:  config.Width,
		Height: config.Height,
		Size:   float64(stat.Size()) / (1024 * 1024),
	}, nil
}

// Keeping the legacy imageStat function for backward compatibility
func imageStat(imgPath string) (ImageStat, error) {
	img, err := vips.NewImageFromFile(imgPath)
	if err != nil {
		return ImageStat{}, fmt.Errorf("failed to load image %s: %v", imgPath, err)
	}
	defer img.Close()

	stat, err := os.Stat(imgPath)
	if err != nil {
		return ImageStat{}, fmt.Errorf("failed to get file stat for %s: %v", imgPath, err)
	}

	return ImageStat{
		Width:  img.Width(),
		Height: img.Height(),
		Size:   float64(stat.Size()) / (1024 * 1024),
	}, nil
}
