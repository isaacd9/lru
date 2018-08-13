# LRU

[![Go Report Card](https://goreportcard.com/badge/github.com/isaacd9/lru)](https://goreportcard.com/report/github.com/isaacd9/lru)

[godoc](https://godoc.org/github.com/isaacd9/lru)

LRU implements the world's simplest LRU cache for Go.

Under the hood this uses the `container/list` package to keep track of each
element in the cache and a map that hashes to a location int the list. Each
time an element is accessed, it's brought to the front of the list. When the
length of the cache exceeds the created size, the last element in the list is
evicted.

## Installation
To download and install simply run:

`go get github.com/isaacd9/lru`

## Supported Methods
The cache implements the following methods. They mostly do what you'd expect.

```go
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
```

## Example!!!
```go
import "github.com/isaacd9/lru"

func main() {
  c := lru.New(10)

  for i := 0; i < 10; i++ {
    c.Insert(i, "test")
  }

  for i := 1; i < 10; i++ {
    c.Get(i)
  }

  c.Insert(10, "cool")
  c.Insert(11, "sup")
  c.Insert(12, "testing")

  // Not found!
  _, ok := c.Get(1)

  // Definitely there!
  c.Get(10)
  c.Get(4)
}
```
