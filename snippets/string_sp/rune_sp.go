package main

import "fmt"

func main() {
	var a rune = 97
	fmt.Println(a, 'a', '\141', '\x61', '\u0061', '\U00000061')
	fmt.Println('\u4f17', '众')
}
