package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fileName := "input1.dat"
	Part1(fileName)
	Part2(fileName)
}

func Part1(fileName string) {
	lines := readFile(fileName)
	grid := make([][]string, 0, len(lines))
	start_row, start_col := 0, 0
	for row, line := range lines {
		grid = append(grid, strings.Split(line, ""))
		if idx := strings.Index(line, "S"); idx != -1 {
			start_row = row
			start_col = idx
		}
	}

	min_col, max_col := 0, len(grid[0])
	beams := make(map[int]int, max_col)
	beams[start_col] = 1

	total_split := 0

	for row := start_row + 1; row < len(grid); row++ {
		for idx, count := range beams {
			if count > 0 {
				if grid[row][idx] == "^" {
					// Left
					if idx >= min_col {
						beams[idx-1] = 1
					}

					// Right
					if idx < max_col {
						beams[idx+1] = 1
					}

					beams[idx] = 0
					total_split += 1
				}
			}
		}
	}

	fmt.Printf("Part1 -- Total beams %d \n", total_split)
}

func Part2(fileName string) {
}

func readFile(fileName string) []string {
	cwd, err := os.Getwd()
	checkErr(err)
	path := filepath.Join(cwd, fileName)
	dat, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(string(dat), "\n")

	result := make([]string, len(lines))
	for i, line := range lines {
		result[i] = strings.TrimSpace(line)
	}

	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
