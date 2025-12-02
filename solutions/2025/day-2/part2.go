package main

import (
	"advent-of-code-website/solutions/utils"
	"strconv"
	"strings"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile()
	if err != nil {
			return -1 , err
	}

	count := 0

	for scanner.Scan(){

		str := scanner.Text()
		ranges := strings.Split(str, ",")

		for _, r := range ranges {

			str := strings.Split(r, "-")

			startStr := str[0]
			endStr := str[1]

			start,_:=strconv.Atoi(startStr)
			end,_:=strconv.Atoi(endStr)

			count += sumInvalidRanges2(start,end)

    }
	}

	return count, nil

}

func sumInvalidRanges2(start int, end int) int {

	sum:=0

	for i := start; i <= end; i++ {

		str := strconv.Itoa(i)

		if checkRepeating(str) {
			sum+=i
		}

	}
	return sum

}

func checkRepeating(str string) bool {

	for i := 1; i <= len(str)/2; i++ {
		if len(str)%i == 0 { // check if substring fits evenly in string
			pattern := str[:i]
			valid := true

			for j := i; j < len(str); j += i {
				if str[j:j+i] != pattern {
					valid = false
					break
				}
			}

			if valid {
				return true
			}
		}
	}

	return false
}