// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/cmplx"
	"os"
)

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
			// img.Set(px, py, hsvColor(uint8(float64(px)/float64(width)*255.)))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func hsvColor(i uint8) color.Color {
	// see https://en.wikipedia.org/wiki/HSL_and_HSV#From_HSL

	// invert (red on the outside, blue as you move in. comment out)
	i = 255 - i
	const c = 1
	h := float64(i) / 255.0 * 240.
	h2 := h / 60.
	x := 1 - math.Abs(math.Mod(h2, 2)-1)
	f2u8 := func(f float64) uint8 {
		return uint8(f * 255.)
	}
	rgba := func(r, g, b float64) color.RGBA {
		return color.RGBA{f2u8(r), f2u8(g), f2u8(b), 255}
	}
	switch {
	case h2 <= 1.:
		return rgba(c, x, 0)
	case h2 <= 2.:
		return rgba(x, c, 0)
	case h2 <= 3.:
		return rgba(0, c, x)
	case h2 <= 4.:
		return rgba(0, x, c)
	case h2 <= 5.:
		return rgba(x, 0, c)
	case h2 <= 6.:
		return rgba(c, 0, x)
	default:
		log.Fatalf("Unexpected value for h2: %f", h2)
	}
	return color.RGBA{0, 0, 0, 255}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return hsvColor(255 - contrast*n)
		}
	}
	return hsvColor(0)
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
