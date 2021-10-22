package fibonacci

import "geekbrains/go/Lesson_8/mycache"

type FibStruct struct {
	Number int64
	Result int64
	Cache  mycache.IntCacher
}

var C1, C2 uint //c1 - счетчик вызова функции c2 - счетчик вызова кеша
var Fb *FibStruct

func (fs *FibStruct) Calculate(n int64, useCache bool) int64 {
	C1 += 1
	if n == 1 || n == -1 || n == 0 {
		fs.Result = n
	} else {
		if useCache {
			C2 += 1
			if n > 0 {
				fs.Result = fs.calculateWithCache(n-2) + fs.calculateWithCache(n-1)
			} else {
				fs.Result = fs.calculateWithCache(n+2) + fs.calculateWithCache(n+1)
			}

			return fs.Result
		}
		if n > 0 {
			fs.Result = fs.Calculate(n-2, useCache) + fs.Calculate(n-1, useCache)
		} else {
			fs.Result = fs.Calculate(n+2, useCache) + fs.Calculate(n+1, useCache)
		}

	}

	return fs.Result
}

func (fs *FibStruct) calculateWithCache(n int64) int64 {
	var data interface{}
	var ok bool

	if data, ok = fs.Cache.GetFromCache(n); !ok {
		fs.Result = fs.Calculate(n, true)
		fs.Cache.AddToCache(n, fs.Result)
	} else {
		if fs.Result, ok = data.(int64); !ok {
			panic("ошибка преобразования типов")
		}
	}

	return fs.Result
}

func init() {
	Fb = &FibStruct{}
}
