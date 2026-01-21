package main

import (
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

	
}

func Part2(fileName string) {
	
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