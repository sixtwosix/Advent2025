package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Coordinates struct {
	row int
	col int
}

func (c *Coordinates) IsEmpty() bool {
	if c == nil {
		return true
	}
	return false
}

func (c *Coordinates) IsEqual(b Coordinates) bool {
	if c.col == b.col && c.row == b.row {
		return true
	}
	return false
}

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

	start := Coordinates{row: start_row, col: start_col}
	total := CalculateAllPaths(grid, start)

	fmt.Printf("Total paths: %d \n", total)
}
/* procedure DFS(G, v) is
  label v as discovered
	for all directed edges from v to w that are in G.adjacentEdges(v) do
			if vertex w is not labeled as discovered then
				recursively call DFS(G, w) */
func CalculateAllPaths(grid [][]string, start Coordinates) int {

	min_col, max_col := 0, len(grid[0])
	max_row := len(grid)
	discovered := make(map[Coordinates]bool)
	memo := make(map[Coordinates]int)
	
	var dfsRec func(v Coordinates) int
	dfsRec = func (v Coordinates) int {

		if val, ok := memo[v]; ok {
			return val
		}

		discovered[v] = true
		// fmt.Println(v)

		if v.row+1 >= max_row {
			memo[v] += 1
			return 1
		}
		total := 0
		switch grid[v.row+1][v.col] {
		case "^":

			// Left
			l := Coordinates{row: v.row+1, col: v.col-1}
			// if _, ok := discovered[l];l.col >= min_col && !ok{
			if l.col >= min_col {
				total += dfsRec(l)
			}
			// Right
			r := Coordinates{row: v.row+1, col: v.col+1}
			// if _, ok := discovered[r];r.col < max_col && !ok{
			if r.col < max_col {
				total += dfsRec(r)
			}
		case ".":
			// Down
			d := Coordinates{row: v.row+1, col: v.col}
			total += dfsRec(d)
		}
		memo[v] += total
		return total
	}

	total := dfsRec(start)

	return total

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


