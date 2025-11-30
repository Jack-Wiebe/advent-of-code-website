package main

import (
	"bufio"
	"fmt"
	"os"
)

func Part2() (int, error){

	file,err := os.Open("input")

	if err != nil{
		fmt.Println("Fail to open file", err)
	}

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	count := 0

	defer file.Close()

	for scanner.Scan(){
		line := scanner.Text()
		str := []rune(line)
		grid = append(grid, str)
	}

	for i := 0; i < len(grid); i++{
		for j := 0; j < len(grid[i]); j++{
			if(grid[i][j] == 'A'){
				if(check_diagonals(grid, i, j)){
					count++
				}
		}
		}
	}
	fmt.Println("word count",count)

	return count, nil
}

func check_diagonals(grid [][]rune, x int, y int) bool {

	if  x+1 >= len(grid){
		return false;
	}
	if  x-1 < 0{
		return false;
	}
	if  y + 1 >= len(grid){
		return false;
	}
	if  y -1 < 0{
		return false;
	}

	check := 0
	if grid[x+1][y+1] == 'M'{
		if grid[x-1][y-1] == 'S'{
			check++
		}
	}
	if grid[x-1][y+1] == 'M'{
		if grid[x+1][y-1] == 'S'{
			check++
		}
	}
	if grid[x+1][y-1] == 'M'{
		if grid[x-1][y+1] == 'S'{
			check++
		}

	}
	if grid[x-1][y-1] == 'M'{
		if grid[x+1][y+1] == 'S'{
			check++
		}
	}

	return check >= 2

}