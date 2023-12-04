package day3gearratios

import (
	"regexp"
	"strconv"
)

type ratio struct {
	startIndex int
	endIndex   int
	val        string
	above      *string
	bellow     *string
}

func newRatio() *ratio {
	return &ratio{}
}

func (r *ratio) addValue(s byte, i int) {
	if r.val == "" {
		r.startIndex = i
	}
	r.val += string(s)
}

func (r *ratio) resetValue() {
	r.startIndex = 0
	r.endIndex = 0
	r.val = ""
}

func (r ratio) getStartIndex() int {
	if r.startIndex > 0 {
		return r.startIndex - 1
	}
	return r.startIndex
}

func (r ratio) getEndIndex() int {
	return r.endIndex + 1
}

func (r ratio) valueIsEmpty() bool {
	return r.val == ""
}

func (r ratio) isDigit(b byte) bool {
	matched, _ := regexp.MatchString(`\d`, string(b))
	return matched
}

func (r ratio) isSpecial(b byte) bool {
	matched, _ := regexp.MatchString(`[^\d.]`, string(b))
	return matched
}

func (r ratio) hasSpecial(s string) bool {
	matched, _ := regexp.MatchString(`[^\d.]`, s)
	return matched
}

func (r ratio) checkLeftRight(s *string) (value int) {
	if r.hasSpecial((*s)[r.getStartIndex():r.getEndIndex()]) {
		value, _ = strconv.Atoi(r.val)
	}
	return value
}

func (r ratio) checkAboveBellow() (value int) {

	// Check above
	if r.above != nil && r.hasSpecial((*r.above)[r.getStartIndex():r.getEndIndex()]) {
		value, _ = strconv.Atoi(r.val)
		return value
	}
	// Check bellow
	if r.bellow != nil && r.hasSpecial((*r.bellow)[r.getStartIndex():r.getEndIndex()]) {
		value, _ = strconv.Atoi(r.val)
		return value
	}
	return value
}

func (r ratio) getValue(d *string, i int) (value int) {
	r.endIndex = i

	if value = r.checkLeftRight(d); value != 0 {
		return value
	} else if value = r.checkAboveBellow(); value != 0 {
		return value
	}
	return value
}

func (r *ratio) setAboveBellow(i int, data *[]string) *ratio {

	if i > 0 {
		r.above = &(*data)[i-1]
	}

	if i+1 < len(*data) {
		r.bellow = &(*data)[i+1]
	}

	return r
}

func Part1(data *[]string) (total int) {

	for i, d := range *data {

		r := newRatio().setAboveBellow(i, data)

		for j := 0; j < len(d); j++ {

			if r.isDigit(d[j]) {

				r.addValue(d[j], j)

			} else if !r.valueIsEmpty() {

				total += r.getValue(&d, j)

				r.resetValue()
			}

			if j+1 == len(d) && !r.valueIsEmpty() {
				total += r.getValue(&d, j)
			}
		}
	}
	return total
}
