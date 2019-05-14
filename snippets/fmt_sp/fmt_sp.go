package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Now you have %g problems.\n", 7.777777777777777)
	fmt.Printf("%d\n", 10)

	s1 := fmt.Sprintf("foo: %s", "bar")
	fmt.Println(s1)

	s2 := fmt.Sprint("ho", "ho", "ho")
	fmt.Println(s2)
}

// https://golang.org/pkg/fmt/
// https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing
// https://gobyexample.com/string-formatting
