package main

import "fmt"

func main() {
	transactions := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
}
