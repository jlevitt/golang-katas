package wiring

import "math"

type DirectionVector struct {
	deltaX, deltaY int
}

var directionVectors = map[Direction]DirectionVector{
	Up: {0, 1},
	Down: {0, -1},
	Left: {-1, 0},
	Right: {1, 0},
}

type Point struct {
	x, y int
}

func (p Point) Distance() int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}

func (p Point) translate(v DirectionVector) Point {
	p.x += v.deltaX
	p.y += v.deltaY

	return p
}

type WireGrid struct {
	wires int
	visited map[Point]int
	Intersections []Point
}

func NewWireGrid() *WireGrid {
	var g WireGrid
	g.visited = make(map[Point]int)

	return &g
}

func (g *WireGrid) AddWire(wire *Wire) {
	cursor := Point{0, 0}
	g.wires++
	wireNumber := g.wires

	for _, segment := range wire.segments {
		v := directionVectors[segment.direction]

		for step := 0; step < segment.length; step++ {
			cursor = cursor.translate(v)

			wire.Visit(cursor)

			if visitedBy, ok := g.visited[cursor]; ok && visitedBy < wireNumber {
				g.Intersections = append(g.Intersections, cursor)
			}
			g.visited[cursor] = wireNumber
		}
	}
}
