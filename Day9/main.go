package main

import (
	"advent2025/Day9/heap"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	fileName := "input1.dat"
	Part1(fileName)
	// Part2(fileName)

}

type Coordinate struct {
	row int
	col int
}

type CoordSet struct {
	tileA Coordinate
	tileB Coordinate
}

func Part1(fileName string) {
	lines := readFile(fileName)

	tiles := make([]Coordinate, 0, len(lines))

	for _, l := range lines {
		data := strings.Split(l, ",")
		col, err := strconv.Atoi(data[0])
		checkErr(err)
		row, err := strconv.Atoi(data[1])
		checkErr(err)
		tiles = append(tiles, Coordinate{row: row, col: col})
	}

	heap := heap.MaxHeap[int]{}
	coordSets := make(map[CoordSet]int)

	for _, t1 := range tiles {
		for _, t2 := range tiles {
			// Calculate manhattan distance
			total := int(math.Abs(float64(t1.col - t2.col)) + 1 + math.Abs(float64(t1.row - t2.row)) + 1)
			heap.Insert(total)
			cSetA := CoordSet{tileA: t1, tileB: t2}
			cSetB := CoordSet{tileA: t2, tileB: t1}
			_, okA := coordSets[cSetA]
			_, okB := coordSets[cSetB]
			if !okA && !okB {
				coordSets[cSetA] = total
			}			
		}
	}

	maxManhattan := heap.Remove()

	filtered := mapFilter(coordSets, func(cs CoordSet, i int) bool {
		return i == maxManhattan
	})

	for key := range filtered {
		area := determineArea(key.tileA, key.tileB)
		fmt.Printf("Biggest area: %d \n", area)
	}
}

func Part2(fileName string) {
}

func determineArea(coordA, coordB Coordinate) int {
	rowDiff := int(math.Abs(float64(coordA.row) - float64(coordB.row))) + 1
	rowCol := int(math.Abs(float64(coordA.col) - float64(coordB.col))) + 1

	return rowCol * rowDiff
}

func mapFilter(m map[CoordSet]int, pred func(CoordSet, int) bool) map[CoordSet]int {
	result := make(map[CoordSet]int)

	for k, v := range m {
		if pred(k, v) {
			result[k] = v
		}
	}

	return result
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