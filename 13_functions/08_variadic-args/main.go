package main

import "fmt"

func main() {
	data := []float64{43, 455, 444, 443}
	n := average(data...)
	fmt.Println(n)

}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
