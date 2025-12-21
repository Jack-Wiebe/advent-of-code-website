package main

import (
	"advent-of-code-website/solutions/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)


func Part1() (int,error) {

	scanner, err := utils.ReadFile("test")
	if err != nil {
			return -1 , err
	}

	count:=0
	var lights [][]rune
	var buttons [][][]int
	//var joltages [][]int ignore for now

	for scanner.Scan(){

		line := scanner.Text()
		fmt.Println(line)
		arr := strings.Split(line, " ")
		lights = append(lights, []rune(strings.Trim(arr[0], "[]")) )

		re := regexp.MustCompile(`\(([^)]+)\)`)
    matches := re.FindAllStringSubmatch(line, -1)
    for _, match := range matches {
			parts := strings.Split(strings.TrimSpace(match[1]), ",")
			var numbers []int
			for _, part := range parts {
				num, _ := strconv.Atoi(strings.TrimSpace(part))
				numbers = append(numbers, num)
			}
			buttons = append(buttons, [][]int{numbers})
    }
	}

	fmt.Printf("%c\n", lights)
	fmt.Println(buttons)


	return count, nil
}


