package day3gearratios

import (
	"log"
	"regexp"
	"strconv"
)

type ratio struct {
	startIndex int
	endIndex   int
	val        string
}

func newRatio() *ratio {
	return &ratio{}

}

func (r *ratio) reset() {
	r.startIndex = 0
	r.endIndex = 0
	r.val = ""

}

func (r ratio) isDigit(b byte) bool {
	matched, _ := regexp.MatchString(`\d`, string(b))
	return matched
}

func (r ratio) isSpecial(b byte) bool {
	matched, _ := regexp.MatchString(`[^\d.]`, string(b))
	return matched
}

func (r ratio) checkLeftRight(d string) (value int) {
	// Check left and right
	if r.isSpecial(d[r.startIndex]) || r.isSpecial(d[r.endIndex-1]) {
		value, _ = strconv.Atoi(r.val)
	}
	return value
}

func (r ratio) checkAboveBellow(above, bellow *string) (value int) {
	for i := r.startIndex; i <= r.endIndex; i++ {

		// Check above
		if above != nil && r.isSpecial((*above)[i]) {
			value, _ := strconv.Atoi(r.val)
			return value
		}
		// Check bellow
		if bellow != nil && r.isSpecial((*bellow)[i]) {
			value, _ := strconv.Atoi(r.val)
			return value

		}
	}
	return value
}

func getAboveBellow(i int, data *[]string) (above *string, bellow *string) {

	if i > 0 {
		above = &(*data)[i-1]
	}

	if i+1 < len(*data) {
		bellow = &(*data)[i+1]
	}

	return above, bellow
}

func Part1(data *[]string) (total int) {

	for i, d := range *data {

		r := newRatio()

		above, bellow := getAboveBellow(i, data)

		for j := 0; j < len(d); j++ {

			if r.isDigit(d[j]) {
				if r.val == "" {
					if j > 0 {
						r.startIndex = j - 1
					} else {
						r.startIndex = j
					}
				}
				r.val += string(d[j])
			} else if r.val != "" {
				if j+1 < len(d) {
					r.endIndex = j + 1
				} else {
					r.endIndex = j
				}

				if val := r.checkLeftRight(d); val != 0 {
					total += val
				} else if val := r.checkAboveBellow(above, bellow); val != 0 {
					total += val
				}

				r.reset()
			}
		}
	}
	log.Println(total)
	return total
}
