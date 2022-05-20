package cache

import "time"

type Cache struct {
	keys map[string]string
}

func NewCache() Cache {
	cache := Cache{}
	cache.keys = make(map[string]string)
	return cache
}

func (cache Cache) Get(key string) (value string, ok bool) {
	value, ok = cache.keys[key]

	return
}

func (cache Cache) Put(key, value string) {
	cache.keys[key] = value
}

func (cache Cache) Keys() (keysArr []string) {

	for key := range cache.keys {
		keysArr = append(keysArr, key)
	}

	return
}

func (cache Cache) removeKeyByDeadline(key string, deadline time.Time) {
	duration := time.Until(deadline)
	time.Sleep(duration)

	delete(cache.keys, key)
}

func (cache Cache) PutTill(key, value string, deadline time.Time) {
	cache.Put(key, value)

	go cache.removeKeyByDeadline(key, deadline)
}
