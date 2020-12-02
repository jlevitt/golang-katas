package wiring

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

var directions = map[rune]Direction {
	'U': Up,
	'D': Down,
	'L': Left,
	'R': Right,
}

type Segment struct {
	direction Direction
	length int
}

func NewSegment(description string) (Segment, error) {
	descRunes := []rune(description)
	directionRune := descRunes[0]
	lengthRunes := descRunes[1:]

	direction := directions[directionRune]
	length, err := strconv.Atoi(string(lengthRunes))
	if err != nil {
		return Segment{}, errors.Wrapf(err, "error parsing segment: %v", description)
	}

	return Segment{
		direction: direction,
		length: length,
	}, nil
}

type Wire struct {
	segments []Segment
	visited map[Point]int
	length int
}

func NewWire(description string) (*Wire, error) {
	segmentDescs := strings.Split(description, ",")

	wire := Wire{}
	for i, desc := range segmentDescs {
		segment, err := NewSegment(desc)
		if err != nil {
			return nil, errors.Wrapf(err, "error parsing wire at position %v", i)
		}
		wire.segments = append(wire.segments, segment)
	}

	wire.visited = make(map[Point]int)

	return &wire, nil
}

func (w *Wire) Visit(p Point) {
	w.length++

	if _, ok := w.visited[p]; !ok {
		w.visited[p] = w.length
	}
}

func (w *Wire) DistanceAt(p Point) int {
	return w.visited[p]
}