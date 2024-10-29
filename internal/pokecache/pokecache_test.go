package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("newtestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key `%s`", c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf(
					"expected to find value `%s` under key `%s`, found `%s`",
					c.val,
					c.key,
					val,
				)
				return
			}
		})
	}
}

func TestReap(t *testing.T) {
	const interval = 5 * time.Millisecond
	cache := NewCache(interval)
	key := "https://reaploopexample.com"
	value := []byte("testdata")
	cache.Add(key, value)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key `%s`", key)
		return
	}

	time.Sleep(interval + 1 * time.Millisecond)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected NOT to find key `%s`", key)
		return
	}
}

func TestReapFail(t *testing.T) {
	const interval = 5 * time.Millisecond
	cache := NewCache(interval)
	key := "https://reaploopexample.com"
	value := []byte("testdata")
	cache.Add(key, value)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key `%s`", key)
		return
	}

	time.Sleep(interval / 2)

	_, ok = cache.Get(key)
	if !ok {
		t.Errorf("expected to still find key `%s`", key)
		return
	}
}
