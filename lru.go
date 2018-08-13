package lru

import (
	"container/list"
)

// Cache implements an LRU cache
type Cache interface {
	// Insert inserts the given key and value into the cache
	Insert(key, value interface{}) bool

	// Get looks up and returns the given key and value from the cache. The second
	// argument indicates whether the key was indeed found.
	Get(key interface{}) (interface{}, bool)

	// Delete deletes the given key from the cache and returns the deleted value.
	// The second argument indicates whether the key was found in the cache.
	Delete(key interface{}) (interface{}, bool)

	// Evict evicts the least recently used element from the cache. The second
	// argument indicates whether there was an element to remove or not.
	Evict() (interface{}, bool)

	// Len returns the current number of items in the cache
	Len() int

	// Size returns the maximum size of the cache
	Size() int
}

type cache struct {
	elements map[interface{}]*list.Element
	usage    *list.List
	size     int
}

type element struct {
	key   interface{}
	value interface{}
}

// New returns a new LRU cache with the given size
func New(size int) (c Cache) {
	c = cache{
		elements: make(map[interface{}]*list.Element),
		usage:    list.New(),
		size:     size,
	}

	return c
}

func (c cache) Insert(key, value interface{}) bool {
	_, ok := c.elements[key]
	c.elements[key] = c.usage.PushFront(element{key, value})
	if c.usage.Len() > c.Size() {
		c.Evict()
	}
	return ok
}

func (c cache) Get(key interface{}) (interface{}, bool) {
	el, ok := c.elements[key]
	if !ok {
		return nil, false
	}
	c.usage.MoveToFront(el)
	return el.Value.(element).value, true
}

func (c cache) Delete(key interface{}) (interface{}, bool) {
	el, ok := c.elements[key]
	if !ok {
		return nil, false
	}
	delete(c.elements, el)
	return c.usage.Remove(el).(element).value, true
}

func (c cache) Evict() (interface{}, bool) {
	b := c.usage.Back()
	if b == nil {
		return nil, false
	}
	removed := c.usage.Remove(b)
	delete(c.elements, removed.(element).key)
	return removed.(element).value, true
}

func (c cache) Len() int {
	return len(c.elements)
}

func (c cache) Size() int {
	return c.size
}
