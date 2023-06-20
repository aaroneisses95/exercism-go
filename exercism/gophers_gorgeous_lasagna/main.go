package main

import "fmt"

const (
	OvenTime  = 40
	layerTime = 2
)

func main() {
	fmt.Println(ElapsedTime(3, 20))
}

func RemainingOvenTime(actual int) int {
	return OvenTime - actual

}

func PreparationTime(numberOfLayers int) int {
	return layerTime * numberOfLayers
}

func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
