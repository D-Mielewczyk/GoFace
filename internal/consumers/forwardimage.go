package consumers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg" // or "image/png" for PNG format
	"log"
	"os"

	"github.com/vee2xx/camtron"
)

var tempImgDir string = "images"
var targetUrls []string

func SaveImage(streamChan chan []byte) {
	if _, err := os.Stat(tempImgDir); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(tempImgDir, os.ModePerm)
		}
	}

	for {
		select {
		// Receive a single frame (or byte slice) from the camera
		case frame, ok := <-streamChan:
			if !ok {
				log.Print("WARNING: Failed to get frame")
				return
			}

			// Assuming frame is a byte slice representing an image
			// Convert it to an image object
			img, _, err := image.Decode(bytes.NewReader(frame))
			if err != nil {
				log.Printf("Failed to decode frame: %v", err)
				continue
			}

			// Save the image to a file
			fileName := tempImgDir + "/captured_image.jpg" // or ".png"
			file, err := os.Create(fileName)
			if err != nil {
				log.Printf("Failed to create image file: %v", err)
				continue
			}
			defer file.Close()

			// Encode the image to JPEG (or PNG)
			err = jpeg.Encode(file, img, nil) // or png.Encode for PNG format
			if err != nil {
				log.Printf("Failed to encode image: %v", err)
			}
			return // We only capture and save one image
		}
	}
}

func StartImageCaptureConsumer() {
	file, _ := os.Open("consumers/conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	targetUrls = configuration.TargetUrls

	fmt.Println()
	// Create a channel to receive the image data
	imgStream := make(chan []byte, 1) // Buffer of 1 since we only need one frame
	camtron.RegisterStream(imgStream)

	go SaveImage(imgStream)
}

type Configuration struct {
	TargetUrls []string
}
