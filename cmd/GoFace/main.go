package main

import (
    "flag"
    "log"
    "os"
    "path/filepath"

    "github.com/D-Mielewczyk/GoFace/internal/detection"
    "github.com/D-Mielewczyk/GoFace/internal/utils"
)

func main() {
    log.Println("main")

    // Define flags
    var outputDir string
    var drawCircle bool

    flag.StringVar(&outputDir, "output", "output", "Specify the output directory.")
    flag.StringVar(&outputDir, "o", "output", "Specify the output directory (shorthand).")
    flag.BoolVar(&drawCircle, "circle", false, "Draw circles around faces instead of rectangles.")
    flag.BoolVar(&drawCircle, "c", false, "Draw circles around faces instead of rectangles (shorthand).")
    flag.Parse()

    args := flag.Args()
    if len(args) < 1 {
        log.Fatalf("Usage: %s -o <output path> -c <image path> or %s all", os.Args[0], os.Args[0])
    }
    arg := args[0]

    processImage := func(imagePath string) {
        log.Printf("Processing file: %s", imagePath)
        detection.DetectFace(imagePath, outputDir, drawCircle, "")
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
