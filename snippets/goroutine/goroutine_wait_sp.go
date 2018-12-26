package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	time.Sleep(400)
	fmt.Println(s)
	wg.Done()
}

func main() {
	go say("world")
	wg.Add(1)
	fmt.Println("hello")
	wg.Wait()
}

// https://programming.guide/go/wait-for-goroutines-waitgroup.html
