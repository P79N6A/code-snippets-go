package main

import "fmt"

func main() {
	// for 的最后一项后面必须同行跟 {
	for i := 5; i > 0; i-- {
		fmt.Println("hello")
	}
	i := 0
	// for 没有最后一项时可以不同行跟 {
	for
	{
		fmt.Printf("i1: %d\n", i)
		i++
		if i > 5 {
			break
		}
	}
	i = 0
	// 由于自动插入行末的 ; 所以可以写成下面的
	for
	;
	i < 6
	i++ {
		fmt.Println("i2: ", i)
	}
}

// https://gfw.go101.org/article/line-break-rules.html
