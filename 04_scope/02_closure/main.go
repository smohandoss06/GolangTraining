package main

import "fmt"

var i = 0

func increment() int {
	i++
	return i
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The credit belongs to the one who is in the ring"
		fmt.Println(y)
	}
	fmt.Println(increment())
	fmt.Println(increment())
	y := 0
	yincrement := func() int {
		y++
		return y
	}
	fmt.Println(yincrement())
	fmt.Println(yincrement())

	wrapperoutput := wrapper()
	fmt.Println(wrapperoutput())
	fmt.Println(wrapperoutput())

}
