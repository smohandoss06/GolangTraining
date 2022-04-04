package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjour!",
	}
	myGreeting["Selva"] = "Cybertruck"
	fmt.Println(len(myGreeting))

}
