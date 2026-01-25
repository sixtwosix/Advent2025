package main

import (
	"advent2025/Day8/heap"
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

func Part1(fileName string) {

	maxPairs := 1000

	lines := readFile(fileName)
	junctionBoxes := make([]coordinate, 0, len(lines))
	for _, line := range lines {
		temp := strings.Split(line,",")
		x, err := strconv.Atoi(temp[0]) 
		checkErr(err)
		y, err := strconv.Atoi(temp[1]) 
		checkErr(err)
		z, err := strconv.Atoi(temp[2]) 
		checkErr(err)
		junctionBoxes = append(junctionBoxes, coordinate{
			x: x,
			y: y,
			z: z,
		})
	}

	array := []float64{}
	minHeap := &heap.MinHeap[float64]{}
	minHeap.BuildHeap(array)

	juncCombos := make(map[CoordinateSet]float64)


	for _, jbA := range junctionBoxes {
		InnerLoop:
		for _, jbB := range junctionBoxes {
			if jbA == jbB {
				continue InnerLoop
			}
			set := CoordinateSet{boxA: jbA, boxB: jbB}
			setRev := CoordinateSet{boxA: jbB, boxB: jbA}
			_, ok := juncCombos[set]
			_, okRev := juncCombos[setRev]
			if ok || okRev {
				continue InnerLoop
			}
			res := determineEuclideanDistance(jbA, jbB)
			minHeap.Insert(res)
			juncCombos[set] = res
		}
	}

	// Pop first 1000 from minHeap

	// Join all sets
	// do merge when two groups are joining together

}

func Part2(fileName string) {
	
}

func determineEuclideanDistance(a, b coordinate) float64 {
	x := math.Pow((float64(a.x)-float64(b.x)),2)
	y := math.Pow((float64(a.y)-float64(b.y)), 2)
	z := math.Pow((float64(a.z)-float64(b.z)),2)
	result := math.Sqrt(x+y+z)
	return result
}

type CoordinateSet struct {
	boxA coordinate
	boxB coordinate
}

type coordinate struct {
	x int
	y int
	z int
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