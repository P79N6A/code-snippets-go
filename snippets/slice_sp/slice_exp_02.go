package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	myPic := make([][]uint8, dy)

	for i := range myPic {
		myPic[i] = make([]uint8, dx)
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

// https://tour.golang.org/moretypes/18
