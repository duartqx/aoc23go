package day1calibrationvalue

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2Chan(p string) <-chan string {

	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < len(p); i++ {
			current := fmt.Sprintf("%c", p[i])

			if _, err := strconv.Atoi(current); err == nil {
				ch <- current
			}

			if i+5 <= len(p) {
				switch p[i : i+5] {
				case "three":
					ch <- "3"
				case "seven":
					ch <- "7"
				case "eight":
					ch <- "8"
				}
			}

			if i+4 <= len(p) {
				switch p[i : i+4] {
				case "four":
					ch <- "4"
				case "five":
					ch <- "5"
				case "nine":
					ch <- "9"
				}
			}

			if i+3 <= len(p) {
				switch p[i : i+3] {
				case "one":
					ch <- "1"
				case "two":
					ch <- "2"
				case "six":
					ch <- "6"
				}
			}

		}
	}()

	return ch
}

func Part2(puzzle <-chan string) (total int) {
	for p := range puzzle {
		var calibrationValues []string

		for i := range Part2Chan(p) {
			calibrationValues = append(calibrationValues, i)
		}

		value, _ := strconv.Atoi(
			strings.Join(
				[]string{
					calibrationValues[0],
					calibrationValues[len(calibrationValues)-1],
				},
				"",
			),
		)

		total += value

	}
	return total
}
