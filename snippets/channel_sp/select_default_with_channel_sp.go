package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: // 其他 case 都没准备好的时候会执行
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
