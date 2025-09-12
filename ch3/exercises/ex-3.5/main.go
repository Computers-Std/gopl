package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x80, 0x00, 0x00, 0xff},
	color.RGBA{0xFF, 0x00, 0x00, 0xff},
	color.RGBA{0xFF, 0x7F, 0x00, 0xff},
	color.RGBA{0xFF, 0xFF, 0x00, 0xff},
	color.RGBA{0x00, 0xFF, 0x00, 0xff},
	color.RGBA{0x00, 0xFF, 0xFF, 0xff},
	color.RGBA{0x00, 0x00, 0xFF, 0xff},
	color.RGBA{0x4B, 0x00, 0x82, 0xff},
	color.RGBA{0x8A, 0x2B, 0xE2, 0xff},
	color.RGBA{0xFF, 0x00, 0xFF, 0xff},
	color.RGBA{0xC7, 0x15, 0x85, 0xff},
	color.RGBA{0xFF, 0xD7, 0x00, 0xff},
	color.RGBA{0xBD, 0xB7, 0x6B, 0xff},
	color.RGBA{0x8B, 0x45, 0x13, 0xff},
	color.RGBA{0xD3, 0xD3, 0xD3, 0xff},
	color.RGBA{0xFF, 0xFF, 0xFF, 0xff},
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // Note: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n%16]
		}
	}
	return color.Black
}
