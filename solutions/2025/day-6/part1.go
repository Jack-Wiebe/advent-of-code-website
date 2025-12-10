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

	var operations [][]string
	count:=0

	for scanner.Scan(){

		line := scanner.Text()

		digits := strings.Fields(line)

		for i,r := range(digits){
			if len(operations)  == 0 {
				operations = make([][]string, len(digits))
			}
			operations[i] = append(operations[i],r)
		}

	}

	for i:=0; i < len(operations); i++{
		total,_:=strconv.Atoi(operations[i][0])
		operation := operations[i][len(operations[i])-1]
		for j:=1; j < len(operations[i]) - 1; j++{
			val,_:=strconv.Atoi(operations[i][j])
			switch operation {
				case "+":
					total+=val
				case "*":
					total = total * val
			}
		}
		count+=total
	}

	return count, nil
}


