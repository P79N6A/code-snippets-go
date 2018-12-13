package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Person2 struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func personStringer() {
	fmt.Println("=== personStringer ===")

	a := Person{"Arthur Dent", 42}
	z := Person2{"Zaphod Beeblebrox", 9001}
	fmt.Println(a)
	fmt.Println(z)
}

type IPAddr [4]byte

func (v IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", v[0], v[1], v[2], v[3])
}

func ipAddrStringer() {
	fmt.Println("=== ipAddrStringer ===")
	ip := IPAddr{192, 168, 1, 222}
	fmt.Println(ip)

	ip2 := [4]byte{192, 168, 1, 222}
	fmt.Println(ip2)
}

func main() {
	personStringer()
	ipAddrStringer()
}
