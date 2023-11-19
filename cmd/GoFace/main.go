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
	var output_dir string
	var draw_circle bool
	var show_help bool
	var measure_performance bool

	flag.StringVar(&output_dir, "output", "output", "Specify the output directory.")
	flag.StringVar(&output_dir, "o", "output", "(shorthand).")
	flag.BoolVar(&draw_circle, "circle", false, "Draw circles around faces instead of rectangles.")
	flag.BoolVar(&draw_circle, "c", false, "(shorthand).")
	flag.BoolVar(&show_help, "help", false, "Show help message.")
	flag.BoolVar(&show_help, "h", false, "(shorthand).")
	flag.BoolVar(&measure_performance, "performance", false, "Measure the performance of image processing.")
	flag.BoolVar(&measure_performance, "p", false, "(shorthand).")

	flag.Usage = func() {
		log.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		log.Println("Provide an image file name or 'all' to process all images in the 'images' directory.")
	}

	flag.Parse()

	if show_help {
		flag.Usage()
		return
	}

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	arg := args[0]

	processImage := func(image_path string) {
		start_time := time.Now()

		log.Printf("Processing file: %s", image_path)
		detection.DetectFace(image_path, output_dir, draw_circle, "")

		if measure_performance {
			duration := time.Since(start_time)
			log.Printf("Processing time for %s: %v", image_path, duration)
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
			image_path := filepath.Join("images", file.Name())
			processImage(image_path)
		}
	}
}
