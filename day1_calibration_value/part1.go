package day1

import (
	"fmt"
	"strconv"
)

func Part1(puzzle <-chan string) (total int) {

	for p := range puzzle {
		var (
			calibrationValue    string
			calibrationValueInt int
		)

		for _, i := range p {
			iStr := fmt.Sprintf("%c", i)
			_, err := strconv.Atoi(iStr)
			if err == nil {
				calibrationValue += iStr
				break
			}
		}

		for i := len(p) - 1; i >= 0; i-- {
			iStr := fmt.Sprintf("%c", p[i])
			_, err := strconv.Atoi(iStr)
			if err == nil {
				calibrationValue += iStr

				break
			}
		}

		if len(calibrationValue) == 1 {
			calibrationValue += calibrationValue
		}

		calibrationValueInt, _ = strconv.Atoi(calibrationValue)
		total += calibrationValueInt
	}

	return total
}
