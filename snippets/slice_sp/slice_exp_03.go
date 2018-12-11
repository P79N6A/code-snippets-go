package main

import (
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

func main() {
	pic.Show(makePic(f1))
}

// https://tour.golang.org/moretypes/18
// An answer of the exercise: Slices on a tour of Go https://gist.github.com/tetsuok/2280162
