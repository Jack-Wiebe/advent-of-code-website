package main

import (
	"advent-of-code-website/solutions/utils"
	"strconv"
	"strings"
)


func Part1() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	var coords [][2]int

	for scanner.Scan(){

		line := scanner.Text()
		//fmt.Println(line)
		strArr:=strings.Split(line, ",")
		var coord [2]int
		coord[0],_ = strconv.Atoi(strArr[0])
		coord[1],_ = strconv.Atoi(strArr[1])

		coords = append(coords, coord)

	}

	largestArea:=0


	for i,r := range(coords){
		for _,nr := range(coords[i+1:]){
			t := calculateArea(r, nr)
			if t > largestArea {
				largestArea=t
			}
		}
	}

	return largestArea, nil
}



func calculateArea(c [2]int, nc [2]int) int {

	dx := utils.AbsInt(nc[0] - c[0]) + 1
	dy := utils.AbsInt(nc[1] - c[1]) + 1

	area := dx * dy

	return area

}


