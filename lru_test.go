package lru

import (
	"testing"
)

func TestLruCache(t *testing.T) {
	t.Run("check internal linked list", func(t *testing.T) {
		checkInternalLinkedList(t, func(c *LruCache) {
			c.Set("a", "foo")
		}, []string{"a"})

		checkInternalLinkedList(t, func(c *LruCache) {
			c.Set("a", "foo")
			c.Set("b", "bar")
			c.Set("c", "baz")
		}, []string{"c", "b", "a"})

		checkInternalLinkedList(t, func(c *LruCache) {
			c.Set("a", "foo")
			c.Set("b", "bar")
			c.Set("c", "baz")
			c.Set("d", "hoge")
		}, []string{"d", "c", "b"})

		checkInternalLinkedList(t, func(c *LruCache) {
			c.Set("a", "foo")
			c.Set("b", "bar")
			c.Set("c", "baz")
			c.Set("d", "hoge")
			c.Set("c", "hoge")
		}, []string{"c", "d", "b"})

		checkInternalLinkedList(t, func(c *LruCache) {
			c.Set("a", "foo")
			c.Set("b", "bar")
			c.Set("c", "baz")
			c.Set("d", "hoge")
			c.GetString("b")
		}, []string{"b", "d", "c"})

		checkInternalLinkedList(t, func(c *LruCache) {
			c.Set("a", "foo")
			c.Set("b", "bar")
			c.Set("c", "baz")
			c.Remove("b")
			c.Set("d", "hoge")
		}, []string{"d", "c", "a"})
	})
}

func checkInternalLinkedList(t *testing.T, f func(c *LruCache), expected []string) {
	c := NewLruCache(3)
	f(c)
	if keys := c.RecentlyUsedKeys(); !isSameSlice(keys, expected) {
		t.Errorf("keys expected %v, but got %v", expected, keys)
	}
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
