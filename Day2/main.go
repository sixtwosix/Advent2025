package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	fileName := "input1.dat"
	Part1(fileName)
	Part2(fileName)

}

func Part1(fileName string) {
	data := readFile(fileName)

	total := 0

	for _, idRange := range data {
		startNum, err := strconv.Atoi(strings.Split(idRange,"-")[0])
		checkErr(err)
		endNum, err := strconv.Atoi(strings.Split(idRange,"-")[1])
		checkErr(err)

		results := GetInvalidIdsP1(startNum, endNum)
		for _, res := range results {
			resInt, err := strconv.Atoi(res)
			checkErr(err)

			total += resInt 
		}
	}

	fmt.Printf("Part1: Total - %d\n", total)
}

func Part2(fileName string) {
	data := readFile(fileName)

	total := 0

	for _, idRange := range data {
		startNum, err := strconv.Atoi(strings.Split(idRange,"-")[0])
		checkErr(err)
		endNum, err := strconv.Atoi(strings.Split(idRange,"-")[1])
		checkErr(err)

		results := GetInvalidIdsP2(startNum, endNum)
		for _, res := range results {
			resInt, err := strconv.Atoi(res)
			checkErr(err)

			total += resInt 
		}
	}

	fmt.Printf("Part2: Total - %d\n", total)
}

func GetInvalidIdsP2(startNum, endNum int) []string {
	invalidIds := make([]string,0)
	for id := startNum; id <= endNum; id++ {
		if !determineIfIdValidP2(id) {
			invalidIds = append(invalidIds, strconv.Itoa(id))
		}
	}

	return invalidIds
}

func determineIfIdValidP2(num int) bool {
		numStr := strconv.Itoa(num)
		ctr := len(numStr)/2

		for i := 1; i <= ctr; i++ {
			res := strings.Replace(numStr, numStr[:i], "", -1)
			if res == "" {
				return false
			}
		}

		return true
}

func GetInvalidIdsP1(startNum, endNum int) []string {
	invalidIds := make([]string,0)
	for id := startNum; id <= endNum; id++ {
		if !determineIfIdValidP1(id) {
			invalidIds = append(invalidIds, strconv.Itoa(id))
		}
	}

	return invalidIds
}

func determineIfIdValidP1(num int) bool {
	numStr := strconv.Itoa(num)
	ctr := len(numStr)/2
	
	if numStr[:ctr] == numStr[ctr:] {
		return false
	}	else {
		return true
	}
}

func readFile(fileName string) []string {

	cwd, err := os.Getwd()
	checkErr(err)
	path := filepath.Join(cwd, fileName)
	dat, err := os.ReadFile(path)
	checkErr(err)
	lines := strings.Split(string(dat), ",")
	
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