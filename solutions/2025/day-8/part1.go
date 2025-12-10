package main

import (
	"advent-of-code-website/solutions/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1() (int,error) {

	scanner, err := utils.ReadFile("")
	if err != nil {
			return -1 , err
	}

	var positions []Position
	var allConnections []Pair
	numberOfConnections := 1000

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
	output := ProcessPairs(allConnections[:numberOfConnections])

	return output, nil
}

func calculateDistance(p Position, np Position) float64 {

	dx := math.Pow(float64(np.x - p.x), 2.0)
	dy := math.Pow(float64(np.y - p.y), 2.0)
	dz := math.Pow(float64(np.z - p.z), 2.0)

	sum := dx + dy + dz

	return math.Sqrt(sum)

}

type Position struct {
	x int
	y int
	z int
}

type Pair struct {
	A, B Position
	distance float64
}

type UnionFind struct {
	parent map[Position]Position
}

func CreateUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[Position]Position),
	}
}

func (uf *UnionFind) Find(p Position) Position {

	if _, exists := uf.parent[p]; !exists {
		uf.parent[p] = p
		return p
	}

	for uf.parent[p] != p {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}
	return p

}

func (uf *UnionFind) Union(a Position, b Position) {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	uf.parent[rootB] = rootA

	if rootA != rootB {
		uf.parent[rootB] = rootA
	}
}

func (uf *UnionFind) GetAllSets() map[Position][]Position {
	sets := make(map[Position][]Position)

	for pos := range uf.parent {
		root := uf.Find(pos)
		sets[root] = append(sets[root], pos)
	}

	return sets
}

func (uf *UnionFind) Sort() [][]Position {
	sets := uf.GetAllSets()

	result := make([][]Position, 0, len(sets))
	for _, group := range sets {
		result = append(result, group)
	}

	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) > len(result[j])
	})

	return result
}

func ProcessPairs(pairs []Pair) int {
	uf := CreateUnionFind()

	for _, pair := range pairs {
		uf.Union(pair.A, pair.B)
	}

	output := uf.Sort()

	return len(output[0]) * len(output[1]) * len(output[2])
}



