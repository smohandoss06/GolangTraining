package main

import "fmt"

func main() {
	name := "Todd"
	//fmt.Println(name)
	changeMe(&name)
	//fmt.Println(name)
}

//Pointers can be datatype
func changeMe(z *string) {

	fmt.Println(z)
	fmt.Println(*z)
	*z = "Rocky"
	fmt.Println(z)
	fmt.Println(*z)
}
