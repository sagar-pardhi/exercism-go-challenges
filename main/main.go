package main

import "fmt"

type Car struct {
	speed        int
	batteryDrain int
}

func NewCar(speed, batteryDrain int) *Car {
	return &Car{
		speed:        speed,
		batteryDrain: batteryDrain,
	}
}

func (car Car) Drive() {
	fmt.Println("car is now Car", car)
}

func main() {
	c := NewCar(5, 8)
	c.Drive()
}
