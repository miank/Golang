package main

import "fmt"

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
	employeeSalary["Alex"] = 2000
	employeeSalary["Sam"] = 1200

	for k1, v1 := range employeeSalary {
		fmt.Printf("Employee Name %s and Salary is %d \n", k1, v1)
	}

	lenOfMap := len(employeeSalary)
	fmt.Println("Length of map is ", lenOfMap)

}

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}

	return keys
}
