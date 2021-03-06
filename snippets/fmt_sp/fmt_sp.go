package main

import "fmt"

type T struct {
	foo string
	bar string
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Now you have %g problems.\n", 7.777777777777777)
	fmt.Printf("%d\n", 10)
	fmt.Printf("%q %q %q\n", 1, 1.1, "foo") // quote char or string (escape)

	s1 := fmt.Sprintf("foo: %s", "bar")
	fmt.Println(s1)

	s2 := fmt.Sprint("ho", "ho", "ho")
	fmt.Println(s2)

	t1 := T{foo:"a", bar:"b"}
	fmt.Println(t1)
	fmt.Printf("%s\n", t1)
}

// https://golang.org/pkg/fmt/
// https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing
// https://gobyexample.com/string-formatting
