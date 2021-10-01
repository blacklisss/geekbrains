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
		os.Exit(1)
	}

	switch operation {
	case OperationSQRT:
		fmt.Printf("Результат операции %s(%.2f) = %.2f\n\r", operation, a, math.Sqrt(a))
	case OperationPlus:
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, a+b)
	case OperationMinus:
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, a-b)
	case OperationMultiply:
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, a*b)
	case OperationDivision:
		fmt.Printf("Результат операции %.2f %s %.2f = %.2f\n\n", a, operation, b, a/b)
	}
}

func inputNumbers(operation string) (a, b float64, err error) {

	switch operation {
	case OperationSQRT:
		fmt.Print("Введите число: ")
		if _, err = fmt.Scanln(&a); err != nil {
			return
		}
	case OperationPlus:
		fallthrough
	case OperationMinus:
		fallthrough
	case OperationMultiply:
		fmt.Print("Введите два числа через пробел: ")
		if _, err = fmt.Scanf("%f %f", &a, &b); err != nil {
			return
		}
	case OperationDivision:
		fmt.Print("Введите два числа через пробел: ")
		if _, err = fmt.Scanf("%f %f", &a, &b); err != nil {
			return
		}

		if b == 0 {
			err = fmt.Errorf("Деление на 0 запрещено")
			return
		}
	default:
		err = fmt.Errorf("Выбрана неправильная операция")
		return
	}

	return
}
