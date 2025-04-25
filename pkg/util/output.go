package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var timeNow string

func init() {
	timeNow = time.Now().Format("20060102_150405")
}

func CreateOutputPath(inputPath, ext string, isBulk bool) string {
	baseName := filepath.Base(inputPath)
	inputExt := filepath.Ext(baseName)
	baseNameNoExt := baseName[:len(baseName)-len(inputExt)]

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}

	baseNameWithSuffix := baseNameNoExt + "_con"

	var outputPath string
	if isBulk {
		outputDir := filepath.Join(currentDir, "converted_"+ext[1:]+"_"+timeNow)
		if err := os.MkdirAll(outputDir, 0755); err == nil {
			outputPath = filepath.Join(outputDir, baseNameWithSuffix+ext)
		} else {
			outputPath = filepath.Join(currentDir, baseNameWithSuffix+ext)
		}
	} else {
		outputPath = filepath.Join(currentDir, baseNameWithSuffix+ext)
	}

	return ensureUniqueFilePath(outputPath)
}

func ensureUniqueFilePath(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path
	}

	dir := filepath.Dir(path)
	ext := filepath.Ext(path)
	baseName := filepath.Base(path)
	baseNameNoExt := baseName[:len(baseName)-len(ext)]

	counter := 1
	for {
		newPath := filepath.Join(dir, fmt.Sprintf("%s_%d%s", baseNameNoExt, counter, ext))
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
		counter++
	}
}
