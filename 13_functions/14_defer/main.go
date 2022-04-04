package main

import "fmt"

func hello() {
	fmt.Printf("hello")
}
func world() {
	fmt.Printf("world")
}
func test() {
	fmt.Printf("test")
}

//defer delay the execution  of the program
func main() {

	defer hello()
	world()
	defer test()
}
