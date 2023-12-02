package day2cubesgame

import (
	"strconv"
	"strings"
)

const (
	red   int = 12
	green     = 13
	blue      = 14
)

type CubeGames struct {
}

func (cg CubeGames) getId(s string) (id int) {
	if strings.Contains(s, ":") {
		s = strings.Split(strings.TrimPrefix(s, "Game "), ":")[0]

		if id, err := strconv.Atoi(s); err == nil {
			return id
		}
	}
	return id
}

func (cg CubeGames) getSets(s string) *[][]string {
	sets := [][]string{}
	for _, set := range strings.Split(strings.Split(s, ":")[1], ";") {
		sets = append(sets, strings.Split(strings.TrimSpace(set), ","))
	}
	return &sets
}

func (cg CubeGames) getValueAndColor(vc string) (value int, color string) {
	valueAndColor := strings.Split(strings.TrimSpace(vc), " ")
	value, _ = strconv.Atoi(strings.TrimSpace(valueAndColor[0]))
	return value, strings.TrimSpace(valueAndColor[1])
}

func (cg CubeGames) FromGame(g string) (total int) {
	isPossible := true
	for _, set := range *cg.getSets(g) {
		for _, c := range set {
			value, color := cg.getValueAndColor(c)
			switch {
			case color == "red" && value > red:
				isPossible = false
			case color == "green" && value > green:
				isPossible = false
			case color == "blue" && value > blue:
				isPossible = false
			}
			if !isPossible {
				break
			}
		}
	}
	if isPossible {
		total += cg.getId(g)
	}
	return total
}

func Part1(data <-chan string) (total int) {

	cg := CubeGames{}

	for g := range data {
		total += cg.FromGame(g)
	}

	return total
}
