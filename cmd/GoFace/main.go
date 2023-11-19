package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/D-Mielewczyk/GoFace/internal/detection"
	"github.com/D-Mielewczyk/GoFace/internal/utils"
)

func main() {
	log.Println("Starting...")

	// Define flags
	var outputDir string
	var drawCircle bool
	var showHelp bool
	var measurePerformance bool

	flag.StringVar(&outputDir, "output", "output", "Specify the output directory.")
	flag.StringVar(&outputDir, "o", "output", "(shorthand).")
	flag.BoolVar(&drawCircle, "circle", false, "Draw circles around faces instead of rectangles.")
	flag.BoolVar(&drawCircle, "c", false, "(shorthand).")
	flag.BoolVar(&showHelp, "help", false, "Show help message.")
	flag.BoolVar(&showHelp, "h", false, "(shorthand).")
	flag.BoolVar(&measurePerformance, "performance", false, "Measure the performance of image processing.")
	flag.BoolVar(&measurePerformance, "p", false, "(shorthand).")

	flag.Usage = func() {
		log.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		log.Println("Provide an image file name or 'all' to process all images in the 'images' directory.")
	}

	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	arg := args[0]

	processImage := func(imagePath string) {
		startTime := time.Now()

		log.Printf("Processing file: %s", imagePath)
		detection.DetectFace(imagePath, outputDir, drawCircle, "")

		if measurePerformance {
			duration := time.Since(startTime)
			log.Printf("Processing time for %s: %v", imagePath, duration)
		}
	}

	files, err := os.ReadDir("images")
	if err != nil {
		log.Fatalf("Cannot read images directory: %v", err)
	}

	if arg != "all" {
		for _, file := range files {
			if file.Name() != arg {
				files = utils.RemoveFromSlice(files, file)
			}
		}

		if len(files) == 0 {
			log.Fatalf("File does not exist: %v", filepath.Join("images", arg))
		}
	}

	for _, file := range files {
		if !file.IsDir() {
			imagePath := filepath.Join("images", file.Name())
			processImage(imagePath)
		}
	}
}
