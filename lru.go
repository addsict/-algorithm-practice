package lru

import (
	"fmt"
	"strings"
)

type LruCache struct {
	cache    map[string]*Entry
	capacity uint
	head     *Entry
	tail     *Entry
}

type Entry struct {
	key   string
	value interface{}
	prev  *Entry
	next  *Entry
}

func (e *Entry) String() string {
	// return fmt.Sprintf("Entry{value: %v, prev: %v, next: %v}", e.value, e.prev, e.next)
	return fmt.Sprintf("Entry{key:%s, value: %v}", e.key, e.value)
}

func NewLruCache(capacity uint) *LruCache {
	return &LruCache{
		cache:    make(map[string]*Entry, capacity),
		capacity: capacity,
	}
}

// TODO: goroutine safe
func (c *LruCache) Set(key string, value interface{}) {
	e := &Entry{
		key:   key,
		value: value,
	}

	if c.head != nil {
		c.Remove(key)
		head := c.head
		head.prev = e
		e.next = head
	}

	if len(c.cache) == int(c.capacity) {
		// eviction
		tail := c.tail
		delete(c.cache, tail.key)
		c.tail = tail.prev
		tail.prev.next = nil
	}

	c.cache[key] = e
	c.head = e
	if c.tail == nil {
		c.tail = e
	}
}

func (c *LruCache) Remove(key string) {
	if c.cache[key] == nil {
		return
	}

	e := c.cache[key]
	delete(c.cache, key)

	if c.head == e && c.tail == e {
		c.head = nil
		c.tail = nil
	} else if c.head == e {
		c.head = e.next
		e.next.prev = nil
	} else if c.tail == e {
		c.tail = e.prev
		e.prev.next = nil
	} else {
		prev := e.prev
		next := e.next
		prev.next = next
		next.prev = prev
	}
}

func (c *LruCache) GetString(key string) (string, bool) {
	e := c.cache[key]
	if e != nil {
		if str, ok := e.value.(string); ok {
			c.Set(key, e.value) // for updating internal structure
			return str, true
		}
	}
	return "", false
}

func (c *LruCache) GetInt(key string) (int, bool) {
	e := c.cache[key]
	if e != nil {
		if i, ok := e.value.(int); ok {
			c.Set(key, e.value) // for updating internal structure
			return i, true
		}
	}
	return 0, false
}

func (c *LruCache) RecentlyUsedKeys() []string {
	keys := make([]string, 0, len(c.cache))
	for e := c.head; e != nil; e = e.next {
		keys = append(keys, e.key)
	}
	return keys
}

func (c *LruCache) String() string {
	// return fmt.Sprintf("cache: %s, capacity: %d, size: %d, keys: %s", c.cache, c.capacity, c.size, strings.Join(keys, " -> "))
	return fmt.Sprintf("capacity: %d, size: %d, keys: %s", c.capacity, len(c.cache), strings.Join(c.RecentlyUsedKeys(), " -> "))
}
