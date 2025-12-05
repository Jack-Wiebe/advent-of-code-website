package main

import (
	"advent-of-code-website/solutions/utils"
	"fmt"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile()
	if err != nil {
			return -1 , err
	}

	count := 0
	var grid [][]rune

	for scanner.Scan(){

		line := scanner.Text()
		str := []rune(line)
		grid = append(grid, str)

	}

	output := make([][]rune, len(grid))
	for i, row := range grid {
			output[i] = make([]rune, len(row))
			copy(output[i], row)
	}

	for {
		fmt.Println(utils.GridToString(grid))
		fmt.Println()
		value := removePaperRoll(&grid, &output)
		if value == 0 {
				break
		}
		count+=value

		for i, row := range output {
			grid[i] = make([]rune, len(row))
			copy(grid[i], row)
		}

	}

	return count, nil

}

func removePaperRoll(grid *[][]rune, output *[][]rune)int{
	count:=0
	for i := 0; i < len(*grid); i++{
		for j := 0; j < len((*grid)[i]); j++{

			if (*grid)[i][j] == '@' && checkSurroundingCells(grid, i, j) < 4 {
				(*output)[i][j] = '.'
				count++
			}
		}
	}
	return count
}