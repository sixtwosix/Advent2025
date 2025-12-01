package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	fileName := "input1.dat"
	// Part1(fileName)
	Part2(fileName)

}

func Part1(fileName string) {
	data := readFile(fileName)

	realPassword := 0
	dial := 50

	for _, line := range data {
		updown := string(line[0])
		count, err := strconv.Atoi(line[1:])
		checkErr(err)

		if updown == "L" {
			dial -= count
		} else {
			dial += count
		}
		for {
			if dial < 0 {
				dial += 100
			} else {
				break
			}
		}
		dial = dial % 100

		if dial == 0 {
			realPassword += 1
		}
	}

	fmt.Printf("Part1: The real password: %d\n", realPassword)

}

func Part2(fileName string) {
	data := readFile(fileName)

	realPassword := 0
	dial := 50

	for _, line := range data {
		updown := string(line[0])
		count, err := strconv.Atoi(line[1:])
		checkErr(err)

		rotation := int(math.Abs(float64(count/100)))
		realPassword += rotation
		count = count % 100

		if updown == "L" {
			if dial == 0 {
				dial += 100
			}
			dial -= count
		} else {
			dial += count
		}

		if dial < 0 {
			realPassword += 1
			dial += 100
		}

		if dial >= 100 {
			realPassword += 1
		}

		if dial == 0 {
			realPassword += 1
		}

		dial = dial % 100

		fmt.Printf("Dial: %d -- RealPassword: %d \n", dial, realPassword)
	}

	fmt.Printf("Part2: The real real password: %d\n", realPassword)
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