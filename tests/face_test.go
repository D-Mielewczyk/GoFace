package tests

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"github.com/D-Mielewczyk/GoFace/internal/detection"
	"github.com/D-Mielewczyk/GoFace/internal/utils"
)

func TestConvertPath(t *testing.T) {
	testCases := []struct {
		name         string
		inputPath    string
		outputDir    string
		expectedPath string
	}{
		{"Test with JPG file", "images/photo.jpg", "output", filepath.Join("output", "photo.jpg")},
		{"Test with PNG file", "images/photo.png", "output", filepath.Join("output", "photo.png")},
		{"Test with nested directory", "images/nested/photo.jpg", "output", filepath.Join("output", "photo.jpg")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotPath := utils.ConvertPath(tc.inputPath, tc.outputDir)
			if gotPath != tc.expectedPath {
				t.Errorf("ConvertPath(%s, %s) = %s; want %s", tc.inputPath, tc.outputDir, gotPath, tc.expectedPath)
			}
		})
	}
}

func TestDetectFace(t *testing.T) {
	// Setup
	imagePath := "data/test.jpg"
	cascadePath := "../cascade/facefinder"
	circle := false

	// Run the function
	detection.DetectFace(imagePath, "output", circle, cascadePath)

	// Load the output image
	outputImg, err := loadImage("output/test.jpg")
	if err != nil {
		t.Fatalf("Failed to load output image: %v", err)
	}

	// Load the reference image
	referenceImg, err := loadImage("data/test_result.jpg")
	if err != nil {
		t.Fatalf("Failed to load reference image: %v", err)
	}

	// Compare images
	if !imagesEqual(outputImg, referenceImg) {
		t.Errorf("Processed image does not match the reference image")
	}
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

func imagesEqual(img1, img2 image.Image) bool {
	// First, check if both images have the same dimensions
	if img1.Bounds() != img2.Bounds() {
		return false
	}

	// Iterate over each pixel in the images
	for y := img1.Bounds().Min.Y; y < img1.Bounds().Max.Y; y++ {
		for x := img1.Bounds().Min.X; x < img1.Bounds().Max.X; x++ {
			if img1.At(x, y) != img2.At(x, y) {
				return false
			}
		}
	}

	return true
}
