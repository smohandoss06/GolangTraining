package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Bonjour!",
		2: "Buenos dias",
		3: "Bongiorno!",
	}
	fmt.Println(myGreeting)
	if val, exists := myGreeting[2]; exists {
		fmt.Println("The value exists", exists)
		fmt.Println("The value", val)
	} else {
		fmt.Println("The value doesn't exists")
	}
	fmt.Println(myGreeting)
}
