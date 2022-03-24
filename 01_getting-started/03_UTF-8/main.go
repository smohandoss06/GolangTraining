package main

import "fmt"

func main() {
	//%q UTF character literal
	for i := 60; i < 122; i++ {

		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
