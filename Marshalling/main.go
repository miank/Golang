package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Address string
}

func main() {

	emp := Employee{
		Name:    "George Smith",
		Age:     30,
		Address: "Newyork USA",
	}

	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(empData))
}
