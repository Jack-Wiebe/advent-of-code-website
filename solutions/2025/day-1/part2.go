package main

import (
	"advent-of-code-website/solutions/utils"
	"strconv"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile()
	if err != nil {
			return -1 , err
	}

	count := 0
	val := 50
	var direction int;

	for scanner.Scan(){

		str := scanner.Text()
		runes := []rune(str)

		dir := string(runes[0])
		amount,_ := strconv.Atoi(string(runes[1:]))

		if dir == "L" {
			direction = -1
		}else{
			direction = 1
		}

		turns := utils.AbsInt(amount / 100)
		count += turns
		delta := utils.Mod(amount, 100)

		new_val := val + (delta * direction)

		if direction == 1 && new_val >= 100{
			count++
		}else if direction == -1 && new_val <= 0 && val != 0{
			count++
		}

		val = utils.Mod(new_val,100)

	}

	return count, nil
}
