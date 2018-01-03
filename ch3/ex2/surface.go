package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var f = fSimpleSin

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok := corner(i+1, j, f)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j, f)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1, f)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1, f)
			if !ok {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int, f func(x, y float64) float64) (float64, float64, bool) {
	// Find point (x, y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func fSin(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}

func fCos(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Cos(r) / r
}

func fASin(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Asinh(r) / r
}

func fHypot(x, y float64) float64 {
	r := math.Hypot(x, y)
	return r / r
}

func fExp(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Exp(r) / r
}

func fGamma(x, y float64) float64 {
	return math.Gamma(x) / math.Gamma(y)
}

func fBessel(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.J0(r)
}

func fBessel_1(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.J1(r)
}

func fLog(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Log(r)
}

func fSimpleSin(x, y float64) float64 {
	return (math.Sin(x) * math.Cos(y*0.5)) * 0.2
}
