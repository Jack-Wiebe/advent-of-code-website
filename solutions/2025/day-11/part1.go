package main

import (
	"advent-of-code-website/solutions/utils"
	"strings"
)


func Part1() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	serverMap := make(map[string] []string)
	count:=0

	for scanner.Scan(){

		line := scanner.Text()

		strArr := strings.Split(line, ": ")
		key:= strArr[0]
		values:=strings.Split(strArr[1], " ")

		serverMap[key] = values

	}
	recursiveFind(&serverMap, "you", &count)

	return count, nil
}

func recursiveFind(serverMap *map[string] []string, key string, count *int) {
	for _,value := range (*serverMap)[key]{
		if value == "out"{
			(*count)++
		}else{
			recursiveFind(serverMap, value, count)
		}
	}
}


