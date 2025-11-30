package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part2() (int, error) {

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening file")
		return -1, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input string
	pattern  := `mul[(](\d+)[,](\d+)[)]|do\(\)|don't\(\)`
	sum := 0
	enabled := true

	for scanner.Scan(){
		input += scanner.Text()
	}

	re, _ := regexp.Compile(pattern)
	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches{

		if(match[0] == "do()"){
			enabled = true
		}else if(match[0] == "don't()"){
			enabled = false
		}else{
			if(enabled){
				num0,_ := strconv.Atoi(match[1])
				num1,_ := strconv.Atoi(match[2])
				sum += (num0 * num1)
			}
		}

	}

	fmt.Println("sum of matches in part 2 is:", sum)
	return sum, nil

}