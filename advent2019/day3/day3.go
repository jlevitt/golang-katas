package day3

import (
	"fmt"
	"github.com/jlevitt/katas/advent2019"
	"github.com/jlevitt/katas/advent2019/day3/wiring"
	"log"
)

func PartOne(path string) {
	lines := advent2019.ReadInputLines(path)
	grid := wiring.NewWireGrid()

	for i, description := range lines {
		wire, err := wiring.NewWire(description)
		if err != nil {
			log.Fatalf("Got error parsing line %v: %v\n", i, err)
		}
		grid.AddWire(wire)
	}

	var intersectionDistances []int
	for _, intersection := range grid.Intersections {
		intersectionDistances = append(intersectionDistances, intersection.Distance())
	}

	shortestDistance, err := advent2019.Min(intersectionDistances)
	if err != nil {
		log.Fatal("No intersections found", err)
	}

	fmt.Printf("Shortest distance: %v\n", shortestDistance)
}

func PartTwo(path string) {
	lines := advent2019.ReadInputLines(path)
	grid := wiring.NewWireGrid()

	wire1, err := wiring.NewWire(lines[0])
	if err != nil {
		log.Fatalf("Got error parsing line 1: %v\n", err)
	}
	grid.AddWire(wire1)

	wire2, err := wiring.NewWire(lines[1])
	if err != nil {
		log.Fatalf("Got error parsing line 2: %v\n", err)
	}
	grid.AddWire(wire2)

	var intersectionDistances []int
	for _, intersection := range grid.Intersections {
		distance := wire1.DistanceAt(intersection) + wire2.DistanceAt(intersection)
		intersectionDistances = append(intersectionDistances, distance)
	}

	shortestDistance, err := advent2019.Min(intersectionDistances)
	if err != nil {
		log.Fatal("No intersections found", err)
	}

	fmt.Printf("Shortest distance: %v\n", shortestDistance)
}
