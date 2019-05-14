package main

import "fmt"

func main() {
	// 一些可以转换的

	fmt.Println(complex128(1 + -1e-1000i)) // 精度不够，被舍弃
	fmt.Println(float32(0.49999999)) // 同样被舍弃
	fmt.Println(float32(17000000000000000))
	fmt.Println(float32(123))
	fmt.Println(uint(1.0))

	fmt.Println(int16(6+0i))
	fmt.Println(complex128(789))

	fmt.Println(string(65))

	// 不可以转换的，问题：truncated、overflows、cannot convert
	//fmt.Println(uint(1.1)) // constant 1.1 truncated to integer
	//fmt.Println(uint8(-1)) // constant -1 overflows uint8
	//fmt.Println(float64(1+2i)) // constant (1+2i) truncated to real
	//fmt.Println(float64(-1e1000)) // constant -1e+1000 overflows float64
	//fmt.Println(int(0x10000000000000000)) // constant 18446744073709551616 overflows int
	//fmt.Println(string(65.0)) // cannot convert 65 (type untyped number) to type string
	//fmt.Println(string(66+0i)) // cannot convert 66 + 0 (type untyped number) to type string
}
