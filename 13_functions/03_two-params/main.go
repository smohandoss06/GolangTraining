package main

import "fmt"

func main() {
	greet("selva", "mohandoss")

}

func greet(fname string, lname string) {
	fmt.Println(fname + lname)
}
