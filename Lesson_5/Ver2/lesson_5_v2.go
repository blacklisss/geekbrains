package main

import (
	"fmt"
	"os"
	"time"
)

type myMap map[int]int

func fib(n int, fn myMap) int {
	if n <= 1 {
		fn[n] = n
	} else if fn[n] == 0 {
		fn[n] = fib(n-2, fn) + fib(n-1, fn)
	}

	return fn[n]
}

func fib2(n int, fn myMap) int {
	if n >= -1 {
		fn[n] = n
	} else if fn[n] == 0 {
		fn[n] = fib2(n+2, fn) + fib2(n+1, fn)
	}

	return fn[n]
}

func main() {
	var n, x int
	fn := make(myMap)

	start := time.Now()

	fmt.Print("Введите число: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if n >= 0 {
		x = fib(n, fn)
	} else {
		x = fib2(n, fn)
	}

	fmt.Println("Число Фибоначи:", x)

	stop := time.Now()
	fmt.Println("[INFO] Время выполнения:", stop.Sub(start))
}
