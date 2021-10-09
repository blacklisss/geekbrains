/**
1)
	а) Не совсем понял вопрос, но видимо речь идет о работе с вводом данных (Scan и т.п.)
	в) При передаче карт, срезов и при захвате внешних переменных замыканиями

2) По количеству операндов и по их типу
*/

package main

import "fmt"

func test_slice(arr []int) {
	fmt.Printf("inner: %v, %p\n", arr, arr)
	arr[1] = 11
	fmt.Printf("inner: %v, %p\n", arr, arr)
}

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("outer i: %v, %p\n", i, &i)
		defer func() {
			fmt.Printf("inner i: %v, %p\n", i, &i)
		}()
	}

	m := []int{
		1,
		2,
		3,
	}

	fmt.Printf("outer: %v, %p\n", m, m)
	test_slice(m)
	fmt.Printf("outer: %v, %p\n", m, m)

	a := 5
	var b = new(int)
	var c **int

	*b = 3
	c = &b

	fmt.Println(a * *b)
	fmt.Println(a * **c)
}
