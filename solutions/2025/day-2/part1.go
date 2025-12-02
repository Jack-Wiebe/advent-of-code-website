package main

import (
	"advent-of-code-website/solutions/utils"
	"strconv"
	"strings"
)

func Part1() (int,error) {

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

			count += sumInvalidRanges(start,end)

    }
	}

	return count, nil

}

func sumInvalidRanges(start int, end int) int {

	sum:=0

	for i := start; i <= end; i++ {

		str := strconv.Itoa(i)
		middle := len(str) / 2
    front := str[:middle]
    back := str[middle:]

		if front == back {
			sum+=i
		}
	}
	return sum

}