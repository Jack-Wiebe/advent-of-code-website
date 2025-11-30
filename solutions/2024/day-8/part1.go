package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part1() (int, error) {

	file,_ := os.Open("input")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var input [][]string

	for scanner.Scan(){
		str := strings.Split(scanner.Text(), "")
		input = append(input, str)
		fmt.Println(str)
	}
	fmt.Println()

	node_map := make(map[string][]Position)

	for row, line := range input{
		for col, cell := range line{
			if cell == "."{
				continue
			}
			node_map[cell] = append(node_map[cell], Position{x:row, y:col})
		}
	}

	for node_type := range node_map{
		anti_node_list := findAntiNode(node_map[node_type])
		for _, node := range anti_node_list{
			if node.x >= len(input) || node.x < 0 || node.y >= len(input[0]) || node.y < 0 {
				continue
			}
			input[node.x][node.y] = "#"
		}
	}

	count:=0
	for _,str := range input{
		for _,char := range str{
			if char == "#"{
				count++
			}
		}
		fmt.Println(str)
	}

	fmt.Println("total anti nodes:", count)
	return count, nil

}

func findAntiNode(nodes []Position) []Position{
	var anti_node_list []Position
	for ind,node := range nodes{
		for _,neighbour := range nodes[ind+1:]{
			anit_vec := Position{x:node.x - neighbour.x, y:node.y-neighbour.y}
			anit_node_alpha := Position{x:node.x + anit_vec.x,y:node.y+anit_vec.y}
			anit_node_beta := Position{x:neighbour.x - anit_vec.x,y:neighbour.y-anit_vec.y}
			anti_node_list = append(anti_node_list, anit_node_alpha, anit_node_beta)
		}
	}
	return anti_node_list
}

type Position struct {
	x int
	y int
}