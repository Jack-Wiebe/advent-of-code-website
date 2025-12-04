package main

import (
	"advent-of-code-website/solutions/utils"
)

var (
	Directions = [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1},           {0, 1},
		{1, -1},  {1, 0},  {1, 1},
	}
)

func Part1() (int,error) {

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

	for i := 0; i < len(grid); i++{
		for j := 0; j < len(grid[i]); j++{

			if grid[i][j] == '@' && checkSurroundingCells(&grid, i, j) < 4 {
				count++
			}

		}
	}

	return count, nil

}

func checkSurroundingCells(grid *[][]rune, x int, y int) int {

	count:=0

	for _, dir := range Directions {
		nx, ny := x+dir[0], y+dir[1]

		if nx >= 0 && nx < len(*grid) && ny >= 0 && ny < len((*grid)[x]) {
				if (*grid)[nx][ny] == '@' {
					count++
				}
		}
	}

	return count
}