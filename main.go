package main

import (
	"fmt"
	"sort"
	"strings"
)

// Глобальные переменные для хранения данных
var currentNumbers []int
var currentResult int

func main() {
	operations := map[string]func(){
		"1":   calculateAverage,
		"2":   calculateSum,
		"3":   calculateMedian,
		"4":   calculateMin,
		"5":   calculateMax,
		"avg": calculateAverage,
		"sum": calculateSum,
		"med": calculateMedian,
		"min": calculateMin,
		"max": calculateMax,
	}

	// Добавляем динамические операции с замыканиями
	addDynamicOperations(operations)

	for {
		fmt.Println("Продвинутый калькулятор")

		operation := operationInput()
		currentNumbers = intInput()

		if len(currentNumbers) == 0 {
			fmt.Println("Вы не ввели ни одного числа.")
			continue
		}

		// Выполняем операцию через map
		if operationFunc, exists := operations[operation]; exists {
			operationFunc() // Вызываем функцию
			fmt.Println("Результат:", currentResult)
		} else {
			fmt.Println("Неверный выбор операции.")
		}

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
	fmt.Println("1 или avg - Среднее (AVG)")
	fmt.Println("2 или sum - Сумма (SUM)")
	fmt.Println("3 или med - Медиана (MED)")
	fmt.Println("4 или min - Минимум (MIN)")
	fmt.Println("5 или max - Максимум (MAX)")
	fmt.Println("pow2 - Возведение в квадрат всех чисел")
	fmt.Println("count - Количество чисел")

	fmt.Scanln(&operation)
	return strings.ToLower(operation)
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

// Функции-обработчики операций
func calculateSum() {
	total := 0
	for _, n := range currentNumbers {
		total += n
	}
	currentResult = total
}

func calculateAverage() {
	if len(currentNumbers) == 0 {
		currentResult = 0
		return
	}
	calculateSum() // Используем уже готовую функцию
	currentResult = currentResult / len(currentNumbers)
}

func calculateMedian() {
	if len(currentNumbers) == 0 {
		currentResult = 0
		return
	}

	sorted := make([]int, len(currentNumbers))
	copy(sorted, currentNumbers)
	sort.Ints(sorted)

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		currentResult = (sorted[mid-1] + sorted[mid]) / 2
	} else {
		currentResult = sorted[mid]
	}
}

func calculateMin() {
	if len(currentNumbers) == 0 {
		currentResult = 0
		return
	}
	min := currentNumbers[0]
	for _, n := range currentNumbers {
		if n < min {
			min = n
		}
	}
	currentResult = min
}

func calculateMax() {
	if len(currentNumbers) == 0 {
		currentResult = 0
		return
	}
	max := currentNumbers[0]
	for _, n := range currentNumbers {
		if n > max {
			max = n
		}
	}
	currentResult = max
}

// Функция для добавления динамических операций с замыканиями
func addDynamicOperations(operations map[string]func()) {
	// Замыкание для возведения в квадрат
	operations["pow2"] = func() {
		total := 0
		for _, n := range currentNumbers {
			total += n * n
		}
		currentResult = total
	}

	// Замыкание для подсчета количества
	operations["count"] = func() {
		currentResult = len(currentNumbers)
	}
}
