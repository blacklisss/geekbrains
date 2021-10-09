package main

import (
	"fmt"
	"os"
	"time"
)

type myMap map[int]int

var fn = make(myMap)
var c1, c2 uint //c1 - счетчик вызова функции c2 - счетчик вызова кеша

func fib(n int, withCache bool) int {
	var res int
	c1 += 1
	if n <= 1 {
		res = n
	} else {
		if withCache {
			if fn[n] != 0 {
				c2 += 1
				return fn[n]
			}
		}
		res = fib(n-2, withCache) + fib(n-1, withCache)
	}

	if withCache {
		fn[n] = res
	}

	return res
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

	start := time.Now()

	fmt.Print("Введите число: ")
	if _, err := fmt.Scanln(&n); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if n >= 0 {
		x = fib(n, false)
	} else {
		x = fib2(n, fn)
	}

	fmt.Println("Число Фибоначи:", x)
	fmt.Println("Вызов функций:", c1)
	fmt.Println("Вызов кеша:", c2)

	stop := time.Now()
	fmt.Println("[INFO] Время выполнения:", stop.Sub(start))
}
