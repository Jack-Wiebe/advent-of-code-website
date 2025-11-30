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

	file, err := os.Open("input")

	if err != nil{
		return -1, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var report [][]string

	for scanner.Scan(){
		str := scanner.Text()
		line := strings.Split(str, " ")
		fmt.Println(line)
		report = append(report, line)
	}
	fmt.Println()

	count := 0
	for i := 0; i < len(report); i++{

		index := reportSafe(report[i])

		if(index == -1){
			fmt.Println("safe")
			count++
		}else{
			fmt.Println(report[i])
			test := utils.Splice(report[i],index)
			fmt.Println(test)

			index = reportSafe(test)
			if(index == -1){
				fmt.Println("safe")
				count++
			}
			fmt.Println(count)
			fmt.Println()
		}
	}

	fmt.Println(count)
	return count, nil

}

func reportSafe(report []string) int {
	dir := 0

	val0 := report[0]
	val1 := report[1]
	int0, err0 := strconv.Atoi(val0)
	int1, err1 := strconv.Atoi(val1)

	if err0 != nil || err1 != nil{
		fmt.Println("FAIL")
	}

	if  int0 == int1{
		fmt.Println("start fail")
		return 0
	}else if int0 > int1{
		dir = -1
	}else{
		dir = 1
	}


	for j := 1; j < len(report); j++{

		val0 := report[j]
		val1 := report[j - 1]
		int0, err0 := strconv.Atoi(val0)
		int1, err1 := strconv.Atoi(val1)

		if err0 != nil || err1 != nil{
			fmt.Println("FAIL")
			return -1
		}

		dif := int0 - int1

		dif = dif * dir

		if(dif == 0){
			fmt.Println("no movement")
			if(j == 2){
				return 0
			}else{
				return j
			}
		}else if(dif < 0){
			fmt.Println("direction change")

			if(j+1 < len(report) && report[j + 1] == report[j - 1]){
				fmt.Println("sandwich")
				if(dir > 0){
					fmt.Println("rising")
					if(report[j] < report[j + 1]){
						return j + 1
					}else{
						return j - 1
					}
				}else{
					fmt.Println("falling")
					if(report[j] < report[j - 1]){
						return j + 1
					}else{
						return j - 1
					}
				}



			}else if(j == 2){
				return j-2
			}else{
				return j
			}
		}else if(dif > 3){
			fmt.Println("overload")
			if(j == 1){
				return 0
			}else{
				return j
			}
		}
	}
	return -1
}

func Part2() (int, error) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numSafe := 0

	for scanner.Scan() {
		line := scanner.Text()
		records := strings.Split(line, " ")

		nums := make([]int, len(records))
		for i, record := range records {
			val, _ := strconv.Atoi(record)
			nums[i] = val
		}

		if testSafe(nums) {
			numSafe++
			continue
		}

		for i := range nums {
			clone := make([]int, len(nums))
			copy(clone, nums)
			resliced := append(clone[:i], clone[i+1:]...)
			if testSafe(resliced) {
				numSafe++
				break
			}
		}
	}

	fmt.Println("Part 2:",numSafe)
	return numSafe, nil
}

func testSafe(nums []int) bool {
	dir := (nums[1] - nums[0]) > 0
	for i := 1; i < len(nums); i++ {
		dif := nums[i] - nums[i-1]
		//fmt.Println(nums)
		if(dif > 0 != dir){
			//fmt.Println("direction change")
			return false
		}
		if(dif == 0){
			//fmt.Println("no movement")
			return false
		}
		if(dif > 3 || -dif > 3){
			//fmt.Println("overload")
			return false
		}
	}

	return true
}
