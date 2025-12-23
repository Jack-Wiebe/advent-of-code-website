package main

import (
	"advent-of-code-website/solutions/utils"
	"regexp"
	"strconv"
	"strings"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

  count := 0
	var grid [][]string

	for scanner.Scan(){

		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))

	}

	re := regexp.MustCompile(`[*+]`)
	re2 := regexp.MustCompile(`\d`)
	problem:=Problem{}

	for i:=0; i < len(grid[0]); i++{
		number:=""
		for j:=len(grid)-1; j >= 0; j--{
			symbol:=grid[j][i]
			if re2.MatchString(symbol){
				number = symbol + number
			}else if re.MatchString(symbol){
				if problem.operand != "" {
					count+=solveProblem(problem)
					problem=Problem{}
				}
				problem.operand = symbol
			}
		}
		if number != "" {
			problem.numbers = append(problem.numbers, number)
		}
	}
	count+=solveProblem(problem)

	return count, nil

}

func solveProblem(problem Problem)int{
	count:=0
	if problem.operand == "*"{
		count++
	}
	for _,r := range(problem.numbers){
			val,_:=strconv.Atoi(r)
			switch problem.operand {
				case "+":
					count+=val
				case "*":
					count = count * val
			}
		}

	return count
}

const (
	MULTIPLY = "*"
	ADD			 = "+"
)

type Problem struct{

	operand string
	numbers [] string

}