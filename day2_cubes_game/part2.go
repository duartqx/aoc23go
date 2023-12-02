package day2cubesgame

type cubeGame struct {
	red   int
	green int
	blue  int
}

func (cg *cubeGame) SetRed(value int) *cubeGame {
	if cg.red < value {
		cg.red = value
	}
	return cg
}

func (cg *cubeGame) SetGreen(value int) *cubeGame {
	if cg.green < value {
		cg.green = value
	}
	return cg
}

func (cg *cubeGame) SetBlue(value int) *cubeGame {
	if cg.blue < value {
		cg.blue = value
	}
	return cg
}

func (cg cubeGame) GetPower() int {
	return cg.red * cg.green * cg.blue
}

func (cg *cubeGame) Play(g string) int {
	for _, set := range *getGameSets(g) {
		// [][]string{
		// 	{"4 blue", "3 green", "7 red"},
		// 	{"4 blue", "3 green", "7 red"},
		// }

		var red, green, blue int

		for _, s := range set {

			// "4 blue" ... "7 red"

			value, color := getValueAndColor(s)
			switch color {
			case "red":
				red += value
			case "green":
				green += value
			case "blue":
				blue += value
			}
		}

		cg.SetRed(red).SetGreen(green).SetBlue(blue)

	}

	return cg.GetPower()
}

func Part2(data <-chan string) (total int) {

	for g := range data {
		cg := cubeGame{}
		total += cg.Play(g)
	}

	return total
}
