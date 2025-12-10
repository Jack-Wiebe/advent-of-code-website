package main

import (
	"advent-of-code-website/solutions/utils"
	"sort"
	"strconv"
	"strings"
)

func Part2() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	var positions []Position
	var allConnections []Pair

	for scanner.Scan(){

		line := scanner.Text()
		strArray := strings.Split(line, ",")
		x,_:=strconv.Atoi(strArray[0])
		y,_:=strconv.Atoi(strArray[1])
		z,_:=strconv.Atoi(strArray[2])
    positions = append(positions, Position{x:x, y:y, z:z})

	}

	for i,p := range(positions){
		for _,np := range(positions[i+1:]){
			distance := calculateDistance(p, np)
			allConnections = append(allConnections, Pair{A: p, B: np, distance: distance})
		}
	}

	sort.Slice(allConnections, func(i, j int) bool {
		return allConnections[i].distance < allConnections[j].distance
	})

	var lastPair Pair
	connectionCount:=0
	uf := CreateUnionFind()

	for _,pair := range(allConnections){
		if uf.Union(pair.A, pair.B) {
			connectionCount++
			lastPair = Pair{A:pair.A, B:pair.B, distance: pair.distance}
		}
		if connectionCount == len(positions){
			break
		}
	}

	output:=lastPair.A.x * lastPair.B.x
	return output, nil
}
