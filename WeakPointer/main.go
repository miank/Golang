// What is a Weak Pointer?
// A weak pointer allows a reference to an object without preventing that object from being garbage-collected.
// This is particularly useful for avoiding cyclic references or when you want to hold a reference without
// affecting the object's lifecycle.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

type WeakCache struct {
	data sync.Map
}

func (wc *WeakCache) Set(key string, value interface{}) {
	// Set the key-value pair in the map
	wc.data.Store(key, value)

	// Attach a finalizer to remove the key when the value is garbage-collected
	runtime.SetFinalizer(value, func(v interface{}) {
		wc.data.Delete(key)
	})
}

func (wc *WeakCache) Get(key string) (interface{}, bool) {
	return wc.data.Load(key)
}

func main() {
	cache := &WeakCache{}

	// Create an object and store it in the cache
	obj := &struct{ Name string }{Name: "WeakRef"}
	cache.Set("example", obj)

	fmt.Println("Before GC:", cache.Get("example"))

	// Remove the strong reference
	obj = nil

	// Force garbage collection
	runtime.GC()

	fmt.Println("After GC:", cache.Get("example"))
}
