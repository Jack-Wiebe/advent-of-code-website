package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func Part1() (int,error){

	file,_ := os.Open("input")

	scanner := bufio.NewScanner(file)

	defer file.Close()

	dict := make(map[int][]int)
	clashes := make(map[int][]int)

	for scanner.Scan(){
		str := scanner.Text()
		line := strings.Split(str, ": ")
		key,_ := strconv.Atoi(line[0])
		val := strings.Split(line[1], " ")
		var int_vals []int

		for _,el := range val{
			el_int,_ := strconv.Atoi(el)
			int_vals = append(int_vals, el_int)
		}

		if len(dict[int(key)]) > 0{
			fmt.Println("FAIL")
			clashes[int(key)] = int_vals

		}else{
			dict[int(key)] = int_vals
		}

	}

	fmt.Println("Clashes in primary map:", clashes)

	sum := testSymbols(dict)
	sum += testSymbols(clashes)

	fmt.Println("The sum of valid inputs is", sum)

	return sum, nil

}

func getBitArrays(size int) [][]int {

	max_bitmask := (1 << size) - 1
	bitmask_length := bits.Len(uint(max_bitmask))

	var output [][]int

	for bitmask := 0; bitmask <= max_bitmask; bitmask++{
		//fmt.Printf("Bitmask: %04b (Decimal: %d)\n", bitmask, bitmask)
		var nested []int

		for i := bitmask_length - 1; i >= 0; i-- {
			bit := (bitmask >> i) & 1
			//fmt.Printf("Bit %d: %d\n", i, bit)
			nested = append(nested, bit)
		}
		output = append(output, nested)

	}

	return output

}

func testSymbols(dict map[int][]int) int {

	sum := 0

	for key, symbol_number := range dict{

		test_arrays := getBitArrays(len(symbol_number) - 1)

		for _, tests := range test_arrays{
			test := symbol_number[0]

			for ind, val := range symbol_number{

				if ind==0{
					continue
				}

				bit := tests[ind - 1]

				if bit == 0{
					test *= val
				}else{
					test += val
				}

			}
			if test == key {
				fmt.Println("VALID INPUT >>>>>", symbol_number, key, tests)
				sum += test
				break;
			}

		}
	}

	return sum
}