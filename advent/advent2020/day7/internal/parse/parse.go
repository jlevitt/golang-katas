package parse

import (
	"regexp"
	"strconv"
	"strings"
)

type ChildBag struct {
	BagColor string
	Count    int
}

type Line struct {
	BagColor string
	Children []ChildBag
}

var pattern *regexp.Regexp = regexp.MustCompile("([[:digit:]])+ (.*) bag")

func NewLine(line string) (Line, error) {
	parentChildSplit := strings.Split(line, " bags contain ")

	parent := parentChildSplit[0]
	children := strings.Trim(parentChildSplit[1], ".")

	result := Line{
		BagColor: parent,
	}

	if children == "no other bags" {
		return result, nil
	}

	childBags := strings.Split(children, ",")
	for _, bag := range childBags {
		matches := pattern.FindStringSubmatch(bag)

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			return result, err
		}

		child := ChildBag{
			BagColor: matches[2],
			Count:    count,
		}
		result.Children = append(result.Children, child)
	}

	return result, nil
}
