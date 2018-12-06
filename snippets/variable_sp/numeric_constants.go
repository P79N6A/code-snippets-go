package main

import "fmt"

// 数字常量高精度，会根据使用场景来决定类型
const (
	// Big，一个很大的数字
	Big = 1 << 100
	// Small，2
	Small = Big >> 99
)

func printTypeAndValue(i interface{}) {
	fmt.Printf("Type: %T Value: %v\n", i, i)
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func needFloat2(x float64) float64 {
	return x
}

func main() {
	// printTypeAndValue(Big) // constant 1267650600228229401496703205376 overflows int
	printTypeAndValue(Small) // Type: int Value: 2
	printTypeAndValue(needInt(Small))
	// printTypeAndValue(needInt(Big)) // constant 1267650600228229401496703205376 overflows int
	printTypeAndValue(needFloat(Big))
	printTypeAndValue(needFloat(Small))
	printTypeAndValue(needFloat2(Small)) // 即使里面没操作还是会变成 float
}
