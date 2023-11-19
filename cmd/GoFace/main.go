package main

import (
    "github.com/D-Mielewczyk/GoFace/internal/detection"
    "log"
    "os"
)

func main() {
    log.Println("main")

    if len(os.Args) < 2 {
        log.Fatalf("Usage: %s <image path>", os.Args[0])
    }
    imagePath := "images/" +os.Args[1]

    detection.DetectFace(imagePath, "")
}