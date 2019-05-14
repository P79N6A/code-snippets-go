package main

import "fmt"

// 常量声明的 = 表示绑定，将字面值绑定到对应的有名常量上，有名常量其实表示一个字面常量

// Pi the Pi
const Pi = 3.14
const π = Pi

// 声明一组
const (
	Yes        = true
	No         = !Yes
	MaxDegrees = 360
	Unit       = "弧度"
)

// 可以在声明的时候指定确定类型，这样声明的常量称为类型确定有名常量
const X1 float32 = 3.14
const (
	A1, B1 int64   = -3, 5
	Y1     float32 = 2.718
)
const X2 = float32(3.14)
const (
	A2, B2 = int64(-3), int64(5)
	Y2     = float32(2.718)
)

//const a uint8 = 256 // constant 256 overflows uint8

// 常量声明的自动补全
const (
	X3 float32 = 3.14
	Y3
	Z3

	A3, B3 = "Go", "language"
	C3, _
)

// iota 首次出现为 0，之后每次出现都 +1，用在上面的自动补全的场景，会有它的作用域
const (
	Failed      = iota - 1   // -1
	Unknown                  // 0
	Succeeded                // 1
	Readable1   = iota + 100 // 103
	Writable1                // 104
	Executable1              // 105
)

const (
	Readable2   = iota + 100 // 100
	Writable2                // 101
	Executable2              // 102
)

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Truth = false // cannot assign to Truth
	const Truth2 bool = false
	fmt.Println("Python rules?", Truth2)

	const DoublePi, HalfPi, Unit2 = π * 2, π * 0.5, "度"
	fmt.Println(Pi, π, DoublePi, HalfPi, Unit2, MaxDegrees, Unit, Yes, No)

	// const Wrold2 := "world" // syntax error: unexpected := at end of statement

	fmt.Println(X3, Y3, Z3, A3, B3, C3)

	fmt.Println(Failed, Unknown, Succeeded)
	fmt.Println(Readable1, Writable1, Executable1)
	fmt.Println(Readable2, Writable2, Executable2)
}
