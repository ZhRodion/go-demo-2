package main

import "fmt"

func main() {

	transactions := []float64{}

	for {

		inputTransaction(&transactions)

		printCalculatedTransactions(transactions)

		fmt.Println("Вы хотите продолжить? (y/n)")
		var continueInput string
		fmt.Scanln(&continueInput)

		if continueInput == "y" || continueInput == "Y" {
			continue
		} else {
			break
		}
	}
}

func inputTransaction(transactions *[]float64) {
	var inputTransaction float64

	for {
		fmt.Println("Введите вашу транзакцию: ")
		fmt.Scanln(&inputTransaction)

		if inputTransaction < 0 {
			fmt.Println("Транзакция не может быть отрицательной!")
			continue
		}
		break
	}

	*transactions = append(*transactions, inputTransaction)
}

func printCalculatedTransactions(transactions []float64) {
	balance := 0.0

	for _, value := range transactions {
		balance += value
		fmt.Println(balance)
	}
}
