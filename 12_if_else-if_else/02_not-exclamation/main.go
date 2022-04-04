package main

import "fmt"

func main() {
	if !true {
		fmt.Println("This is  not ran")
	}
	if !false {
		fmt.Println("This ran")
	}
}
