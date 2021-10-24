package main

import (
	"fmt"
	"geekbrains/go/Lesson_8/configuration"
	"geekbrains/go/Lesson_8/fibonacci"
	"geekbrains/go/Lesson_8/mycache"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var config *configuration.Config
	var err error

	if config, err = configuration.Load("configuration/config.env"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print("Введите число: ")
	if _, err = fmt.Scanln(&fibonacci.Fb.Number); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	mapCache := mycache.NewCache()
	mapCache.SetUseCache(config.UseCache)

	fibonacci.Fb.Cache = mapCache

	_ = fibonacci.Fb.Calculate(fibonacci.Fb.Number, mapCache.IsUseCache())

	fmt.Println("Число Фибоначи:", fibonacci.Fb.Result)
	fmt.Println("Вызов функций:", fibonacci.C1)
	fmt.Println("Вызов кеша:", fibonacci.C2)
	fmt.Println("Вывод поля конфигурации [URL]:", config.Url)

	fmt.Println("[INFO] Время выполнения:", time.Since(start).Seconds(), "сек")
}
