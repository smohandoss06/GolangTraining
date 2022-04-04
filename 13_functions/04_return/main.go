package main

import "fmt"

func main() {
	fmt.Println(greet("Selva", "Mohandoss"))

}

func greet(fname, lname string) string {
	return fmt.Sprintf(fname, lname)
}
