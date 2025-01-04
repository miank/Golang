package main

import (
	"fmt"
	"runtime"
)

type Object struct {
	reference *Object
}

func main() {
	// Initial setup
	obj1 := &Object{}                // This object will be marked as grey initially
	obj2 := &Object{reference: obj1} // Another object referencing obj1

	// This is where the tri-color algorithm kicks in
	obj1.reference = obj2 // Creating a circular reference

	// Unsetting obj1 and obj2 to make them eligible for GC
	obj1 = nil
	obj2 = nil

	// Forcing a garbage collection cycle
	runtime.GC()

	fmt.Println("Garbage Collection executed")
}
