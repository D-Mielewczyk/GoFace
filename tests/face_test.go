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
	test_cases := []struct {
		name         string
		input_path    string
		output_dir    string
		expected_path string
	}{
		{"Test with JPG file", "images/photo.jpg", "output", filepath.Join("output", "photo.jpg")},
		{"Test with PNG file", "images/photo.png", "output", filepath.Join("output", "photo.png")},
		{"Test with nested directory", "images/nested/photo.jpg", "output", filepath.Join("output", "photo.jpg")},
	}

	for _, tast_case := range test_cases {
		t.Run(tast_case.name, func(t *testing.T) {
			got_path := utils.ConvertPath(tast_case.input_path, tast_case.output_dir)
			if got_path != tast_case.expected_path {
				t.Errorf("ConvertPath(%s, %s) = %s; want %s", tast_case.input_path, tast_case.output_dir, got_path, tast_case.expected_path)
			}
		})
	}
}

func TestDetectFace(t *testing.T) {
	// Setup
	image_path := "data/test.jpg"
	cascade_path := "../cascade/facefinder"
	circle := false

	// Run the function
	detection.DetectFace(image_path, "output", circle, cascade_path)

	// Load the output image
	output_img, err := loadImage("output/test.jpg")
	if err != nil {
		t.Fatalf("Failed to load output image: %v", err)
	}

	// Load the reference image
	reference_img, err := loadImage("data/test_result.jpg")
	if err != nil {
		t.Fatalf("Failed to load reference image: %v", err)
	}

	// Compare images
	if !imagesEqual(output_img, reference_img) {
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
