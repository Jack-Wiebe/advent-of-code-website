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
	val := 50

	for scanner.Scan(){

		str := scanner.Text()
		runes := []rune(str)

		dir := string(runes[0])
		amount,_ := strconv.Atoi(string(runes[1:]))

		if dir == "L" {
			val += (amount * -1)
		}else{
			val += amount
		}

		val = utils.Mod(val, 100)

		if val == 0 {
			count++
		}
	}

	return count, nil

}