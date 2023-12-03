package day3gearratios

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type checker struct {
	start *regexp.Regexp
	end   *regexp.Regexp
}

func newChecker() *checker {
	return &checker{
		start: regexp.MustCompile(`^[^\d].*`),
		end:   regexp.MustCompile(`.*[^\d]$`),
	}
}

func (ch checker) checkLeftRight(s string) (value int) {
	var (
		currentInt int
	)

	if s == "" {
		return value
	}

	// Check left and right
	if ch.start.MatchString(s) {
		currentInt, _ = strconv.Atoi(s[1:])
		value += currentInt
	} else if ch.end.MatchString(s) {
		currentInt, _ = strconv.Atoi(s[:len(s)-1])
		value += currentInt
	}
	return value
}

func (ch checker) getIndexes(d, s string) (int, int, error) {
	startIndex := strings.Index(d, s)
	if startIndex == -1 {
		return 0, 0, errors.New("Not found")
	}
	endIndex := startIndex + len(s)
	if endIndex < len(d) {
		endIndex++
	}
	if startIndex != 0 {
		startIndex--
	}
	return startIndex, endIndex, nil
}

func (ch checker) check(m, s string) (val int) {
	if m != "." && ch.start.MatchString(m) {
		val, _ = strconv.Atoi(s)
	}
	return val
}

func Part1(data <-chan string) (total int) {

	ch := newChecker()

	var (
		slicedData []string
	)

	for d := range data {
		slicedData = append(slicedData, d)
	}

	for i, d := range slicedData {

		dotSliced := strings.Split(d, ".")

		for _, s := range dotSliced {

			hasLeftRight := ch.checkLeftRight(s)
			if hasLeftRight != 0 {
				total += hasLeftRight
				continue
			}

			startIndex, endIndex, err := ch.getIndexes(d, s)
			if err != nil {
				continue
			}
			for k := startIndex; k < endIndex; k++ {
				// Check above
				if i != 0 {
					if val := ch.check(string(slicedData[i-1][k]), s); val != 0 {
						total += val
						break
					}
				}

				// Check bellow
				if i < (len(slicedData) - 1) {
					if val := ch.check(string(slicedData[i+1][k]), s); val != 0 {
						total += val
						break
					}
				}
			}
		}

	}
	// [467  114  ]
	// [   *      ]
	// [  35  633 ]
	// [      #   ]
	// [617*      ]
	// [     + 58 ]
	// [  592     ]
	// [      755 ]
	// [   $ *    ]
	// [ 664 598  ]
	log.Println(total)
	return total
}
