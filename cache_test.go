package cache

import (
	"testing"

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

	cache.Put(`abcd`, `efgh`)
	testGetKeyValue(t, cache, `abcd`, `efgh`)

	cache.Put(`ijkl`, `mnop`)
	cache.Put(`ijkl`, `qrst`)
	testGetKeyValue(t, cache, `ijkl`, `qrst`)

}
