package main

import "fmt"

func main() {
	// 基础算术操作
	var (
		a, b float32 = 12.0, 3.14
		c, d int16   = 15, -6
		e    uint8   = 7
	)
	fmt.Println('A'+12, 12-a, a*b, c%d)
	fmt.Println(c+int16(e), uint8(c)+e)
	fmt.Println(a/b, c/d, -100/-9, -100/-9.0) // int/int 得到的是 int
	fmt.Println(10 % 3)
	//fmt.Println(10 % 3.1) // illegal constant expression: floating-point % operation

	// 位操作
	fmt.Println("=== bit operator ===")
	fmt.Println(1|1, 1|0, 0|1, 0|0)
	fmt.Println(1&1, 1&0, 0&1, 0&0)
	fmt.Println(1^1, 1^0, 0^1, 0^0)
	fmt.Println(^1, ^0)                 // 位反（位补）
	fmt.Println(1&^1, 1&^0, 0&^1, 0&^0) // 清位，1 的时候就把对应的那一位清掉，0 就不变，相当于 a & (^b)
	//fmt.Println(1.1^1.1) // internal compiler error: binaryOp: bad operation: 1.1 ^ 1.1

	fmt.Println(1>>1, 1<<1, -1>>1, -1>>2) // -1 怎么右移都是 -1
	u := uint8(0xff)
	fmt.Println(u, u>>2, u>>4, u>>7) // unsigned 类型右移最高位的 1 不会再补
	//fmt.Println(3>>-1) // invalid negative shift count: -1

	f1, g1 := 1, uint8(32)
	//h1 := f1 << g1 // invalid operation: f << g (shift count type int, must be unsigned integer)
	h1 := f1 << g1
	fmt.Println(f1, g1, h1)

	f2, g2 := int32(1), uint8(32)
	h2 := f2 << g2         // 0，因为 f2 是 int32
	var h3 int32 = 1 << g2 // 0，因为结果变量是 int32
	h4 := int64(f2 << g2)  // 0
	h5 := int64(f2) << g2  // 4294967296
	h6 := 1 << g2          // 4294967296
	fmt.Println(f2, g2, h2, h3, h4, h5, h6)
}
