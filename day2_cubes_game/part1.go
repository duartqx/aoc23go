package day2cubesgame

import (
	"strconv"
	"strings"
)

func getId(s string) (id int) {
	if strings.Contains(s, ":") {
		s = strings.Split(strings.TrimPrefix(s, "Game "), ":")[0]

		if id, err := strconv.Atoi(s); err == nil {
			return id
		}
	}
	return id
}

func getGameSets(s string) *[][]string {
	sets := [][]string{}
	for _, set := range strings.Split(strings.Split(s, ":")[1], ";") {
		sets = append(sets, strings.Split(strings.TrimSpace(set), ","))
	}
	return &sets
}

func getValueAndColor(vc string) (value int, color string) {
	valueAndColor := strings.Split(strings.TrimSpace(vc), " ")
	value, _ = strconv.Atoi(strings.TrimSpace(valueAndColor[0]))
	return value, strings.TrimSpace(valueAndColor[1])
}

func FromGamePart1(g string) (total int) {
	isPossible := true
	for _, set := range *getGameSets(g) {
		for _, c := range set {
			value, color := getValueAndColor(c)
			switch {
			case color == "red" && value > 12:
				isPossible = false
			case color == "green" && value > 13:
				isPossible = false
			case color == "blue" && value > 14:
				isPossible = false
			}
			if !isPossible {
				break
			}
		}
	}
	if isPossible {
		total += getId(g)
	}
	return total
}

func Part1(data <-chan string) (total int) {

	for g := range data {
		total += FromGamePart1(g)
	}

	return total
}
