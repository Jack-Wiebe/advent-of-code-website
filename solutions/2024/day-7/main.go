package main

import (
	"advent-of-code-website/solutions/utils"
	"advent-of-code-website/types"
	"fmt"
)

func main() {

    result1, _ := Part1()

    fmt.Printf("Part 1: %d\n", result1)

		result := types.SolutionOutput{}
    result.Part1 = result1

		utils.OutputJSON(result)
}