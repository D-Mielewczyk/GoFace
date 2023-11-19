package detection

import (
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	pigo "github.com/esimov/pigo/core"

	"github.com/D-Mielewczyk/GoFace/internal/utils"
)

func DetectFace(image_path, output_dir string, circle bool, cascade_path string) {
	// image_path must be provided hover you can ommit cascade_path for a default value
	log.Println("Starting...")

	if cascade_path == ""{
		cascade_path = "cascade/facefinder"
	}
	cascadeFile, err := os.ReadFile(cascade_path)
	if err != nil {
		log.Fatalf("Error reading the cascade file %v, beacouse of:\n%v", cascade_path, err)
	}
	log.Println("Classifier loaded.")

	src, err := pigo.GetImage(image_path)
	if err != nil {
		log.Fatalf("Cannot open the image file %v, beacouse of:\n%v", image_path, err)
	}
	log.Println("Image loaded.")

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

	inFile, err := os.Open(image_path)
	if err != nil {
		log.Fatalf("Cannot open the image file: %v", err)
	}
	defer inFile.Close()

	img, _, err := image.Decode(inFile)
	if err != nil {
		log.Fatalf("Error decoding image: %v", err)
	}

	// Create a new image for the output
	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, img.Bounds(), img, image.Point{}, draw.Src)

	// Draw rectangles around detected faces
	for _, det := range dets {
		if det.Q > 5 {
			if circle {
				utils.DrawCircle(dst, det.Col, det.Row, det.Scale/2)
			} else {
				utils.DrawRectangle(dst, det.Col-det.Scale/2, det.Row-det.Scale/2, det.Scale, det.Scale)
			}
		}
	}

	// Save the modified image
	output_path := utils.ConvertPath(image_path, output_dir)
	outFile, err := os.Create(output_path)
    if err != nil {
        log.Fatalf("Cannot create output file %v, because of:\n%v", output_path, err)
    }
    defer outFile.Close()

    err = jpeg.Encode(outFile, dst, nil)
    if err != nil {
        log.Fatalf("Cannot save the image %v, because of:\n%v", output_path, err)
    }

    log.Printf("Output image saved as %v", output_path)
}


