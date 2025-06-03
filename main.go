package main

import "fmt"

func main() {
	transactions := [5]int{1, 2, 3, 4, 5}
	transactionsPartial := transactions[1:]
	transactionsPartialNew := transactions[:1]

	transactionsPartial[0] = 30
	fmt.Println(transactions)
	fmt.Println(len(transactionsPartial), cap(transactionsPartial))
	fmt.Println(transactionsPartialNew)
	fmt.Println(len(transactionsPartialNew), cap(transactionsPartialNew))
}
