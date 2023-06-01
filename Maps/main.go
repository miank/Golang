package main

import "fmt"

type employee struct {
	name   string
	salary string
}

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
}
