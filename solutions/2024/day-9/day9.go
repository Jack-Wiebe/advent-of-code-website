package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() (int, error) {

	file,_ := os.Open("input")

	scanner := bufio.NewScanner(file)

	defer file.Close()

	var input []string
	var output []string
	var output_prime []string

	for scanner.Scan(){
		input = strings.Split(scanner.Text(), "")
	}

	for ind,val := range input{
		num, _ := strconv.Atoi(val)
		if ind % 2 == 0{
			for range num  {
				output = append(output, strconv.Itoa((ind/2)))
			}
		}else{
			for range num{
				output = append(output, ".")
			}
		}
	}

	fmt.Println(output)

	offset:=1
	for i,val := range output{
		if offset > (len(output) - i){
			output_prime = append(output_prime, ".")
			continue
		}
		if val != "."{
			output_prime = append(output_prime, output[i])
		}else{
			for output[len(output)-offset] == "."{
				offset++
				continue
			}
			temp := output[len(output)-offset]
			output_prime = append(output_prime, temp)
			offset++
		}
	}

	fmt.Println(output_prime)
	count:=0
	for ind,val := range output_prime{
		if(val == "."){
			break
		}
		num,_ := strconv.Atoi(val)
		count = count + (num * ind)
	}

	fmt.Println(count)

	return count, nil

}