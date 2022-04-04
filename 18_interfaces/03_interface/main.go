package main

import "fmt"

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {

	s := square{10}
	info(s)

}
