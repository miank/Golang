package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// creating a map
	sample := map[string]string{
		"a": "x",
		"b": "y",
	}

	fmt.Println(sample)
	for k, v := range sample {
		fmt.Printf("Key : %s value : %s\n", k, v)
	}

	// Looping throught keys
	for k := range sample {
		fmt.Printf("Key :%s\n", k)
	}

	// Looping throught values
	for _, v := range sample {
		fmt.Printf("value: %s\n", v)
	}

	keys := getAllKeys(sample)
	fmt.Println(keys)

	// length of the map
	employeeSalary := make(map[string]int)
	employeeSalary["Tom"] = 2000
	employeeSalary["Sam"] = 1200

	for k1, v1 := range employeeSalary {
		fmt.Printf("Employee Name %s and Salary is %d \n", k1, v1)
	}

	lenOfMap := len(employeeSalary)
	fmt.Println("Length of map is ", lenOfMap)

	// Check If key exist
	fmt.Println("Key exists case")
	val, ok := employeeSalary["Tom"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)

	fmt.Println("Key doesn't exists case")

	val, ok = employeeSalary["Sam"]
	fmt.Printf("Val: %d, ok: %t\n", val, ok)

	m1 := map[string]int{}
	fmt.Println(m1)

	m2 := make(map[string]string)
	fmt.Println(m2)

	// Creating a nil map
	var employeeSalary1 map[string]int
	fmt.Println(employeeSalary1)
	// This will panic as we are trying to create
	// employeeSalary1["Tom"] = 2000

	// Maps to JSON
	a := make(map[int]string)
	a[1] = "John"
	j, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error : %s", err.Error())
	} else {
		fmt.Println(string(j))
	}

	// Unmarshall json to map
	var b map[int]string
	json.Unmarshal(j, &b)
	fmt.Println(b)

}

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}

	return keys
}
