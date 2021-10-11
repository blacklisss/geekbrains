package mycache

import "fmt"

type IntCacher interface {
	AddToCache(key int64, data int64)
	GetFromCache(key int64) (data interface{}, err error)
	IsUseCache() bool
}

type MapCacheInt struct {
	UseCache bool
	Cache    map[int64]int64
}

func (mc *MapCacheInt) AddToCache(key int64, data int64) {
	mc.Cache[key] = data
}

func (mc *MapCacheInt) GetFromCache(key int64) (dat interface{}, err error) {
	var ok bool
	if dat, ok = mc.Cache[key]; !ok {
		return nil, fmt.Errorf("ошибка чтения кеша с ключем %d", key)
	}

	return dat, nil
}

func (mc *MapCacheInt) IsUseCache() bool {
	return mc.UseCache
}
