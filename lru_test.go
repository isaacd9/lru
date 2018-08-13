package lru

import "testing"

func TestLookup(t *testing.T) {
	tests := make(map[interface{}]interface{})

	k := "hi"
	v := "bye"

	tests[1] = 1
	tests[200] = 300
	tests[k] = v
	tests[struct {
		test string
		num  int
	}{k, 100}] = v

	c := New(10)
	for k, v := range tests {
		c.Insert(k, v)
		val, _ := c.Get(k)
		if v != val {
			t.Fatalf("Expected %v for key %v. Got %v.", v, k, val)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := make(map[interface{}]interface{})

	k := "hi"
	v := "bye"

	tests[1] = 1
	tests[200] = 300
	tests[k] = v
	tests[struct {
		test string
		num  int
	}{k, 100}] = v

	c := New(10)
	for k, v := range tests {
		c.Insert(k, v)
		val, _ := c.Delete(k)
		if v != val {
			t.Fatalf("Expected %v for key %v. Got %v.", v, k, val)
		}
	}
}

func TestLRUOne(t *testing.T) {
	c := New(3)

	tests := make(map[interface{}]interface{})

	tests["hi"] = "bye"
	tests[100] = 200
	tests[300] = 400

	for k, v := range tests {
		c.Insert(k, v)
		val, _ := c.Get(k)
		if v != val {
			t.Fatalf("Expected %v for key %v. Got %v.", v, k, val)
		}
	}

	c.Get(300)
	c.Get(100)

	delete(tests, "hi")
	c.Insert("new", "test")
	tests["new"] = "test"

	for k, v := range tests {
		val, _ := c.Get(k)
		if v != val {
			t.Fatalf("Expected %v for key %v. Got %v.", v, k, val)
		}
	}
}

func TestLRUTwo(t *testing.T) {
	c := New(10)

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
	if _, ok := c.Get(1); ok {
		t.Fatalf("1 should not be in the cache anymore!")
	}

	// Definitely there!
	if _, ok := c.Get(10); !ok {
		t.Fatalf("10 should now be in the cache")
	}

	if _, ok := c.Get(4); !ok {
		t.Fatalf("4 should still be in the cache")
	}

	if c.Size() != 10 {
		t.Fatalf("The cache should be size 10!")
	}

	if c.Len() != 10 {
		t.Fatalf("The cache should be len 10!")
	}
}
