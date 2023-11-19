package main

import (
	"fmt"
	"github.com/D-Mielewczyk/GoFace/internal/consumers"
	"github.com/vee2xx/camtron"
)

func main() {
	fmt.Println("Starting...")
	consumers.StartImageCaptureConsumer()
	fmt.Println("Working...")
	camtron.StartCam()
	fmt.Println("Ending...")
}
