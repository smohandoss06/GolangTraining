package main

import "fmt"

func main() {
	switch "Mhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}

/*
fallthrough statement transfer controls to the next case
no default fallthrough
fallthrough is optional
-- you can specify fallthrough by explicitly stating it
-- break isn't needed like in other languages
*/
