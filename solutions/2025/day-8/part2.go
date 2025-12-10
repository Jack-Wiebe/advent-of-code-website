package main

import (
	"advent-of-code-website/solutions/utils"
	"math"
	"strconv"
	"strings"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	var positions []Position
	minDistance:=math.MaxFloat64
	lastMinDistance:=0.0

	for scanner.Scan(){

		line := scanner.Text()
		//fmt.Println(line)
		strArray := strings.Split(line, ",")
		x,_:=strconv.Atoi(strArray[0])
		y,_:=strconv.Atoi(strArray[1])
		z,_:=strconv.Atoi(strArray[2])
    positions = append(positions, Position{x:x, y:y, z:z})

	}

	//fmt.Println(positions)

	var nodeA Position
	var nodeB Position
	var lastPair Pair
	exit:=false
	uf := CreateUnionFind()
	for{

		for i,p := range(positions){
			for _,np := range(positions[i+1:]){
				distance := calculateDistance(p, np)
				if distance < minDistance && distance > lastMinDistance{
					minDistance = distance
					nodeA = p
					nodeB = np
					uf.Union(nodeA, nodeB)
					lastPair = Pair{A:nodeA, B:nodeB}
				}
			}
		}

		lastMinDistance = minDistance
		minDistance = math.MaxFloat64
		if exit {
			break
		}
		if len(uf.parent) == len(positions){
			exit = true
		}
	}

	return lastPair.A.x * lastPair.B.x, nil
}
