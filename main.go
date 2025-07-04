package main

import (
	"fmt"
	"sort"
	"strings"
)

// Определяем тип функции для операций
type OperationFunc func(*[]int) int

func main() {
	// Создаем map с функциями
	operations := map[string]OperationFunc{
		"1": average,
		"2": sum,
		"3": median,
	}

	for {
		fmt.Println("Продвинутый калькулятор")

		operation := operationInput()
		numbers := intInput()

		if len(numbers) == 0 {
			fmt.Println("Вы не ввели ни одного числа.")
			continue
		}

		result := calculateOperation(operation, &numbers, &operations)
		fmt.Println("Результат:", result)

		fmt.Println("Вы хотите продолжить? (y/n)")
		var continueInput string
		fmt.Scanln(&continueInput)

		if continueInput != "y" && continueInput != "Y" {
			break
		}
	}
}

func operationInput() string {
	var operation string

	fmt.Println("Выберите операцию:")
	fmt.Println("1. Среднее (AVG)")
	fmt.Println("2. Сумма (SUM)")
	fmt.Println("3. Медиана (MED)")

	fmt.Scanln(&operation)
	return operation
}

func intInput() []int {
	var input string
	fmt.Println("Введите числа через запятую:")
	fmt.Scanln(&input)

	var numbers []int
	parts := strings.Split(input, ",")

	for _, part := range parts {
		var num int
		_, err := fmt.Sscanf(strings.TrimSpace(part), "%d", &num)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func sum(numbers *[]int) int {
	total := 0
	for _, n := range *numbers {
		total += n
	}
	return total
}

func average(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}
	return sum(numbers) / len(*numbers)
}

func median(numbers *[]int) int {
	if len(*numbers) == 0 {
		return 0
	}

	sorted := make([]int, len(*numbers))
	copy(sorted, *numbers)
	sort.Ints(sorted)

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

func calculateOperation(operation string, numbers *[]int, operations *map[string]OperationFunc) int {
	if opFunc, exists := (*operations)[operation]; exists {
		return opFunc(numbers)
	}
	fmt.Println("Неверный выбор операции.")
	return 0
}
