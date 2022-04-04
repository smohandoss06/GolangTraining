package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Bonjout",
	}
	fmt.Println(myGreeting["Tim"])
}
