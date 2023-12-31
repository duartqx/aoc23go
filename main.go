package main

import (
	// day1 "aoc23/day1_calibration_value"
	// day2 "aoc23/day2_cubes_game"
	day3 "aoc23/day3_gear_ratios"
	"aoc23/getdata"
	"log"
)

func main() {

	data, err := getdata.GetInputSlice("./day3_gear_ratios/input")
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(day1.Part1(data))
	// log.Println(day1.Part2(data))
	// log.Println(day2.Part1(data))
	// log.Println(day2.Part2(data))
	log.Println(day3.Part1(data)) // (532422 *wrong) | (535235 right)
}
