package main

import (
	"fmt"
	"os"
)
import "math"

func main() {
	var N int

	fmt.Print("Введите целое число N: ")
	if _, err := fmt.Scanln(&N); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	numbers := make([]int, N+1)
	seive := make([]bool, N+1)

	for i := 0; i <= N; i++ {
		numbers[i] = i
	}

	limit := int(math.Sqrt(float64(N))) + 1

	for i := 2; i < limit; i++ {
		if !seive[i] {
			for j := i * i; j <= N; j += i {
				seive[j] = true
			}
		}
	}

	count := 0
	for i := 2; i <= N; i++ {
		if !seive[i] {
			count++
		}
	}

	fmt.Printf("Простых чисел от 0 до %d включительно = %d\n", N, count)
	fmt.Print("Простые числа: ")

	for i := 2; i <= N; i++ {
		if !seive[i] {
			fmt.Print(numbers[i], " ")
		}
	}
	fmt.Print("\n\r")
}
