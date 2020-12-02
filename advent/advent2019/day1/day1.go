package day1

import (
	"fmt"
	"log"
	"strconv"
	"github.com/jlevitt/katas/advent"
)

func PartTwo(path string) {
	moduleMasses := advent.ReadInputLines(path)

	totalFuelRequired := 0
	for _, massStr := range moduleMasses {
		mass, err := strconv.Atoi(massStr)
		if err != nil {
			log.Fatal(err)
		}

		fuelRequired := FuelRequired(mass)
		totalFuelRequired += fuelRequired

		extraFuel := FuelRequired(fuelRequired)
		for extraFuel > 0 {
			totalFuelRequired += extraFuel
			extraFuel = FuelRequired(extraFuel)
		}
	}

	fmt.Printf("Total fuel required: %v\n", totalFuelRequired)
}

func FuelRequired(moduleMass int) int {
	return moduleMass / 3 - 2
}