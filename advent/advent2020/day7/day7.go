package day7

import (
	"fmt"

	"github.com/jlevitt/katas/advent/advent2020/day7/internal/parse"
	"github.com/jlevitt/katas/advent/input"
)

type node struct {
	bagColor     string
	children     []*node
	visited      bool
	containsGold bool
}

func (n *node) walk() {
	if n.visited {
		return
	}

	n.visited = true

	for _, child := range n.children {
		child.walk()
		if child.bagColor == "shiny gold" || child.containsGold {
			n.containsGold = true
			return
		}
	}
}

func PartOne(path string) error {
	lines := input.ReadInputLines(path)
	parsedLines, err := parseLines(lines)
	if err != nil {
		return err
	}

	bagMap := buildBagMap(parsedLines)

	result := 0
	for _, bag := range bagMap {
		bag.walk()
		if bag.containsGold {
			result++
		}
	}

	fmt.Printf("Result: %v\n", result)

	return nil
}

func buildBagMap(lines []parse.Line) map[string]*node {
	result := map[string]*node{}

	for _, line := range lines {
		result[line.BagColor] = &node{
			bagColor: line.BagColor,
			children: []*node{},
		}
	}

	for _, line := range lines {
		parent, ok := result[line.BagColor]
		if !ok {
			parent = &node{
				bagColor: line.BagColor,
				children: []*node{},
			}
			result[line.BagColor] = parent
		}

		for _, child := range line.Children {
			childNode, ok := result[child.BagColor]
			if !ok {
				childNode = &node{
					bagColor: child.BagColor,
					children: []*node{},
				}
				result[child.BagColor] = childNode
			}

			parent.children = append(parent.children, childNode)
		}
	}

	return result
}

func parseLines(lines []string) ([]parse.Line, error) {
	var result []parse.Line

	for i, l := range lines {
		line, err := parse.NewLine(l)
		if err != nil {
			return nil, fmt.Errorf("parsing line %v: %w", i+1, err)
		}

		result = append(result, line)
	}

	return result, nil
}
