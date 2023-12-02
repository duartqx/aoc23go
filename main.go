package main

import (
	day1 "aoc23/day1_calibration_value"
	"aoc23/getdata"
	"log"
)

func main() {

	data, err := getdata.GetInputChannel("./day1_calibration_value/input")
	if err != nil {
		log.Fatal(err)
	}

	// log.Println(day1.Part1(data))
	log.Println(day1.Part2(data))
}
