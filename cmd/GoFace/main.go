package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/D-Mielewczyk/GoFace/internal/detection"
)

func main() {
	log.Println("main")

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <image path> or %s all [circle]", os.Args[0], os.Args[0])
	}
	arg := os.Args[1]
	drawCircle := len(os.Args) > 2 && os.Args[2] == "circle"

	processImage := func(imagePath string) {
		log.Printf("Processing file: %s", imagePath)
		detection.DetectFace(imagePath, "", drawCircle)
	}

	if arg == "all" {
		files, err := os.ReadDir("images")
		if err != nil {
			log.Fatalf("Cannot read images directory: %v", err)
		}

		for _, file := range files {
			if !file.IsDir() {
				imagePath := filepath.Join("images", file.Name())
				processImage(imagePath)
			}
		}
	} else {
		imagePath := filepath.Join("images", arg)
		processImage(imagePath)
	}
}
