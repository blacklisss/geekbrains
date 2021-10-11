package main

import (
	"fmt"
	"geekbrains/go/Lesson_7/mycache"
	"os"
	"time"
)

var c1, c2 uint //c1 - счетчик вызова функции c2 - счетчик вызова кеша

type Cacher mycache.IntCacher

type FibStruct struct {
	number int64
	result int64
}

func (fs *FibStruct) Calculate(n int64, cache Cacher) int64 {
	var data interface{}
	var err error
	var ok bool

	c1 += 1
	if n == 1 || n == -1 || n == 0 {
		fs.result = n
	} else {
		if cache.IsUseCache() {
			if data, err = cache.GetFromCache(n); err == nil {
				c2 += 1
				if fs.result, ok = data.(int64); !ok {
					panic("ошибка преобразования типов")
				}
				return fs.result
			}
		}
		if n > 0 {
			fs.result = fs.Calculate(n-2, cache) + fs.Calculate(n-1, cache)
		} else {
			fs.result = fs.Calculate(n+2, cache) + fs.Calculate(n+1, cache)
		}

	}

	if cache.IsUseCache() {
		cache.AddToCache(n, fs.result)
	}

	return fs.result
}

func main() {
	start := time.Now()

	fb := &FibStruct{}

	fmt.Print("Введите число: ")
	if _, err := fmt.Scanln(&fb.number); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mapCache := &mycache.MapCacheInt{
		UseCache: true,
		Cache:    make(map[int64]int64),
	}

	_ = fb.Calculate(fb.number, mapCache)

	fmt.Println("Число Фибоначи:", fb.result)
	fmt.Println("Вызов функций:", c1)
	fmt.Println("Вызов кеша:", c2)

	stop := time.Now()
	fmt.Println("[INFO] Время выполнения:", stop.Sub(start))
}
