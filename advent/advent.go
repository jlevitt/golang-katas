package advent

import (
	"errors"
)

func Min(values []int) (int, error) {
	if len(values) == 0 {
		return 0, errors.New("values is empty")
	}

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min, nil
}
