package main

import (
	"advent-of-code-website/solutions/utils"
	"fmt"
)


func Part2() (int,error) {

	scanner, err := utils.ReadFile("test")
	if err != nil {
			return -1 , err
	}

	count:=0
	mergeCount:=1
	merge:=false
	var tachyonIndex []int

	for scanner.Scan(){

		line := scanner.Text()
		fmt.Println(line)
		t := tachyonIndex

		merge=false
		for i,r := range(line){

			if r == 'S' {
				tachyonIndex = append(tachyonIndex, i)
			}else if (r == '^' && utils.Contains(t, i)) {
				tachyonIndex = utils.RemoveElement(tachyonIndex, utils.IndexOf(tachyonIndex,i))
				if utils.IndexOf(tachyonIndex, i-1) == -1 && i-1 >= 0 {
					tachyonIndex = append(tachyonIndex, i-1)
					count++
				}else {
					merge=true
				}
				if utils.IndexOf(tachyonIndex, i+1) == -1 && i+1 <= len(line) {
					tachyonIndex = append(tachyonIndex, i+1)
					count++
				}else {
					merge=true
				}
			}

		}
		if merge {
			mergeCount++
		}

	}
	fmt.Println(mergeCount, count)
	return count, nil
}


