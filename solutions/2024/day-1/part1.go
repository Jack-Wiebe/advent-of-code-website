package main

import (
	"advent-of-code-website/solutions/utils"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1() (int,error) {

	scanner, err := utils.ReadFile()
	if err != nil {
			return -1 , err
	}

	dif := 0
	var left_list []int
	var right_list []int

	for scanner.Scan(){
		//fmt.Println(scanner.Text())
		str := scanner.Text()
		result := strings.Fields(str)

		int0, _ := strconv.Atoi(result[0])
		int1, _ := strconv.Atoi(result[1])

		left_list = append(left_list, int0)
		right_list = append(right_list, int1)

	}

	sort.Ints(right_list)
	sort.Ints(left_list)

	for i := 0; i < len(left_list); i++ {
		dif += utils.AbsInt(left_list[i] - right_list[i])
	}

	fmt.Println(dif)

	return dif, nil

}

func Part2() (int,error) {

	file, err := os.Open("input")

	if err != nil {
		fmt.Println("Error opening file", err)
		return -1,err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left_list []int
	var right_list []int

	for scanner.Scan(){
		//fmt.Println(scanner.Text())
		str := scanner.Text()
		result := strings.Fields(str)

		int0, _ := strconv.Atoi(result[0])
		int1, _ := strconv.Atoi(result[1])

		left_list = append(left_list, int0)
		right_list = append(right_list, int1)

	}

	foo := 0

	for i := 0; i < len(left_list); i++ {
		mult := count(right_list, left_list[i])
		foo += mult * left_list[i]
	}

	fmt.Println(foo)
	return foo,nil

}

func count(slice []int, value int) int {
	count := 0
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}
