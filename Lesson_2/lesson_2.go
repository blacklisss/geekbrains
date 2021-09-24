package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Task_1 = iota + 1
	Task_2
	Task_3
	Task_3_2
)

func main() {
	var task_number int
	fmt.Printf(`
Выберите задание: 
1) Вычислене площади прямоугольника
2) Вычисление диаметра и длины окружности по заданной площади круга
3) Вывод цифр, соответствующие количество сотен, десятков и единиц в числе
4) Вывод цифр, соответствующие количество сотен, десятков и единиц в числе (Вариант 2)
`)

	fmt.Print("Введите номер задания (число): ")
	if _, err := fmt.Scanf("%d", &task_number); err != nil {
		fmt.Println(err)
		return
	}

	switch task_number {
	case Task_1:
		task1()
		break
	case Task_2:
		task2()
		break
	case Task_3:
		task3()
		break
	case Task_3_2:
		task3_2()
		break
	default:
		fmt.Println("Номер задания должен быть от 1 до 3")
	}
}

func task1() {
	var a, b int // Длины сторон

	fmt.Println("-------------------------")
	fmt.Print("Введите длины сторон a и b (через пробел): ")
	if _, err := fmt.Scanf("%d %d", &a, &b); err != nil {
		fmt.Println("Длины сторон должны быть числами")
		return
	}

	fmt.Println("Площадь прямоугольника: ", a*b)
}

func task2() {
	var s float64    // Площать круга
	var d, l float64 // d-диаметр круга, l-длина окружности

	fmt.Println("-------------------------")
	fmt.Print("Введите площать круга: ")
	if _, err := fmt.Scanf("%f", &s); err != nil {
		fmt.Println("Площадь круга должна быть числом")
		return
	}

	d = 2 * math.Sqrt(s/math.Pi)
	fmt.Printf("Диаметр круга: %.2f \n\r", d)

	l = math.Pi * d
	fmt.Printf("Длина окружности: %.2f \n\r", l)
}

func task3() {
	var number int

	fmt.Println("-------------------------")
	fmt.Print("Введите трехзначное число: ")
	if _, err := fmt.Scanf("%d", &number); err != nil {
		fmt.Println(err)
		return
	}

	str := strconv.Itoa(number)

	if len(str) != 3 {
		fmt.Println("Число должно быть 3-х значное")
		return
	}

	fmt.Printf("Сотни: %s\n\rДесятки: %s\n\rЕдиницы: %s\n\r", string(str[0]), string(str[1]), string(str[2]))
}

func task3_2() {
	var number, count, nTmp int
	var numbersArray []int // Результирующий массив

	fmt.Println("-------------------------")
	fmt.Print("Введите трехзначное число: ")
	if _, err := fmt.Scanf("%d", &number); err != nil {
		fmt.Println(err)
		return
	}

	nTmp = number
	for {
		count++
		nTmp /= 10
		if nTmp == 0 {
			break
		}
	}

	if count != 3 {
		fmt.Println("Число должно быть 3-х значное")
		return
	}

	numbersArray = append(numbersArray, number/100)
	numbersArray = append(numbersArray, number/10%10)
	numbersArray = append(numbersArray, number%10)

	fmt.Printf("Сотни: %d\n\rДесятки: %d\n\rЕдиницы: %d\n\r", numbersArray[0], numbersArray[1], numbersArray[2])
}
