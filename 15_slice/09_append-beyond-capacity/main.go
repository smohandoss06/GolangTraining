package main

import "fmt"

func main() {
	greeting := make([]string, 3, 4)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "gidday")
	greeting = append(greeting, "fdfsdf")
	fmt.Println(greeting)
}
