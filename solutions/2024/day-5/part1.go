package main

import (
	"advent-of-code-website/solutions/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() (int, error) {

	file,err := os.Open("input")

	if err != nil{
		fmt.Println("error opening file", err)
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

	fmt.Println(updates)
	fmt.Println()
	fmt.Println(page_map)
	fmt.Println()

	var output [][]string
	count := 0


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
			output = append(output, update)
			foo,_ := strconv.Atoi(update[(len(update)-1)/2])
			count += foo
		}

	}

	fmt.Println(output)
	fmt.Println()
	fmt.Println("part 1 count is:", count)
	fmt.Println()

	return count, nil
}

func check_pages(page_map map[string][]string, key string, substring []string) bool {

	check := true
	for _,val := range substring{
		if !utils.Contains(page_map[key], string(val)){
			check = false
			break
		}
	}
	return check

}

// func contains(slice []string, value string) bool {
// 	for _, v := range slice {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }