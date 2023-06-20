package main

import "fmt"

const (
	OvenTime  = 40
	layerTime = 2
)

func main() {
	// knightIsAwake := false
	knightIsAwake := false
	archerIsAwake := false
	prisonerIsAwake := true
	petDogIsPresent := false
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))
}

func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	} else {
		return true
	}
}

func CanSpy(knightIsAwake bool, archerIsAwake bool, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

func CanSignalPrisoner(archerIsAwake bool, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (petDogIsPresent && !archerIsAwake) || (!petDogIsPresent && prisonerIsAwake && !knightIsAwake && !archerIsAwake)
}
