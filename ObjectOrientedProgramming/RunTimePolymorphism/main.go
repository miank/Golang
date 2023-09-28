package main

import "fmt"

type taxSystem interface {
	calculateTax()
}

type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}
func main() {

	indianTax := indianTax{
		taxPercentage: 30,
		income:        1000,
	}

	totalIndiaTax := indianTax.calculateTax()
	fmt.Println(totalIndiaTax)

	singaporeTax := &singaporeTax{
		taxPercentage: 10,
		income:        2000,
	}

	totalSingaporeTax := singaporeTax.calculateTax()
	fmt.Println(totalSingaporeTax)
}

// func calculateTotalTax(taxSystems []taxSystem) int {
// 	totalTax := 0
// 	for _, t := range taxSystems {
// 		totalTax += t.calculateTax() //This is where runtime polymorphism happens
// 	}
// 	return totalTax
// }
