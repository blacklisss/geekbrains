package mycache

type IntCacher interface {
	AddToCache(key int64, data int64)
	GetFromCache(key int64) (data interface{}, ok bool)
	IsUseCache() bool
}

type MapCacheInt struct {
	useCache bool
	Cache    map[int64]int64
}

type Cacher IntCacher

func (mc *MapCacheInt) AddToCache(key int64, data int64) {
	mc.Cache[key] = data
}

func (mc *MapCacheInt) GetFromCache(key int64) (dat interface{}, ok bool) {
	if dat, ok = mc.Cache[key]; !ok {
		return nil, ok
	}

	return dat, ok
}

func (mc *MapCacheInt) IsUseCache() bool {
	return mc.useCache
}

func (mc *MapCacheInt) SetUseCache(useCache bool) {
	mc.useCache = useCache
}

func NewCache() *MapCacheInt {
	return &MapCacheInt{
		Cache: make(map[int64]int64),
	}
}
