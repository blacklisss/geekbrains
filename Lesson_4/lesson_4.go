package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int

	myscanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Введите числа через запятую или пробел: ")
	if ok := myscanner.Scan(); !ok {
		fmt.Println("Возникла ошибка ввода")
		os.Exit(1)
	}

	line := myscanner.Text()

	splitFunc := func(r rune) bool {
		return strings.ContainsRune(", ", r)
	}

	tmp := strings.FieldsFunc(line, splitFunc)
	for _, num := range tmp {
		num = strings.TrimSpace(num)
		n, err := strconv.Atoi(num)

		if err == nil {
			numbers = append(numbers, n)
		}
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
