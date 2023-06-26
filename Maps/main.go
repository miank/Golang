package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	empSalary := map[string]int{}
	fmt.Println(empSalary)

	empSalary["Tom"] = 20000
	empSalary = map[string]int{
		"John": 1000,
		"Sam":  1200,
	}

	fmt.Println(empSalary)

	salary := empSalary["Tom"]
	fmt.Printf("Salary: %d", salary)

	// using make
	emp1 := make(map[string]int)

	// Adding a key value
	emp1["Test"] = 2000
	fmt.Println(emp1)

	// Delete the key
	fmt.Println(emp1)

	// Length of map
	lenOfMap := len(emp1)
	fmt.Println("Length of map", lenOfMap)

	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	for k, v := range sample {
		fmt.Printf("Key :%s value : %s \n", k, v)
	}

	m2 := map[string]int{
		"a": 1,
	}
	fmt.Println(m2)

	// Define a map

	sample1 := map[string]string{
		"a": "x",
		"b": "y",
	}

	fmt.Println("Map ", sample1)

	// iterate through the map
	for x, value := range sample1 {
		fmt.Printf("The value of %s is %s \n", x, value)
	}

	for _, v := range sample1 {
		fmt.Printf("value : %s \n", v)
	}

	// length of a map

	employeeSalary := make(map[string]int)
	employeeSalary["Tom"] = 2000
	employeeSalary["Sam"] = 1200

	lenOfMap1 := len(employeeSalary)
	fmt.Println("Length of the map is ", lenOfMap1)

	// Check if key exist or not
	val, ok := employeeSalary["Tom"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)

	val, ok = employeeSalary["Sam1"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)

	// json to map
	a := make(map[int]string)

	a[1] = "John"

	j, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
	} else {
		fmt.Println("Map to Json ", string(j))
	}

	// Make a map
	var emp map[string]int
	emp["Tom"] = 2000

	empSalary2 := map[string]int{
		"John": 1000,
		"Sam":  2000,
	}

	fmt.Println(empSalary2)

	
}
