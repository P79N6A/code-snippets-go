package main

import (
	"fmt"
	"github.com/binderclip/code-snippets-go/utils"
	"strings"
)

func string1() {
	utils.PrintFuncName()
	fmt.Println("abc", "\141\142\143", "\x61\x62\x63")
	fmt.Println("\u4f17", "\xe4\xbc\x97", "众")

	fmt.Println("1\n1 1\n1 2 1")
	fmt.Println(`1
1 1
1 2 1`)
	// 但是 `` 中不能用转义符
}

func split() {
	utils.PrintFuncName()
	s := "ab.c.de."
	fmt.Println(strings.Split(s, "."))
}

func main() {
	string1()
	split()
}

// https://gfw.go101.org/article/basic-types-and-value-literals.html
