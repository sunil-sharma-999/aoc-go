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

func hasSymbolsAround(numRange []int, rowIndex int, lines []string, regex *regexp.Regexp) bool {

	start := numRange[0]
	end := numRange[1]
	specialCharRegex := regexp.MustCompile(`\w|\.`)

	answer := false

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
				answer = true
				break loop
			}
		}
		// check bottom
		bottomIndex := rowIndex + 1
		if bottomIndex < len(lines) {
			char := lines[bottomIndex][i]
			boolean := specialCharRegex.FindString(string(char))
			if boolean == "" {
				answer = true
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
					answer = true
					break loop
				}
			}

			// check top left
			if topIndex > 0 && leftIndex >= 0 {
				char := lines[topIndex][leftIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					answer = true
					break loop
				}
			}
			// check bottom left
			if bottomIndex < len(lines) && leftIndex >= 0 {
				char := lines[bottomIndex][leftIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					answer = true
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
					answer = true
					break loop
				}
			}

			// check top right
			if topIndex > 0 && rightIndex < len(lines[rowIndex]) {
				char := lines[topIndex][rightIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					answer = true
					break loop
				}
			}
			// check bottom right
			if bottomIndex < len(lines) && rightIndex < len(lines[rowIndex]) {
				char := lines[bottomIndex][rightIndex]
				boolean := specialCharRegex.FindString(string(char))
				if boolean == "" {
					answer = true
					break loop
				}
			}
		}

	}

	return answer
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
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		matches := regex.FindAllStringIndex(line, -1)
		allMatches = append(allMatches, matches)
		lines = append(lines, line)
	}
	total := 0
	for rowIndex, match := range allMatches {
		for _, numRange := range match {
			isTrue := hasSymbolsAround(numRange, rowIndex, lines, regex)
			if isTrue {
				line := lines[rowIndex]
				numStr := line[numRange[0]:numRange[1]]
				num, _ := strconv.Atoi(numStr)
				total += num
			}
		}
	}
	fmt.Println(total)

}
