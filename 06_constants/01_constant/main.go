package main

import "fmt"

const p = "death & taxes"

const (
	pi       = 3.14
	language = "Go"
)

const (
	C0 = iota
	C1 = iota
	C2 = iota
)
const (
	a = iota
	b
	c
)
const (
	d = iota
	e
	f
)
const (
	_   = iota
	b11 = iota * 10
	c11 = iota * 20
)

func main() {
	const q = 2
	fmt.Println("p -", p)
	fmt.Println("q -", q)
	fmt.Println(pi)
	fmt.Println(language)
	fmt.Println(C0)
	fmt.Println(C1)
	fmt.Println(C2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	//fmt.Println(_)
	fmt.Println(c11)
	fmt.Println(b11)

}
