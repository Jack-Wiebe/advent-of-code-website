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
	joltage := ""
	length := 12


	for scanner.Scan(){
		str := scanner.Text()

		for i := range(length-1){

			strSize := length - i - 1
			substring := str[:len(str)-strSize]
			digit, index := getMax(substring)
			str = str[index+1:]

			joltage+=strconv.Itoa(digit)

		}

		lastDigit,_ := getMax(str)
		joltage+=strconv.Itoa(lastDigit)

		//fmt.Println(joltage)
		val,_ := strconv.Atoi(joltage)
		count+=val
		joltage=""
	}

	return count, nil
}

func getMax(substring string) (int,int) {

	digit:=0
	index:=0
	//find highest digit
	for i,n := range(substring){
			d,_ := strconv.Atoi(string(n))
			if d > digit{
				digit = d
				index = i
				if d == 9 { //exit early if we find a max value digit
					return d, i
				}
			}
		}


	return digit, index

}