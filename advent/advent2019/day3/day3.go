package day3

import (
	"fmt"

	"github.com/jlevitt/katas/advent"
	"github.com/jlevitt/katas/advent/advent2019/day3/wiring"
	"github.com/jlevitt/katas/advent/input"
)

func PartOne(path string) error {
	lines := input.ReadInputLines(path)
	grid := wiring.NewWireGrid()

	for i, description := range lines {
		wire, err := wiring.NewWire(description)
		if err != nil {
			return fmt.Errorf("got error parsing line %v: %w\n", i, err)
		}
		grid.AddWire(wire)
	}

	var intersectionDistances []int
	for _, intersection := range grid.Intersections {
		intersectionDistances = append(intersectionDistances, intersection.Distance())
	}

	shortestDistance, err := advent.Min(intersectionDistances)
	if err != nil {
		fmt.Errorf("no intersections found: %w", err)
	}

	fmt.Printf("Shortest distance: %v\n", shortestDistance)

	return nil
}

func PartTwo(path string) error {
	lines := input.ReadInputLines(path)
	grid := wiring.NewWireGrid()

	wire1, err := wiring.NewWire(lines[0])
	if err != nil {
		return fmt.Errorf("got error parsing line 1: %w\n", err)
	}
	grid.AddWire(wire1)

	wire2, err := wiring.NewWire(lines[1])
	if err != nil {
		return fmt.Errorf("got error parsing line 2: %w\n", err)
	}
	grid.AddWire(wire2)

	var intersectionDistances []int
	for _, intersection := range grid.Intersections {
		distance := wire1.DistanceAt(intersection) + wire2.DistanceAt(intersection)
		intersectionDistances = append(intersectionDistances, distance)
	}

	shortestDistance, err := advent.Min(intersectionDistances)
	if err != nil {
		return fmt.Errorf("no intersections found: %w", err)
	}

	fmt.Printf("Shortest distance: %v\n", shortestDistance)

	return nil
}
