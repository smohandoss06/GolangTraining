package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjou!",
		2: "Buenos dias!",
		4: "Bongiorno",
	}
	//fmt.Println(myGreeting)
	delete(myGreeting, 7)
	fmt.Println(myGreeting)
}
