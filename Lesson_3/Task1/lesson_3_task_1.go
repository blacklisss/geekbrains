package main

import (
	"fmt"
	"math"
	"os"
)

const (
	OperationPlus     = "+"
	OperationMinus    = "-"
	OperationMultiply = "*"
	OperationDivision = "/"
	OperationSQRT     = "sqrt"
)

func main() {
	var operation string
	var a, b float64
	var err error

	fmt.Printf(`
Введите операцию: 
Сложение: %s
Вычетание: %s
Умножение: %s
Деление: %s
Квадратный корень: %s
`, OperationPlus, OperationMinus, OperationMultiply, OperationDivision, OperationSQRT)

	if _, err = fmt.Scanln(&operation); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if a, b, err = inputNumbers(operation); err != nil {
		fmt.Println(err.Error())
	}

	switch operation {
	case OperationSQRT:
		result := math.Sqrt(a)
		fmt.Printf("Результат операции %s(%.2f) = %.2f\n\r", operation, a, result)
	case OperationPlus:
		result := a + b
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, result)
	case OperationMinus:
		result := a - b
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, result)
	case OperationMultiply:
		result := a * b
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, result)
	case OperationDivision:
		if b == 0 {
			fmt.Println("Деление на 0 запрещено")
			os.Exit(1)
		}
		result := a / b
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, result)
	}
}

func inputNumbers(operation string) (a, b float64, err error) {

	switch operation {
	case OperationSQRT:
		fmt.Print("Введите число: ")
		if _, err = fmt.Scanln(&a); err != nil {
			return
		}
	case OperationPlus, OperationMinus, OperationMultiply, OperationDivision:
		fmt.Print("Введите два числа через пробел: ")
		if _, err = fmt.Scanf("%f %f", &a, &b); err != nil {
			return
		}
	default:
		err = fmt.Errorf("Выбрана неправильная операция")
		return
	}

	return
}
