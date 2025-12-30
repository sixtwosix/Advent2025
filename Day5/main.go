package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

type idRange struct {
	start int
	end   int
}

type newIdType struct {
	startend int
	id int
}

func main() {
	fileName := "input1.dat"
	Part1(fileName)
	Part2(fileName)
}

func Part1(fileName string) {
	lines := readFile(fileName)

	idx := slices.Index(lines, "")

	ranges := lines[:idx]
	ids := lines[idx+1:]

	dict := determineRanges(ranges)

	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		checkErr(err)

		if r := determineFreshIngredients(id, dict); r.start != -1 {
			dict[r] = append(dict[r], id)
		}
	}

	total := determineTotalFresh(dict)

	fmt.Printf("Part1 -- Total fresh ingredients %d \n", total)

}

func Part2(fileName string) {
	lines := readFile(fileName)

	idx := slices.Index(lines, "")

	ranges := lines[:idx]
	
	freshIds := make([]idRange, 0)

	for _, rangeStr := range ranges {
		start, err := strconv.Atoi(strings.Split(rangeStr, "-")[0])
		checkErr(err)
		end, err := strconv.Atoi(strings.Split(rangeStr, "-")[1])
		checkErr(err)
		freshIds = append(freshIds, idRange{
			start: start,
			end: end,
		})
	}

	newRanges := newRangeSortingAlgorithm(freshIds)

	total := countFresh(newRanges)
	
	fmt.Printf("Part2 -- Total fresh ingredients %d \n", total)
}

func newRangeSortingAlgorithm(freshIds []idRange) []idRange {
	newRanges := make([]newIdType,0)

	for _, idr := range freshIds {
		newRanges = append(newRanges, newIdType{ startend: 0, id: idr.start})
		newRanges = append(newRanges, newIdType{ startend: 1, id: idr.end})
	}

	slices.SortFunc(newRanges, func(a,b newIdType) int {
		if a.id < b.id {
			return -1
		} else if a.id > b.id {
			return 1
		} else  {
			if a.startend < b.startend {
				return -1
			} else if a.startend > b.startend {
				return 1
			} else {
				return 0
			}
		}			
	})

	finalRanges := make([]idRange, 0)

	currVal := 0
	prevVal := 0

	startId := -1
	endId := -1
	for _, val := range newRanges {
		if val.startend == 0 {
			currVal += 1
		} else {
			currVal -= 1
		}

		if currVal == 1 && prevVal == 0 {
			startId = val.id
		}
		if currVal == 0 && prevVal == 1 {
			endId = val.id
		}
		if startId != -1 && endId != -1 {
			finalRanges = append(finalRanges, idRange{start: startId, end: endId})
			startId = -1
			endId = -1
		}
		prevVal = currVal
	}

	return finalRanges
}

func countFresh(freshIds []idRange) int {

	total := 0

	for _, val := range freshIds {
		total += (val.end - val.start) + 1
	}

	return total
}

func determineTotalFresh(dict map[idRange][]int) int {
	total := 0
	for _, val := range dict {
		total += len(val)
	}
	return total
}

func determineFreshIngredients(id int, dict map[idRange][]int) idRange {
	for key := range dict {
		if (id >= key.start) && (id <= key.end) {
			return key
		}
	}
	return idRange{start: -1, end: -1}
}

func determineRanges(ranges []string) map[idRange][]int {
	dict := make(map[idRange][]int)

	for _, r := range ranges {
		beginEnd := strings.Split(r,"-")
		begin, err := strconv.Atoi(beginEnd[0])
		checkErr(err)
		end, err := strconv.Atoi(beginEnd[1])
		checkErr(err)
		dict[idRange{start: begin, end: end}] = []int {}
	}

	return dict
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
