package main

import (
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileName := "input1.dat"
	Part1(fileName)
	Part2(fileName)
}

func Part1(fileName string) {
	lines := readFile(fileName)

	data := make([][]string, 0, len(lines))

	for _, line := range lines {
		data = append(data, strings.Fields(line))
	}

	total := 0

	for i := range len(data[0]) {
		sign := data[len(data)-1][i]
		values := make([]int, 0, len(data)-1)
		for j := range len(data) - 1 {
			val, err := strconv.Atoi(data[j][i])
			checkErr(err)
			values = append(values, val)
		}

		res := calculateResults(sign, values...)
		total += res
	}

	fmt.Printf("Part1 -- Homework answer is: %d \n", total)
}

func Part2(fileName string) {
	lines := readFile(fileName)

	blockSizes := determineStringGroups(lines[len(lines) - 1])	

	total := 0

	start_idx := 0
	for i, size := range blockSizes {
		sign := strings.Fields(lines[len(lines)-1])[i]
		block := make([]string, 0, len(lines)-1)
		end_idx := start_idx + size
		for j := range len(lines) - 1 {
			block = append(block, lines[j][start_idx:end_idx])
		}
		
		values := determineCorrectValues(block...)

		res := calculateResults(sign, values...)
		total += res

		start_idx = end_idx + 1
	}

	fmt.Printf("Part2 -- Homework answer is: %d \n", total)
}

func determineStringGroups(line string) []int {
	sign_count := strings.Count(line,"*") + strings.Count(line,"+")

	block_sizes := make([]int, 0, sign_count)

	prev_start := 0
	for i, s := range line {
		if i == 0 {
			continue
		}
		if s == '+' || s == '*' {
			block_sizes = append(block_sizes, i - prev_start - 1)
			prev_start = i
		}
	}
	block_sizes = append(block_sizes, len(line) - prev_start)

	return block_sizes
}

func determineCorrectValues(inc_vals ...string) []int {
	maxValSize := len(slices.MaxFunc(inc_vals, func(a string, b string) int {
		return cmp.Compare(len(a), len(b))
	}))

	new_vals := make([]string, maxValSize)

	for i := maxValSize - 1; i >= 0; i-- {
		for _, val := range inc_vals {
			if len(val)-1 < i {
				continue
			}
			s := strings.Split(val,"")[i]
			if s != " " {
				new_vals[i] += s			
			}
		}
	}

	values := make([]int, 0, len(new_vals))
	for _, val := range new_vals {
		int_val, err := strconv.Atoi(val)
		checkErr(err)
		values = append(values, int_val)
	}

	return values
}

func calculateResults(sign string, values ...int) int {
	total := 0
	switch sign {
	case "+":
		for _, val := range values {
			total += val
		}
	case "*":
		total = 1
		for _, val := range values {
			total *= val
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

	return lines
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
