package main

import (
	"advent-of-code-website/solutions/utils"
	"sort"
	"strconv"
	"strings"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	var ranges []Range
	var union []Range
	var previousRange Range
  count := 0

	for scanner.Scan(){

		line := scanner.Text()

		digits := strings.Split(line, "-")

		if len(digits) > 1{

			start,_:=strconv.Atoi(digits[0])
			end,_:=strconv.Atoi(digits[1])

			ranges = append(ranges, Range{start:start, end: end})

		}

	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	for _,r := range(ranges){

		if previousRange == (Range{}) {
			previousRange = r
		}else if r.start > previousRange.end {
			union = append(union, Range{start:previousRange.start, end:previousRange.end})
			previousRange = r
		}else{
			previousRange = Range{start:previousRange.start, end:max(r.end, previousRange.end)}
		}

	}

	union = append(union, previousRange)

	for _,r:= range(union){
		count+= r.end - r.start + 1
	}

	return count, nil

}