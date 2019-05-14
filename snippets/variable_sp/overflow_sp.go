package main

import "fmt"

const n = 1 << 64
const r = 'a' + 0x7FFFFFFF
const x = 2e+308

func main() {
	//fmt.Println(n, r, x) // 会报一堆 overflow，因为超出了默认的范围
	fmt.Println(n>>2, r-0x7FFFFFFF, x/2) // 计算好的结果没有溢出就不会报错

	var y1 uint8 = 128
	fmt.Println(y1)
	y1 = y1 + y1 + 10
	fmt.Println(y1) // 没有报错，溢出之后直接从 0 开始计数

	//var y2 = uint8(256) // constant 256 overflows uint8
	t := 256
	var y2 = uint8(t)
	fmt.Println(y2) // 没有报错，溢出之后直接从 0 开始计数
}
