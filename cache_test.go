package cache

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()
	m := make(map[string]string)
	assert.DeepEqual(t, m, cache.keys)
	cache.Put(`abcd`, `efgh`)

	value, ok := cache.Get(`abcd`)
	assert.Assert(t, ok, `Error to Get() value by 'abcd' key`)
	assert.Equal(t, value, `efgh`, `Unequal '%v' value for 'abcd' key`, value)

	cache.Put(`ijkl`, `mnop`)
	cache.Put(`ijkl`, `qrst`)

	value, ok = cache.Get(`ijkl`)
	assert.Assert(t, ok, `Error to Get() value by 'ijkl' key`)
	assert.Equal(t, value, `qrst`, `Unequal '%v' value for 'qrst' key`, value)
}
