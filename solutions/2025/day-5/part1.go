package main

import (
	"advent-of-code-website/solutions/utils"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end int
}

func Part1() (int,error) {

	scanner, err := utils.ReadFile()
	if err != nil {
			return -1 , err
	}

	var ranges []Range
	var ids []int


	count := 0

	for scanner.Scan(){

		line := scanner.Text()
		//fmt.Println(line)

		digits := strings.Split(line, "-")

		if len(digits) == 1{

			id,_ :=strconv.Atoi(digits[0])
			ids = append(ids, id)

		}else{

			start,_:=strconv.Atoi(digits[0])
			end,_:=strconv.Atoi(digits[1])

			ranges = append(ranges, Range{start:start, end:end})

		}

	}

	//fmt.Println(ranges)
	//fmt.Println(ids)

	for _,id := range(ids){
		for _,r := range(ranges){
			if id >= r.start && id <= r.end {
				count++
				break;
			}
		}
	}


	return count, nil

}