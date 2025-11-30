package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Part2() (int, error){

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
		anti_node_list := findAntiNodeRange(node_map[node_type], len(input), len(input[0]))
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

func findAntiNodeRange(nodes []Position, grid_width int, grid_height int) []Position{
	fmt.Println(grid_width, grid_height)
	var anti_node_list []Position
	for ind,node := range nodes{
		for _,neighbour := range nodes[ind+1:]{
			anit_vec := Position{x:node.x - neighbour.x, y:node.y-neighbour.y}
			i:=1
			// /j:=1
			fmt.Println(anit_vec)

			for {
				i++
				if node.x + (anit_vec.x * i) >= grid_width || node.y + (anit_vec.y * i) >= grid_height || node.x - (anit_vec.x * i) < 0 || node.y - (anit_vec.y * i) < 0{
					break
				}
				anit_node_alpha := Position{x:node.x + anit_vec.x,y:node.y+anit_vec.y}
				anti_node_list = append(anti_node_list, anit_node_alpha)
			}
			// for node.x - (anit_vec.x * j) < 0 || node.y - (anit_vec.y * j) < 0{
			// 	j++
			// 	anit_node_beta := Position{x:neighbour.x - anit_vec.x,y:neighbour.y-anit_vec.y}
			// 	anti_node_list = append(anti_node_list, anit_node_beta)
			// }



		}
	}
	return anti_node_list
}

