package main

import "fmt"

// Vechile is an interface that all the vechile types must implement

type Vechile interface {
	Drive() string
}

type car struct{}

func (c car) Drive() string {
	return "Driving a car !!"
}

type bike struct{}

func (b bike) Drive() string {
	return "Driving a bike"
}

func NewVehicle(vechileType string) Vechile {
	if vechileType == "car" {
		return car{}
	} else if vechileType == "bike" {
		return bike{}
	}
	return nil
}

func main() {
	car := NewVehicle("car")
	fmt.Println(car.Drive())

	bike := NewVehicle("bike")
	fmt.Println(bike.Drive())

	// Handling an unknown vehicle type
	unknown := NewVehicle("truck")
	if unknown == nil {
		fmt.Println("Unknown vehicle type!")
	}
}
