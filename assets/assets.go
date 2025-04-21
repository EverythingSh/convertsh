package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetAvailableImages returns a slice of image file paths in the assets directory, prefixed with the current working directory.
func GetAvailableImages() ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %w", err)
	}

	cwd = strings.Replace(cwd, "/pkg/image", "", 1)

	assetsDir := filepath.Join(cwd, "assets")
	files, err := os.ReadDir(assetsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read assets directory: %w", err)
	}

	var images []string
	for _, file := range files {
		if !(file.IsDir() || strings.HasSuffix(file.Name(), ".go")) {
			images = append(images, filepath.Join(assetsDir, file.Name()))
		}
	}
	return images, nil
}
