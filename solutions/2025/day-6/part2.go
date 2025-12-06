package main

import (
	"advent-of-code-website/solutions/utils"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

  count := 0

	for scanner.Scan(){

		//line := scanner.Text()


	}

	return count, nil

}