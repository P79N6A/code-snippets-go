package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount 计算句子中不同词语出现的次数
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, field := range strings.Fields(s) {
		m[field]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
