package main

import (
	"fmt"
	"math"
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
	banks := readFile(fileName)
	total := 0

	for _, row := range banks {
		res := getMostJoltage(row, 2)
		// fmt.Printf(" -- %d -- \n", res)
		total += res
	}

	fmt.Printf("Part1 -- Total = %d \n", total)
}

func Part2(fileName string) {
		banks := readFile(fileName)
	total := 0

	for _, row := range banks {
		res := getMostJoltage(row, 12)
		// fmt.Printf(" -- %d -- \n", res)
		total += res
	}

	fmt.Printf("Part2 -- Total = %d \n", total)
}

func getMostJoltage(s string, maxDepth int) int {
	firstNum := 0
	total := 0

	for i, x := range s {
		if i == len(s) - (maxDepth - 1) {
			break
		}
		// substract '0' to get digital value
		num := int(x - '0')
		if num > firstNum {
			firstNum = num
			max := 0

			if maxDepth == 2 {
				for _, y := range s[i+1:] {
				if int(y - '0') > max {
					max = int(y - '0')
				}
				}
				if max == 0 {
					continue
				}
			} else {
				max = getMostJoltage(s[i+1:], maxDepth - 1)
			}
			if (firstNum*(int(math.Pow10(maxDepth-1))) + max) > total {
				total = firstNum*(int(math.Pow10(maxDepth-1))) + max
			}	
		}
	}

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