package main

import (
    "fmt"
    "log"
    "os"
    _ "image/jpeg"

    pigo "github.com/esimov/pigo/core"
)

func main() {
    fmt.Println("Starting...")
    cascadeFile, err := os.ReadFile("cascade/facefinder")
    if err != nil {
        log.Fatalf("Error reading the cascade file: %v", err)
    }
    fmt.Println("Classifier loaded.")

    src, err := pigo.GetImage("images/ja.jpg")
    if err != nil {
        log.Fatalf("Cannot open the image file: %v", err)
    }
    fmt.Println("Image loaded.")
    
    pixels := pigo.RgbToGrayscale(src)
    cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y
    
    cParams := pigo.CascadeParams{
        MinSize:     20,
        MaxSize:     1000,
        ShiftFactor: 0.1,
        ScaleFactor: 1.1,
    
        ImageParams: pigo.ImageParams{
            Pixels: pixels,
            Rows:   rows,
            Cols:   cols,
            Dim:    cols,
        },
    }
    
    pigo := pigo.NewPigo()
    // Unpack the binary file. This will return the number of cascade trees,
    // the tree depth, the threshold and the prediction from tree's leaf nodes.
    classifier, err := pigo.Unpack(cascadeFile)
    if err != nil {
        log.Fatalf("Error reading the cascade file: %s", err)
    }
    
    angle := 0.0 // cascade rotation angle. 0.0 is 0 radians and 1.0 is 2*pi radians
    
    // Run the classifier over the obtained leaf nodes and return the detection results.
    // The result contains quadruplets representing the row, column, scale and detection score.
    dets := classifier.RunCascade(cParams, angle)
    
    // Calculate the intersection over union (IoU) of two clusters.
    dets = classifier.ClusterDetections(dets, 0.2)
}