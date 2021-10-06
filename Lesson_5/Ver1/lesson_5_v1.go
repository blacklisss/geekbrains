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

func main() {
	var n int

	start := time.Now()

	fmt.Print("Введите число: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	x := fib(n)
	fmt.Println("Число Фибоначи:", x)

	stop := time.Now()
	fmt.Println("[INFO] Время выполнения:", stop.Sub(start))
}
