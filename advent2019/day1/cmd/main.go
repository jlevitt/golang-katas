package main

import (
	"fmt"
	"github.com/jlevitt/katas/advent2019"
	"github.com/jlevitt/katas/advent2019/day1"
)

func main() {
	moduleMasses := advent2019.ReadInput()

	totalFuelRequired := 0
	for _, mass := range moduleMasses {
		fuelRequired := day1.FuelRequired(mass)
		totalFuelRequired += fuelRequired

		extraFuel := day1.FuelRequired(fuelRequired)
		for extraFuel > 0 {
			totalFuelRequired += extraFuel
			extraFuel = day1.FuelRequired(extraFuel)
		}
	}

	fmt.Printf("Total fuel required: %v\n", totalFuelRequired)
}
