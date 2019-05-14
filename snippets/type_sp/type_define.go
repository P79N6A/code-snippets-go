package main

import "fmt"

func main() {
	type status bool // status 和 bool 不是同一种类型
	type boolean = bool // boolean 和 bool 是同一种类型
	var s status = true

	fmt.Println(s)
	var b boolean = false // 因为是同一种类型，所以这里 IDE 会提示 boolean 类型生命可以删掉
	fmt.Println(b)
}

// https://gfw.go101.org/article/basic-types-and-value-literals.html
