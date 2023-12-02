package day2cubesgame

func getTotal(g string) (total int) {
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
		total += getTotal(g)
	}

	return total
}
