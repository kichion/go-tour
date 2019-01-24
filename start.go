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

func (img Image) ColorModel() image.ColorModel {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8(x^y)
	return color.RGBA{v, v, 128, 189}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
