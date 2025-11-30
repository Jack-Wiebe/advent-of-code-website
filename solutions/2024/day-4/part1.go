package main

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() (int, error) {

	directions := [8][2]int {
		{0,1},  //N
		{1,1},  //NE
		{1,0},  //E
		{1,-1}, //SE
		{0,-1}, //S
		{-1,1}, //SW
		{-1,0}, //W
		{-1,-1},//NW
	}

	file,err := os.Open("input")

	if err != nil{
		fmt.Println("Fail to open file", err)
	}

	scanner := bufio.NewScanner(file)
	var grid [][]rune
	word := "MAS"
	count := 0

	defer file.Close()

	for scanner.Scan(){
		line := scanner.Text()
		//fmt.Println(line)
		str := []rune(line)
		grid = append(grid, str)
	}

	for i := 0; i < len(grid); i++{
		for j := 0; j < len(grid[i]); j++{
			if(grid[i][j] == 'X'){
				//fmt.Println()
				//fmt.Println("START",i,j)

				for _, dir := range directions{
					if(check_direction(grid, dir, i, j, word)){
						//fmt.Println("Found XMAS!!!!")
						count++
					}
				}
			}
		}
	}

	//fmt.Println()
	fmt.Println("word count",count)

	return count, nil

}

func check_direction(grid [][]rune, dir [2]int, x int, y int, word string) bool {

	//fmt.Println("Check Direction", dir)

	if  x + (dir[0] * len(word)) >= len(grid){
		//fmt.Println("out of bounds bottom")
		return false;
	}
	if  x + (dir[0] * len(word)) < 0{
		//fmt.Println("out of bounds top")
		return false;
	}
	if  y + (dir[1] * len(word)) >= len(grid){
		//fmt.Println("out of bounds right")
		return false;
	}
	if  y + (dir[1] * len(word)) < 0{
		//fmt.Println("out of bounds left")
		return false;
	}

	found := false

	for i, val := range word {
		if grid[x+(dir[0] * (i+1))][y+(dir[1] * (i+1))] == val{
			//fmt.Printf("found letter %c\n", val)
			if(i >= len(word)-1){
				found = true
			}
		}else{
			break
		}
	}

	return found
}