package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		a, ok1 := <-ch1
		b, ok2 := <-ch2
		// fmt.Println(a, ok1, b, ok2)
		if !(ok1 || ok2) {
			return true // 同时结束，相同
		} else if !(ok1 && ok2) {
			return false // 不同时结束，不相同
		}
		if a != b {
			return false // 出现不相同元素，不相同
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
