package main

import (
	"advent-of-code-website/solutions/utils"
	"advent-of-code-website/types"
	"fmt"
)

func main() {

    result1, _ := Part1()
    result2, _ := Part2()

    fmt.Printf("Part 1: %d\n", result1)
    fmt.Printf("Part 2: %d\n", result2)

		result := types.SolutionOutput{}
    result.Part1 = result1
    result.Part2 = result2
		utils.OutputJSON(result)
}