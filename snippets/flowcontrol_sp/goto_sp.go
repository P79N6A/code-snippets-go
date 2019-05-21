package main

import "fmt"

func goto1() {
	fmt.Println("a")
	goto C
	fmt.Println("b")
C:
	fmt.Println("c")
}

func goto2() {
	i := 0
FOR:
	for ; i < 3; i++ {
		for j := i; j < 3; j++ {
			if i == 0 && j == 1 {
				i = i + 1
				goto FOR
			}
			fmt.Println(i, j)
		}
	}
}

func main() {
	goto1()
	goto2()
}
