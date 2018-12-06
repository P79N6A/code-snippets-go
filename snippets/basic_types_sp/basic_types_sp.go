package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// ToBe: to be or not to be
var (
	ToBe          = false // bool
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i) // complex128
)

func printTypeAndValue(i interface{}) {
	fmt.Printf("Type: %T Value: %v\n", i, i)
}

func typeAndValues() {
	fmt.Println("=== typeAndValues ===")
	printTypeAndValue(ToBe)
	printTypeAndValue(MaxInt)
	printTypeAndValue(z)
}

func typeConversions() {
	fmt.Println("=== typeConversions ===")
	x, y := 3, 4
	// var f float64 = math.Sqrt(x*x + y*y) // cannot use x * x + y * y (type int) as type float64 in argument to math.Sqrt
	f := math.Sqrt(float64(x*x + y*y))
	// var z uint = f // cannot use f (type float64) as type uint in assignment
	z := uint(f)
	fmt.Println(x, y, f, z)
}

func typeInference() {
	fmt.Println("=== typeInference ===")
	var i int
	j := i
	k := 42
	l := 3.142      // float64
	m := 9.8 + 0.5i // complex128
	printTypeAndValue(i)
	printTypeAndValue(j)
	printTypeAndValue(k)
	printTypeAndValue(l)
	printTypeAndValue(m)
}

func main() {
	typeAndValues()
	typeConversions()
	typeInference()
}

// https://tour.golang.org/basics/11
// [should have comment or be unexported · Issue #191 · golang/lint](https://github.com/golang/lint/issues/191)
// https://stackoverflow.com/questions/40145569/how-do-you-make-a-function-accept-multiple-types-in-go
