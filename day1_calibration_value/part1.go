package day1

import (
	"fmt"
	"strconv"
)

func Part1(puzzle <-chan string) (total int) {

	for p := range puzzle {
		var (
			calValues    []string
			calValuesInt int
		)

		for _, s := range p {
			str := fmt.Sprintf("%c", s)
			if _, err := strconv.Atoi(str); err == nil {
				calValues = append(calValues, str)
			}
		}

		calValuesInt, _ = strconv.Atoi(calValues[0] + calValues[len(calValues)-1])
		total += calValuesInt
	}

	return total
}
