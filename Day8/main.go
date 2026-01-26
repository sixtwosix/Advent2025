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
	junctionBoxes := make([]Coordinate, 0, len(lines))
	for _, line := range lines {
		temp := strings.Split(line,",")
		x, err := strconv.Atoi(temp[0]) 
		checkErr(err)
		y, err := strconv.Atoi(temp[1]) 
		checkErr(err)
		z, err := strconv.Atoi(temp[2]) 
		checkErr(err)
		junctionBoxes = append(junctionBoxes, Coordinate{
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
	for i := 0; i < maxPairs; i++ {
		minDistance := minHeap.Remove()

		filteredMap := mapFilter(juncCombos, func(c CoordinateSet, val float64) bool {
			return val == minDistance
		})
		for k, v := range filteredMap {

		}

	}

	// Join all sets
	// do merge when two groups are joining together

}

func Part2(fileName string) {
	
}


func mapFilter(m map[CoordinateSet]float64, pred func(CoordinateSet, float64) bool) map[CoordinateSet]float64 {
	result := make(map[CoordinateSet]float64)

	for k, v := range m {
		if pred(k,v) {
			result[k] = v
		}
	}

	return result
}

func determineEuclideanDistance(a, b Coordinate) float64 {
	x := math.Pow((float64(a.x)-float64(b.x)),2)
	y := math.Pow((float64(a.y)-float64(b.y)), 2)
	z := math.Pow((float64(a.z)-float64(b.z)),2)
	result := math.Sqrt(x+y+z)
	return result
}

type CoordinateSet struct {
	boxA Coordinate
	boxB Coordinate
}

type Coordinate struct {
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