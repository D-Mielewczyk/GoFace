package main

import (
    "github.com/D-Mielewczyk/GoFace/internal/detection"
    "log"
)

func main() {
    log.Println("main")
    detection.DetectFace("images/ja.jpg", "")
}