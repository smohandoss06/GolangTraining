package main

import "fmt"

func greeting() {
	fmt.Println("Hello World")
}

func main() {
	greeting()

	greet := func() {
		fmt.Println("Hello World")
	}
	greet()
	fmt.Printf("%T\n", greet)
	greeter := makeGreeter()
	fmt.Println(greeter())
}

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}
