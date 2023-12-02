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

func day01Part02Start(p string) string {

	for i := 0; i < len(p); i++ {
		current := fmt.Sprintf("%c", p[i])

		if _, err := strconv.Atoi(current); err == nil {
			return current
		}

		switch current {
		case "o":
			if i+3 < len(p) && p[i:i+3] == "one" {
				return "1"
			}
		case "t":
			if i+3 < len(p) && p[i:i+3] == "two" {
				return "2"
			} else if i+5 < len(p) && p[i:i+5] == "three" {
				return "3"
			}
		case "f":
			if i+4 < len(p) {
				if p[i:i+4] == "four" {
					return "4"
				} else if p[i:i+4] == "five" {
					return "5"
				}
			}
		case "s":
			if i+3 < len(p) && p[i:i+3] == "six" {
				return "6"
			} else if i+5 < len(p) && p[i:i+5] == "seven" {
				return "7"
			}
		case "e":
			if i+5 < len(p) && p[i:i+5] == "eight" {
				return "8"
			}
		case "n":
			if i+4 < len(p) && p[i:i+4] == "nine" {
				return "9"
			}
		}
	}
	return ""
}

func day01Part02End(p string) string {

	for i := len(p) - 1; i >= 0; i-- {
		current := fmt.Sprintf("%c", p[i])

		if _, err := strconv.Atoi(current); err == nil {
			return current
		}

		switch current {
		case "e":
			if i-2 >= 0 && p[i-2:i] == "on" {
				return "1"
			} else if i-4 >= 0 && p[i-4:i] == "thre" {
				return "3"
			} else if i-3 >= 0 && p[i-3:i] == "nin" {
				return "9"
			} else if i-3 >= 0 && p[i-3:i] == "fiv" {
				return "5"
			}
		case "o":
			if i-2 >= 0 && p[i-2:i] == "tw" {
				return "2"
			}
		case "r":
			if i-3 >= 0 && p[i-3:i] == "fou" {
				return "4"
			}
		case "x":
			if i-2 >= 0 && p[i-2:i] == "si" {
				return "6"
			}
		case "t":
			if i-4 >= 0 && p[i-4:i] == "eigh" {
				return "8"
			}
		case "n":
			if i-4 >= 0 && p[i-4:i] == "seve" {
				return "7"
			}
		}
	}
	return ""
}

func day01Part02(puzzle <-chan string) (total int) {
	for p := range puzzle {

		var calibrationValues []string

		if start := day01Part02Start(p); start != "" {
			calibrationValues = append(calibrationValues, start)
		}

		if end := day01Part02End(p); end != "" {
			calibrationValues = append(calibrationValues, end)
		}

		if len(calibrationValues) == 1 {
			calibrationValues = append(calibrationValues, calibrationValues[0])
		}

		value, _ := strconv.Atoi(strings.Join(calibrationValues, ""))
		total += value

	}
	return total
}

func main() {
	data, err := GetInputChannel("./input2")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(day01Part01(data))
	log.Println(day01Part02(data))
}
