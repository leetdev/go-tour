// https://tour.golang.org/methods/25

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	// (x+y)/2
	// x*y
	// x^y
	// x*x+y*y+x*y
	// 2*x*y
	// x^y + y^x

	return color.RGBA{uint8(x * y), uint8(x ^ y), uint8(x*x + y*y + x*y), uint8((x + y) / 2)}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
