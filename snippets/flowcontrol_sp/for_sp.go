package main

import (
	"fmt"
	"math"
)

func basicForTest() {
	fmt.Println("=== basicForTest ===")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 { // for ; sum < 1000; {，分号其实被省略了
		sum += sum
	}
	fmt.Println(sum)

	i := 0
	// 无限循环
	for {
		i++
		if i > 10 {
			break
		}
	}
	fmt.Println("DONE")
}

func sqrt(x float64) float64 {
	z := x / 2
	last := 0.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		if math.Abs(last-z) < 1e-10 {
			// 差距过小，提前退出
			break
		}
		// test
		// _result := math.Sqrt(x)
		// fmt.Println(i, "|", z, "-", _result, "=", _result-z, "|", last-z)
		last = z
	}
	return z
}

func sqrtTest() {
	fmt.Println("=== sqrtTest ===")
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(20))
	fmt.Println(sqrt(500))
	fmt.Println(sqrt(715))
}

func main() {
	basicForTest()
	sqrtTest()
}
