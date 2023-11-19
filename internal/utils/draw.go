package utils

import (
	"image"
	"image/color"
)

func DrawRectangle(img *image.RGBA, x, y, width, height int) {
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

func DrawCircle(img *image.RGBA, centerX, centerY, radius int) {
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
