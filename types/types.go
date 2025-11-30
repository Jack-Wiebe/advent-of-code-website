package types

type Solution struct {
	Day    string `json:"day"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
type Year struct {
	Name string `json:"name"`
}

type SolutionOutput struct {
	Day    string `json:"day"`
	Output string `json:"output"`
	Part1  int    `json:"part1,omitempty"`
	Part2  int    `json:"part2,omitempty"`
	Error  string `json:"error,omitempty"`
}
