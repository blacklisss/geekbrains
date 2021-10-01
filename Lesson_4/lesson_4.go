package main

import (
	"fmt"
	"os"
)

func main() {
	var n, tmp int
	var numbers []int

	fmt.Print("Введите количество чисел: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("Введите %d число: ", i)
		if _, err := fmt.Scanln(&tmp); err != nil {
			fmt.Println("Необходимо вводить числа")
			os.Exit(1)
		}

		numbers = append(numbers, tmp)
	}

	fmt.Println("----------------------------")
	fmt.Println("Срез до сортировки:")
	fmt.Println(numbers)
	fmt.Println("----------------------------")

	insertionSort(numbers)

	fmt.Println("Срез после сортировки:")
	fmt.Println(numbers)
}

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
