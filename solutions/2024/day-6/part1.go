package main

import (
	"bufio"
	"fmt"
	"os"
)

func Part1() (int, error) {

	directions := [][]int{
		{-1,  0},
		{ 0,  1},
		{ 1,  0},
		{ 0, -1},
	}

	current_direction := directions[0]

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("file could not be loaded", err)
		return -1, err
	}

	scanner := bufio.NewScanner(file)

	defer file.Close()

	var input [][]rune

	for scanner.Scan(){

		line := scanner.Text()
		var row []rune

		for _, char := range(line){
			row = append(row, char)
		}

		input = append(input, row)
		//fmt.Println(line)

	}

	output := input
	current_position:=findStartingPos(input)
	done:=false

	for !done {
		current_position, done = moveForward(current_direction, current_position, input, &output)

		current_direction = getNextDir(directions, current_direction)
	}

	count := 0
	for _, row := range output{
		fmt.Println(string(row))
		for _,char := range row{
			if char == '*'{
				count++
			}
		}
	}

	fmt.Println("visited", count, "cells")

	return count, nil

}

func getNextDir(dirs[][]int, dir[]int) []int{

	for ind, val := range dirs{
		if val[0] == dir[0] && val[1] == dir[1]{
			next := (ind + 1) % len(dirs)
			//fmt.Println(next)
			return dirs[next]
		}
	}

	return []int{-1,-1}
}

func findStartingPos(input [][]rune) Position {

	for x, line := range input {
		for y, val := range line {
			if val == '^' {
				//fmt.Println(x, y)
				return Position{x:x, y:y}
			}
		}
	}

	return Position{x:-1,y:-1}

}

func moveForward(dir []int, pos Position, input [][]rune, output *[][]rune) (Position, bool){

	next_position := pos
	//fmt.Println(next_position)
	done := false

	for {
			if next_position.x + dir[0] < 0 || next_position.x + dir[0] >= len(input) || next_position.y + dir[1] < 0 || next_position.y + dir[1] >= len(input[next_position.x]) {
					//fmt.Println("Made it out at", next_position.x, next_position.y)
					done = true
					return next_position, done
			}

			nextX, nextY := next_position.x + dir[0], next_position.y + dir[1]
			if input[nextX][nextY] == '#' {
					//fmt.Println("Obstacle Found at", next_position.x, next_position.y)
					return next_position, done
			}else{
				next_position = Position{x:nextX, y: nextY}
				(*output)[nextX][nextY] = '*'
				//fmt.Println("No obstacle", next_position.x, next_position.y)
			}
	}
}

type Position struct {
	x int
	y int
}