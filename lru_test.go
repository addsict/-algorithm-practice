package lru

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	c := NewLruCache(3)

	c.Set("a", "foo")
	if keys := c.RecentlyUsedKeys(); !isSameSlice(keys, []string{"a"}) {
		t.Errorf("keys expected %v, but got %v", []string{"a"}, keys)
	}

	c.Set("b", "bar")
	if keys := c.RecentlyUsedKeys(); !isSameSlice(keys, []string{"b", "a"}) {
		t.Errorf("keys expected %v, but got %v", []string{"b", "a"}, keys)
	}

	c.Set("c", "baz")
	if keys := c.RecentlyUsedKeys(); !isSameSlice(keys, []string{"c", "b", "a"}) {
		t.Errorf("keys expected %v, but got %v", []string{"c", "b", "a"}, keys)
	}

	c.Set("d", "hoge")
	if keys := c.RecentlyUsedKeys(); !isSameSlice(keys, []string{"d", "c", "b"}) {
		t.Errorf("keys expected %v, but got %v", []string{"d", "c", "b"}, keys)
	}

	c.Set("c", "baz")
	if keys := c.RecentlyUsedKeys(); !isSameSlice(keys, []string{"c", "d", "b"}) {
		t.Errorf("keys expected %v, but got %v", []string{"c", "d", "b"}, keys)
	}

	c.Set("c", "baz")
	if keys := c.RecentlyUsedKeys(); !isSameSlice(keys, []string{"c", "d", "b"}) {
		t.Errorf("keys expected %v, but got %v", []string{"c", "d", "b"}, keys)
	}

	fmt.Println(c)
}

func isSameSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
