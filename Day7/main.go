package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"advent2025/Day7/queue"
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
	fileName := "test_input1.dat"
	// Part1(fileName)
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

	fmt.Printf("Grid: \n %+v \n", grid)

	res := BFS(grid, Coordinates{row: start_row, col: start_col})

	starts := make([]Coordinates, 0, len(grid[0]))
	for key := range res {
		if key.row == len(grid) - 1 {
			starts = append(starts, key)
		}
	}

	fmt.Printf("Starting points: \n %+v \n", starts)
	
	total := 0

	for key, val := range res {
/* 		fmt.Printf("%+v \t %+v \n", key, val) */
		if slices.Contains(starts, key) {
			total += len(val)
		}
	}

	fmt.Printf("Total paths: %d \n", total)
}

func TotalPaths()  {
	
}

func BFS(grid [][]string, root Coordinates) map[Coordinates][]Coordinates {
	queue := queue.NewQueue[Coordinates]()
	explored := make(map[Coordinates][]Coordinates)

	explored[root] = []Coordinates{root}
	queue.Push(root)

	min_col, max_col := 0, len(grid[0])
	max_row := len(grid)

	for len(*queue) > 0 {
		v, ok := queue.Pop()
		if !ok {
			break
		}

		if v.row+1 >= max_row {
			continue
		}

		if grid[v.row+1][v.col] == "^" {

			if v.col-1 >= min_col {
				w1 := Coordinates{row: v.row + 1, col: v.col - 1}
				if _, ok := explored[w1]; !ok {
					explored[w1] = []Coordinates{v}
				} else {
					explored[w1] = append(explored[w1], v)
				}
				queue.Push(w1)
			}
			if v.col+1 < max_col {
				w2 := Coordinates{row: v.row + 1, col: v.col + 1}
				if _, ok := explored[w2]; !ok {
					explored[w2] = []Coordinates{v}					
				} else {
					explored[w2] = append(explored[w2], v)
				}
				queue.Push(w2)
			}

		}

		if grid[v.row+1][v.col] == "." {
			w := Coordinates{row: v.row + 1, col: v.col}
			if _, ok := explored[w]; !ok {
				explored[w] = []Coordinates{v}
				queue.Push(w)
			} else {
				explored[w] = append(explored[w], v)
			}
		}

	}

	return explored
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
