package detection

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	pigo "github.com/esimov/pigo/core"
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
				drawCircle(dst, det.Col, det.Row, det.Scale/2)
			} else {
				drawRectangle(dst, det.Col-det.Scale/2, det.Row-det.Scale/2, det.Scale, det.Scale)
			}
		}
	}

	// Save the modified image
	output_path := ConvertPath(image_path, output_dir)
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

func drawRectangle(img *image.RGBA, x, y, width, height int) {
	red := color.RGBA{255, 0, 0, 255}

	// Draw horizontal lines
	for i := x; i < x+width; i++ {
		img.Set(i, y, red)
		img.Set(i, y+height, red)
	}

	// Draw vertical lines
	for i := y; i < y+height; i++ {
		img.Set(x, i, red)
		img.Set(x+width, i, red)
	}
}

func drawCircle(img *image.RGBA, centerX, centerY, radius int) {
	color := color.RGBA{255, 0, 0, 255}

	putPixel := func(x, y int) {
		img.Set(centerX+x, centerY+y, color)
		img.Set(centerX-x, centerY+y, color)
		img.Set(centerX+x, centerY-y, color)
		img.Set(centerX-x, centerY-y, color)
		img.Set(centerX+y, centerY+x, color)
		img.Set(centerX-y, centerY+x, color)
		img.Set(centerX+y, centerY-x, color)
		img.Set(centerX-y, centerY-x, color)
	}

	x, y, d := 0, radius, 3-2*radius
	putPixel(x, y)

	for y >= x {
		x++
		if d > 0 {
			y--
			d = d + 4*(x-y) + 10
		} else {
			d = d + 4*x + 6
		}
		putPixel(x, y)
	}
}

func ConvertPath(input_path string, output_dir string) string {
	fileName := filepath.Base(input_path)
	return filepath.Join(output_dir, fileName)
}
