package main

import (
	"advent-of-code-website/solutions/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part2() (int, error) {

	file,err := os.Open("input")

	if err != nil{
		fmt.Println("error opening file", err)
		return -1, err
	}

	scanner := bufio.NewScanner(file)

	defer file.Close()

	var input string
	for scanner.Scan(){
		str := scanner.Text()
		input += str + "\n"
	}

	sections := strings.Split(input, "\n\n")
	pages := strings.Split(sections[0], "\n")
	updates := strings.Split(sections[1], "\n")

	page_map := make(map[string][]string)

	for _,val := range pages{
		temp := strings.Split(val, "|")
		page_map[temp[0]] = append(page_map[temp[0]], temp[1])
	}
	count := 0
	count2 := 0


	for _,val := range updates {
		pass := true
		update := strings.Split(val, ",")

		for i := 0; i < len(update); i++{
			if(check_pages(page_map, update[i], update[i+1:])){
				//fmt.Println("PASS")
			}else{
				//fmt.Println("FAIL")
				pass = false
				break
			}
		}

		if(pass){
			//output = append(output, update)
			foo,_ := strconv.Atoi(update[len(update)/2])
			count += foo
		}else{
			count2 += fix_fail(update, page_map)
		}

	}

	fmt.Println("part 2 count is:", count2)

	return count2, nil
}

func fix_fail(update []string, page_map map[string][]string) int {


	//this method is not working as it assumes moving the out of place value forward solves the soltion
	// instead it needs to be checked again each value ahead of it to find its correct place
	//or seperate out each val in wrong spot and resort them in on a second pass
	fmt.Println(update)

	var temp []int

	for ind,key := range update{
		for i,val := range update[ind+1:]{
			if !utils.Contains(page_map[key],val){
				temp = append(temp, ind+i+1)
				move_element(update, ind+i+1, ind+i)
			}
		}
	}

	// fmt.Println(temp)

	for i := range temp{
		for j := len(update); j > 0; j--{
			if(!check_pages(page_map, update[temp[i]], update[j:])){
				move_element(update, temp[i], j)
			}
		}
	}

	middle,_ := strconv.Atoi(update[(len(update)-1)/2])
	fmt.Println("Fixing Broken Page")
	fmt.Println(update)
	fmt.Println(middle)
	fmt.Println()
	return middle
}


func move_element[T any](input []T, fromIndex int, toIndex int) []T {

	element := input[fromIndex]
	input = append(input[:fromIndex], input[fromIndex+1:]...)
	input = append(input[:toIndex], append([]T{element}, input[toIndex:]...)...)

	return input
}