package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if 跟一个短语句
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// 设法促成
		// else 中也可以访问 if 后的短语句里声明的变量，但这里 lint 工具其实已经在提醒去掉这个 else 了
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func elseIf(a int) {
	if a < 0 {
		fmt.Println("(-∞, 0)")
	} else if a < 2 {
		fmt.Println("[0, 2)")
	} else {
		fmt.Println("[2, +∞)")
	}
}

func main() {
	fmt.Println(sqrt(2), sqrt(-3))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	elseIf(-5)
	elseIf(1)
	elseIf(10)
}
