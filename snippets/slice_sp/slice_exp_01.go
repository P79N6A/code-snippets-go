package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var myPic [][]uint8
	for i := 0; i < dx; i++ {
		myPic = append(myPic, make([]uint8, dx))
		for j := 0; j < dy; j++ {
			// dot := (i + j) / 2
			// dot := i * j
			dot := i ^ j
			myPic[i][j] = uint8(dot)
		}
	}
	return myPic
}

func main() {
	pic.Show(Pic)
}
