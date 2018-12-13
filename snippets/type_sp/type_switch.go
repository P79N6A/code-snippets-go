package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	// case int, float64: // 多个 type 的时候类型不会确定下来，只会用开始时候的类型
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// a := 12
	// 必须在 type switch 中使用这种语法
	// t := a.(type) // use of .(type) outside type switch

	do(21)
	do(1.414)
	do("hello")
	do(true)
}

// https://stackoverflow.com/questions/40575033/golang-multiple-case-in-type-switch
