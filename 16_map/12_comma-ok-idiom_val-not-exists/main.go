package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	if val, exists := myGreeting[7]; exists {
		delete(myGreeting, 7)
		fmt.Println(val)
		fmt.Println("exists :", exists)
	} else {
		fmt.Println("The value doesn't exists")
	}
}
