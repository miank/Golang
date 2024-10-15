package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	RollNo    int
}

func main() {
	// Initialize a map with Person structs as keys and strings as values.
	peopleMap := map[Person]string{
		{FirstName: "Alice", LastName: "Johnson", Age: 32, RollNo: 1}:  "Engineer",
		{FirstName: "Bob", LastName: "Smith", Age: 25, RollNo: 2}:      "Doctor",
		{FirstName: "Carol", LastName: "Williams", Age: 25, RollNo: 3}: "Teacher",
		{FirstName: "Dave", LastName: "Brown", Age: 40, RollNo: 4}:     "Artist",
	}

	keys := make([]Person, 0, len(peopleMap))

	for key := range peopleMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {

		return keys[i].RollNo < keys[j].RollNo
	})

	// Display the sorted map based on the sorted keys.
	for _, key := range keys {
		fmt.Printf("%+v: %s\n", key, peopleMap[key])
	}

}
