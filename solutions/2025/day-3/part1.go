package main

import (
	"advent-of-code-website/solutions/utils"
	"strconv"
)

func Part1() (int,error) {

	scanner, err := utils.ReadFile()
	if err != nil {
			return -1 , err
	}

	count := 0

	for scanner.Scan(){

		str := scanner.Text()

		firstDigit:=0
		secondDigit:=0
		index:=0

		for i,n := range(str){
			if i == len(str)-1{
				break
			}
			test,_ := strconv.Atoi(string(n))
			if test > firstDigit{
				firstDigit = test
				index = i
			}
		}

		for _,n := range(str[index+1:]){
			test,_ := strconv.Atoi(string(n))
			if test > secondDigit{
				secondDigit = test
			}
		}

		output,_ := strconv.Atoi(strconv.Itoa(firstDigit)+strconv.Itoa(secondDigit))
		count+=output

	}

	return count, nil

}