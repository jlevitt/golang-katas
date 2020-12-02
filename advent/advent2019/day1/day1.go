package day1

import (
	"fmt"
	"strconv"

	"github.com/jlevitt/katas/advent/input"
)

func PartTwo(path string) error {
	moduleMasses := input.ReadInputLines(path)

	totalFuelRequired := 0
	for _, massStr := range moduleMasses {
		mass, err := strconv.Atoi(massStr)
		if err != nil {
			return err
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

	return nil
}

func FuelRequired(moduleMass int) int {
	return moduleMass/3 - 2
}
