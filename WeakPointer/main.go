// What is a Weak Pointer?
// A weak pointer allows a reference to an object without preventing that object from being garbage-collected.
// This is particularly useful for avoiding cyclic references or when you want to hold a reference without
// affecting the object's lifecycle.

package main

import (
	"fmt"
	"runtime"
)

type MyData struct {
	Name string
}

func main() {
	obj := &MyData{Name: "WeakPointerExample"}

	// Set a finalizer to observe garbage collection
	runtime.SetFinalizer(obj, func(o *MyData) {
		fmt.Println("Object finalized:", o.Name)
	})

	// Remove strong references
	obj = nil

	// Force garbage collection (not recommended in production)
	runtime.GC()
}
