package main

import (
	"fmt"
	"github.com/jlevitt/katas/advent2019"
	"github.com/jlevitt/katas/advent2019/day1"
	"log"
	"strconv"
)

func main() {
	moduleMasses := advent2019.ReadInputLines()

	totalFuelRequired := 0
	for _, massStr := range moduleMasses {
		mass, err := strconv.Atoi(massStr)
		if err != nil {
			log.Fatal(err)
		}

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
