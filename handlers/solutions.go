package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"advent-of-code-website/types"

	"github.com/gorilla/mux"
)

func ListSolutionsHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    year := vars["year"]

    solutions := []types.Solution{}

    files, err := os.ReadDir("./solutions/" + year)
    if err != nil {
        http.Error(w, "Cannot read solutions directory", http.StatusInternalServerError)
        return
    }

    for _, file := range files {
        if file.IsDir() && strings.HasPrefix(file.Name(), "day") {
            solutions = append(solutions, types.Solution{
                Day:    file.Name(),
                Name:   fmt.Sprintf("Day %s", strings.TrimPrefix(file.Name(), "day")),
                Status: "Available",
            })
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(solutions)
}

func ListYearsHandler(w http.ResponseWriter, r *http.Request) {
    years := []types.Year{}

    files, err := os.ReadDir("./solutions")
    if err != nil {
        http.Error(w, "Cannot read solutions directory", http.StatusInternalServerError)
        return
    }

    pattern:=`^\d{4}$`
    re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Printf("Error compiling regex: %v\n", err)
		return
	}

    for _, file := range files {
        if file.IsDir(){
            if re.MatchString(file.Name()){
            years = append(years, types.Year{
                Name:    file.Name(),
            })
        }
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(years)
}

func RunSolutionHandler(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    day := vars["day"]
    year := vars["year"]

    result := types.SolutionOutput{Day: day}

    if !strings.HasPrefix(day, "day") {
        result.Error = "Invalid day format"
        sendJSONResponse(w, result)
        return
    }

    solutionPath := filepath.Join("./solutions/", year, day)
    if _, err := os.Stat(solutionPath); os.IsNotExist(err) {
        result.Error = fmt.Sprintf("Solution for %s not found", day)
        sendJSONResponse(w, result)
        return
    }

    originalDir, _ := os.Getwd()
    defer os.Chdir(originalDir)

    if err := os.Chdir(solutionPath); err != nil {
        result.Error = fmt.Sprintf("Cannot change to solution directory: %v", err)
        sendJSONResponse(w, result)
        return
    }

    cmd := exec.Command("go", "run", ".")
    output, err := cmd.CombinedOutput()
    outputStr := string(output)

    if err != nil {
        result.Error = fmt.Sprintf("Execution error: %v", err)
        sendJSONResponse(w, result)
        return
    }

    jsonStr, consoleStr := extractJSONFromOutput(outputStr)
    if jsonStr == "" {
        result.Error = "No JSON found in output"
        sendJSONResponse(w, result)
        return
    }
    result.Output = consoleStr

    var solutionOutput SolutionOutput
    if err := json.Unmarshal([]byte(jsonStr), &solutionOutput); err != nil {
        result.Error = "Failed to parse JSON: " + err.Error()
    } else {
        result.Part1 = solutionOutput.Part1
        result.Part2 = solutionOutput.Part2
        if solutionOutput.Error != "" {
            result.Error = solutionOutput.Error
        }
    }

    sendJSONResponse(w, result)
}

func extractJSONFromOutput(output string) (jsonStr string, cleanOutput string) {
    lines := strings.Split(output, "\n")
    var outputLines []string
    var foundJSON string

    for i := len(lines) - 1; i >= 0; i-- {
        line := strings.TrimSpace(lines[i])
        if line == "" {
            outputLines = append([]string{lines[i]}, outputLines...)
            continue
        }

        if strings.HasPrefix(line, "{") && strings.HasSuffix(line, "}") && foundJSON == "" {
            foundJSON = line
            continue
        }

        outputLines = append([]string{lines[i]}, outputLines...)
    }

    return foundJSON, strings.Join(outputLines, "\n")
}

type SolutionOutput struct {
    Part1 int    `json:"part1"`
    Part2 int    `json:"part2"`
    Error string `json:"error,omitempty"`
    Output  string `json:"output,omitempty"`
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}