package cache

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewCache(t *testing.T) {
	cache := NewCache()
	assert.Equal(t, Cache{}, cache)
}
