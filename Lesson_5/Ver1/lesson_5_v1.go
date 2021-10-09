package main

import (
	"fmt"
	"os"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-2) + fib(n-1)
}

func fib2(n int) int {
	if n >= -1 {
		return n
	}

	return fib2(n+2) + fib2(n+1)
}

func main() {
	var n, x int

	start := time.Now()

	fmt.Print("Введите число: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if n >= 0 {
		x = fib(n)
	} else {
		x = fib2(n)
	}

	fmt.Println("Число Фибоначи:", x)

	stop := time.Now()
	fmt.Println("[INFO] Время выполнения:", stop.Sub(start))
}
