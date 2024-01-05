package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Matches struct {
	value int
	coord [2]int
}

func hasSymbolsAround(numRange []int, rowIndex int, lines []string) ([2]int, bool) {

	start := numRange[0]
	end := numRange[1]
	specialCharRegex := regexp.MustCompile(`\w|\.`)

	ok := false
	answer := [2]int{-1, -1}

loop:
	for i := start; i < end; i++ {
		// check right
		// check top right
		// check bottom right

		// for all
		// check top
		topIndex := rowIndex - 1
		if topIndex > 0 {
			char := lines[topIndex][i]
			boolean := specialCharRegex.FindString(string(char))
			if boolean == "" {
				ok = true
				answer = [2]int{topIndex, i}
				break loop
			}
		}
		// check bottom
		bottomIndex := rowIndex + 1
		if bottomIndex < len(lines) {
			char := lines[bottomIndex][i]
			boolean := specialCharRegex.FindString(string(char))
			if boolean == "" {
				ok = true
				answer = [2]int{bottomIndex, i}
				break loop
			}
		}

		// only for start
		if i == start {
			// check left
			leftIndex := i - 1
			if leftIndex >= 0 {
				char := lines[rowIndex][leftIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					ok = true
					answer = [2]int{rowIndex, leftIndex}
					break loop
				}
			}

			// check top left
			if topIndex > 0 && leftIndex >= 0 {
				char := lines[topIndex][leftIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					ok = true
					answer = [2]int{topIndex, leftIndex}
					break loop
				}
			}
			// check bottom left
			if bottomIndex < len(lines) && leftIndex >= 0 {
				char := lines[bottomIndex][leftIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					ok = true
					answer = [2]int{bottomIndex, leftIndex}
					break loop
				}
			}
		}
		if i == end-1 {
			// check right
			rightIndex := i + 1
			if rightIndex < len(lines[rowIndex]) {
				char := lines[rowIndex][rightIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					ok = true
					answer = [2]int{rowIndex, rightIndex}
					break loop
				}
			}

			// check top right
			if topIndex > 0 && rightIndex < len(lines[rowIndex]) {
				char := lines[topIndex][rightIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					ok = true
					answer = [2]int{topIndex, rightIndex}
					break loop
				}
			}
			// check bottom right
			if bottomIndex < len(lines) && rightIndex < len(lines[rowIndex]) {
				char := lines[bottomIndex][rightIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					ok = true
					answer = [2]int{bottomIndex, rightIndex}
					break loop
				}
			}
		}

	}

	return answer, ok
}

func main() {
	inputPath := flag.String("input", "./input.txt", "Enter input file path")
	flag.Parse()

	file, err := os.Open(*inputPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`[0-9]+`)

	allMatches := [][][]int{}
	lines := []string{}

	total := 0
	ratio := 0
	var allGears []Matches

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		matches := regex.FindAllStringIndex(line, -1)
		allMatches = append(allMatches, matches)
		lines = append(lines, line)
	}

	for rowIndex, match := range allMatches {
		for _, numRange := range match {
			coord, ok := hasSymbolsAround(numRange, rowIndex, lines)
			if ok {
				line := lines[rowIndex]
				numStr := line[numRange[0]:numRange[1]]
				num, _ := strconv.Atoi(numStr)
				total += num

				allGears = append(allGears, Matches{value: num, coord: coord})
			}
		}
	}

	allCoords := map[[2]int][]int{}
	for _, gear := range allGears {
		allCoords[gear.coord] = append(allCoords[gear.coord], gear.value)
	}
	for _, value := range allCoords {
		if len(value) == 2 {
			ratio += value[0] * value[1]
		}
	}

	fmt.Printf("What is the sum of all of the part numbers in the engine schematic?\n%d\n", total)
	fmt.Printf("What is the sum of all of the gear ratios in your engine schematic??\n%d\n", ratio)

}
