package main

import (
    "github.com/D-Mielewczyk/GoFace/internal/detection"
    "log"
    "os"
    "path/filepath"
)

func main() {
    log.Println("main")

    if len(os.Args) < 2 {
        log.Fatalf("Usage: %s <image path> or %s all", os.Args[0], os.Args[0])
    }
    arg := os.Args[1]

    if arg == "all" {
        files, err := os.ReadDir("images")
        if err != nil {
            log.Fatalf("Cannot read images directory: %v", err)
        }

        for _, file := range files {
            if !file.IsDir() {
                imagePath := filepath.Join("images", file.Name())
                log.Printf("Processing file: %s", imagePath)
                detection.DetectFace(imagePath, "")
            }
        }
    } else {
        imagePath := filepath.Join("images", arg)
        detection.DetectFace(imagePath, "")
    }
}
