package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good Morning"
	myGreeting["Jenny"] = "Will"
	fmt.Println(myGreeting)
}
