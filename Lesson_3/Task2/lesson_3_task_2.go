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

	for i := 0; i <= N; i++ {
		numbers[i] = i
	}

	limit := int(math.Sqrt(float64(N))) + 1

	for i := 2; i < limit; i++ {
		if numbers[i] != 0 {
			for j := i * i; j <= N; j += i {
				numbers[j] = 0
			}
		}
	}

	for i := 2; i <= N; i++ {
		if numbers[i] != 0 {
			fmt.Println(numbers[i])
		}
	}

}
