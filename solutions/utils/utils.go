package utils

import (
	"advent-of-code-website/types"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Mod(x int, m int) int {
    result := x % m
    if result < 0 {
        result += m
    }
    return result
}

func Contains[T comparable](input []T, value T) bool {
	for _, v := range input {
		if v == value {
			return true
		}
	}
	return false
}

func MoveElement[T any](input []T, fromIndex int, toIndex int) []T {

	element := input[fromIndex]
	input = append(input[:fromIndex], input[fromIndex+1:]...)
	input = append(input[:toIndex], append([]T{element}, input[toIndex:]...)...)

	return input
}

func Splice[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	test := append(slice[:index], slice[index+1:]...)
	return test
}

func GridToString(grid [][]rune) string {
    if grid == nil {
        return "nil"
    }
    if len(grid) == 0 {
        return "[]"
    }

    var sb strings.Builder
    for i, row := range grid {
        if i > 0 {
            sb.WriteString("\n")
        }
        for _, ch := range row {
            if ch == 0 {
                sb.WriteString("·") // Show zeros as dots
            } else {
                sb.WriteRune(ch)
            }
        }
    }
    return sb.String()
}

func ReadFile(fileName ...string) (*bufio.Scanner, error) {

	fn := "input"

	if len(fileName) > 0 && fileName[0] != "" {
        fn = fileName[0]
    }
    file, err := os.Open(fn)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    return bufio.NewScanner(strings.NewReader(string(content))), nil
}

func OutputJSON(result types.SolutionOutput) {
    jsonData, err := json.Marshal(result)
    if err != nil {
        // Fallback to plain text if JSON fails
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Println(string(jsonData))
}