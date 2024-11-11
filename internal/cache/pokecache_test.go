package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c1CacheMapTest := make(map[string]cacheEntry)

	c1CacheMapTest["hello"] = cacheEntry{
		createdAt: time.Now(),
		val:       []byte("world"),
	}

	c1Test := Cache{
		cacheMap: c1CacheMapTest,
	}

	c1CacheMapExpected := make(map[string]cacheEntry)

	c1CacheMapExpected["hello"] = cacheEntry{
		createdAt: c1CacheMapTest["hello"].createdAt,
		val:       []byte("world"),
	}

	c1Expected := Cache{
		cacheMap: c1CacheMapExpected,
	}

	cases := []struct {
		input    Cache
		expected Cache
	}{
		{
			input:    c1Test,
			expected: c1Expected,
		},
	}

	for _, c := range cases {
		if c.input.cacheMap["hello"].createdAt != c.expected.cacheMap["hello"].createdAt {
			t.Errorf("error: time dates don't match up")
		}
	}
}

func TestAddGetCache(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "hello",
			val: []byte("world!"),
		},
		{
			key: "world!",
			val: []byte("hello"),
		},
	}

	cache := NewCache(10)

	for _, v := range cases {
		err := cache.Add(v.key, v.val)
		if err != nil {
			t.Errorf("adding failed")
		}
	}

	for _, v := range cases {
		val, exists := cache.Get(v.key)
		if !exists {
			t.Errorf("key does not exist")
		}

		if string(val) != string(v.val) {
			t.Errorf("value does not match")
		}
	}

}
