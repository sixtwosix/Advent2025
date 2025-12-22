package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type location struct {
	row int
	col int
}

func main() {

	fileName := "input1.dat"
	// Part1(fileName)
	Part2(fileName)

}

func Part1(fileName string) {
	data := readFile(fileName)
	
	count := 0

	for i, row := range data {
		for j, col := range row {
			if col == "@" {
				if checkMoveable(i, j, data) {
					count += 1
					// fmt.Printf("%d -- The paper at (%d - %d) is moveable \n", count, i, j)
				}
			}
		}
	}

	fmt.Printf("Part1 -- We can move a total of %d papers", count)
}

func Part2(fileName string) {
	data := readFile(fileName)

	totalRemoved := 0
	

	for {
		roundRemoved := 0
		collectedPaper := make(map[location]bool)
		for i, row := range data {
			for j, col := range row {
				if col == "@" {
					if checkMoveable(i,j,data) {
						collectedPaper[location{row: i, col: j}] = false
					}
				}
			}
		}
		for loc, removed := range collectedPaper {
			if !removed {
				data[loc.row][loc.col] = "."
				roundRemoved += 1
			}
		}
		totalRemoved += roundRemoved

		// fmt.Printf("Moved %d papers this round\n", roundRemoved)

		if roundRemoved <= 0 {
			break
		}
	}

	fmt.Printf("Part2 -- We moved %d papers\n", totalRemoved)
}

func checkMoveable(row, col int, data [][]string) bool {
	dx := []int {-1, 0, 1, 1, 1, 0, -1, -1}
	dy := []int {-1, -1, -1, 0, 1, 1, 1, 0}

	maxSurrounds := 4
	total := 0

	rowMax, colMax := determineMax(data)

	for i, rowDy := range dy {
		colDx := dx[i]

		newRow := rowDy + row
		newCol := colDx + col

		if (newRow < rowMax) && (newRow >= 0) && (newCol < colMax) && (newCol >= 0) {
			if data[newRow][newCol] == "@" {
				total += 1
			}
		}
	}

	if total >= maxSurrounds {
		return false
	} else {
		return true
	}
}

func determineMax(data [][]string) (rowMax, colMax int) {
	rowMax = len(data)
	colMax = len(data[0])
	return
}

func readFile(fileName string) [][]string {

	cwd, err := os.Getwd()
	checkErr(err)
	path := filepath.Join(cwd, fileName)
	dat, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(string(dat), "\n")
	
	result := make([][]string, len(lines))
	for i, line := range lines {
		result[i] = strings.Split(strings.TrimSpace(line),"")
	}

	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}