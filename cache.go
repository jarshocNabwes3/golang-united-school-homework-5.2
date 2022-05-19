package cache

import "time"

type Cache struct {
}

func NewCache() Cache {
	return Cache{}
}

func (cache Cache) Get(key string) (value string, ok bool) {

	return
}

func (cache Cache) Put(key, value string) {
}

func (cache Cache) Keys() (keys []string) {
	return
}

func (cache Cache) PutTill(key, value string, deadline time.Time) {
}
