package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

func makePic(f func(x, y int) uint8) func(dx, dy int) [][]uint8 {
	pic := func(dx, dy int) [][]uint8 {
		myPic := make([][]uint8, dy)

		for i := range myPic {
			myPic[i] = make([]uint8, dx)
			for j := 0; j < dy; j++ {
				myPic[i][j] = f(i, j)
			}
		}
		return myPic
	}
	return pic
}

func f1(x, y int) uint8 {
	return uint8(x ^ y)
}

func f2(x, y int) uint8 {
	return uint8(x + y)
}

// Image that as image.Image
type Image struct{}

// ColorModel of Image
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
