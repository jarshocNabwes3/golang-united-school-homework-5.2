package cache

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()
	var m map[string]string
	assert.DeepEqual(t, m, cache.keys)
	cache.Put(`abcd`, `efgh`)
	value, ok := cache.Get(`abcd`)
	assert.Assert(t, ok, `Error to Get() value by 'abcd' key`)
	assert.Equal(t, value, `efgh`, `Unequal value for 'abcd' key`)
}
