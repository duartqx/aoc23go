package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day01Part01(puzzle <-chan string) (total int) {

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

func day01Part02Chan(p string) <-chan string {

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

func day01Part02(puzzle <-chan string) (total int) {
	for p := range puzzle {
		var calibrationValues []string

		for i := range day01Part02Chan(p) {
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

func main() {
	data, err := GetInputChannel("./input")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(day01Part01(data))
	log.Println(day01Part02(data))
}
