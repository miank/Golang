package main

import (
	"fmt"
	"math"
)

func main() {
	res := math.Ceil(0.75)
	fmt.Println(res)

	res = math.Ceil(-1.6)
	fmt.Println(res)

	res = math.Ceil(1)
	fmt.Println(res)

	res1 := math.Floor(1.6)
	fmt.Println(res1)

	res1 = math.Floor(-1.6)
	fmt.Println(res1)

	res1 = math.Floor(1)
	fmt.Println(res1)

	// Get integer value of a float
	res2 := math.Trunc(1.6)
	fmt.Println(res2)

	res2 = math.Trunc(-1.6)
	fmt.Println(res2)

	res2 = math.Trunc(1.6)
	fmt.Println(res2)

}
