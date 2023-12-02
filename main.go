package main

import (
	// day1 "aoc23/day1_calibration_value"
	day2 "aoc23/day2_cubes_game"
	"aoc23/getdata"
	"log"
)

func main() {

	data, err := getdata.GetInputChannel("./day2_cubes_game/input")
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(day1.Part1(data))
	// log.Println(day1.Part2(data))
	log.Println(day2.Part1(data))
}
