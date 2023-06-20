package main

import (
	"fmt"
)

const ()

func main() {
	fmt.Println(CalculateCost(21))
}

func CalculateWorkingCarsPerHour(cars, succesrate float32) float32 {
	return (cars * succesrate / 100)
}

func CalculateWorkingCarsPerMinute(cars, succesrate int) int {
	return cars * succesrate / 6000
}

func CalculateCost(cars int) int {
	a := cars % 10
	b := int(cars / 10)

	price := (a * 10000) + (b * 95000)

	return price
}
