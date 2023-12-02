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
