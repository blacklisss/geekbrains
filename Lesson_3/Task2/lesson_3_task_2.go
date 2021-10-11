package main

import (
	"fmt"
	"os"
)

func main() {
	var N int

	fmt.Print("Введите целое число N: ")
	if _, err := fmt.Scanln(&N); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	/*	numbers := make([]int, N+1)

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
		}*/

	for i := 2; i <= N; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Print(i, " ")
		}
	}

	fmt.Print("\n")
}
