package cache

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func testGetKeyValue(t *testing.T, cache Cache, key, valueExpected string) {
	value, ok := cache.Get(key)
	assert.Assert(t, ok, `Error to Get() value by '%v' key`, key)
	assert.Equal(t, value, valueExpected, `Unequal '%v' value: '%v' differs from '%v' `, key, value, valueExpected)
}

func TestNewCache(t *testing.T) {
	cache := NewCache()
	assert.DeepEqual(t, cache.keys == nil, false)
}

func TestPutGet(t *testing.T) {
	cache := NewCache()

	cache.Put(`abcd`, `efgh`)
	testGetKeyValue(t, cache, `abcd`, `efgh`)

	cache.Put(`ijkl`, `mnop`)
	cache.Put(`ijkl`, `qrst`)
	testGetKeyValue(t, cache, `ijkl`, `qrst`)

}

func TestKeys(t *testing.T) {
	cache := NewCache()

	cache.Put(`abcd`, `efgh`)
	cache.Put(`ijkl`, `mnop`)
	cache.Put(`ijkl`, `qrst`)

	keys := cache.Keys()
	assert.DeepEqual(t, keys, []string{`abcd`, `ijkl`})

}

func TestPutTill(t *testing.T) {
	cache := NewCache()
	cache.PutTill(`abcd`, `efgh`, time.Now().Add(200*time.Millisecond))
	testGetKeyValue(t, cache, `abcd`, `efgh`)
	time.Sleep(100 * time.Millisecond)
	testGetKeyValue(t, cache, `abcd`, `efgh`)

	time.Sleep(200 * time.Millisecond)

	_, ok := cache.Get(`abcd`)
	assert.Assert(t, !ok, `No error to Get() value by '%v' key after 300 msec`, `abcd`)
}
