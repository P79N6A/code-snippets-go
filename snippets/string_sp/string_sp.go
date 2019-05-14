package main

import "fmt"

func main() {
	fmt.Println("abc", "\141\142\143", "\x61\x62\x63")
	fmt.Println("\u4f17", "\xe4\xbc\x97", "众")

	fmt.Println("1\n1 1\n1 2 1")
	fmt.Println(`1
1 1
1 2 1`)
	// 但是 `` 中不能用转义符
}
