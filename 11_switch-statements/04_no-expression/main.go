package main

import "fmt"

func main() {
	myFriendsName := "Mar"
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	default:
		fmt.Println("nothing matched, this is the default")
	}
	//expression not needed
	//if no expression provided,go check for the first case
}
