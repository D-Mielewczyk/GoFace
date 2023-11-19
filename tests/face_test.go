package tests

import (
    "image"
    "os"
    "testing"
	"path/filepath"

    "github.com/D-Mielewczyk/GoFace/internal/detection"
)

func TestConvertPath(t *testing.T) {
    testCases := []struct {
        name         string
        inputPath    string
        expectedPath string
    }{
        {"Test with JPG file", "images/photo.jpg", filepath.Join("output", "photo.jpg")},
        {"Test with PNG file", "images/photo.png", filepath.Join("output", "photo.png")},
        {"Test with nested directory", "images/nested/photo.jpg", filepath.Join("output", "photo.jpg")},
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            gotPath := detection.ConvertPath(tc.inputPath)
            if gotPath != tc.expectedPath {
                t.Errorf("convertPath(%s) = %s; want %s", tc.inputPath, gotPath, tc.expectedPath)
            }
        })
    }
}

func TestDetectFace(t *testing.T) {
    // Setup
    imagePath := "data/test.jpg"
    cascadePath := "../cascade/facefinder"
    circle := false // or true, depending on your test case

    // Run the function
    detection.DetectFace(imagePath, cascadePath, circle)

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
