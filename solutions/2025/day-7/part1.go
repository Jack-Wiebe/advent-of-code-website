package main

import (
	"advent-of-code-website/solutions/utils"
)


func Part1() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	count:=0
	var tachyonIndex []int

	for scanner.Scan(){

		line := scanner.Text()
		//fmt.Println(line)
		t := tachyonIndex
		for i,r := range(line){
			if r == 'S' {
				tachyonIndex = append(tachyonIndex, i)
			}else if (r == '^' && utils.Contains(t, i)) {
				tachyonIndex = utils.RemoveElement(tachyonIndex, utils.IndexOf(tachyonIndex,i))
				if utils.IndexOf(tachyonIndex, i-1) == -1 && i-1 >= 0 {
					tachyonIndex = append(tachyonIndex, i-1)
				}
				if utils.IndexOf(tachyonIndex, i+1) == -1 && i+1 <= len(line) {
					tachyonIndex = append(tachyonIndex, i+1)
				}
				count++
			}
		}

	}

	return count, nil
}


