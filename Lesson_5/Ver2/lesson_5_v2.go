package main

import (
	"fmt"
	"os"
	"time"
)

type myMap map[int]int

func fib(n int, fn myMap) int {
	for i := 0; i <= n; i++ {
		var f int
		if i == 0 {
			f = 0
		} else if i == 1 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	var n int
	fn := make(myMap)

	start := time.Now()

	fmt.Print("Введите число: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	x := fib(n, fn)
	fmt.Println("Число Фибоначи:", x)

	stop := time.Now()
	fmt.Println("[INFO] Время выполнения:", stop.Sub(start))
}
