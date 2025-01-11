package main

import (
	"bufio"
	"fmt"
	"os"
)

const MATCH_WORD = "XMAS"
const MATCH_WORD_LENGTH = len(MATCH_WORD)

func checkWordAtIndex(xAxis, yAxis int, matrix []string) int {
	count := 0
	// center to left
	if xAxis >= MATCH_WORD_LENGTH {
		line := ""
		for i := MATCH_WORD_LENGTH - 1; i >= 0; i-- {
			line += string(matrix[yAxis][xAxis-i])
		}
		if line == MATCH_WORD {
			count++
		}
	}
	// center to right
	if xAxis+MATCH_WORD_LENGTH <= len(matrix[yAxis]) {
		line := matrix[yAxis][xAxis : xAxis+MATCH_WORD_LENGTH]
		if line == MATCH_WORD {
			count++
		}
	}
	// center to top
	if yAxis >= MATCH_WORD_LENGTH {
		line := ""
		for i := MATCH_WORD_LENGTH - 1; i >= 0; i-- {
			line += string(matrix[yAxis-i][xAxis])
		}
		if line == MATCH_WORD {
			count++
		}
	}
	// center to bottom
	if yAxis+MATCH_WORD_LENGTH <= len(matrix) {
		line := ""
		for i := 0; i < MATCH_WORD_LENGTH; i++ {
			line += string(matrix[yAxis+i][xAxis])
		}
		if line == MATCH_WORD {
			count++
		}
	}
	// center to top left
	if xAxis >= MATCH_WORD_LENGTH && yAxis >= MATCH_WORD_LENGTH {
		line := ""
		for i := 0; i < MATCH_WORD_LENGTH; i++ {
			line += string(matrix[yAxis-i][xAxis-i])
		}
		if line == MATCH_WORD {
			count++
		}
	}
	// center to top right
	if xAxis+MATCH_WORD_LENGTH <= len(matrix[yAxis]) && yAxis >= MATCH_WORD_LENGTH {
		line := ""
		for i := 0; i < MATCH_WORD_LENGTH; i++ {
			line += string(matrix[yAxis-i][xAxis+i])
		}
		if line == MATCH_WORD {
			count++
		}
	}
	// center to bottom right
	if xAxis+MATCH_WORD_LENGTH <= len(matrix[yAxis]) && yAxis+MATCH_WORD_LENGTH <= len(matrix) {
		line := ""
		for i := 0; i < MATCH_WORD_LENGTH; i++ {
			line += string(matrix[yAxis+i][xAxis+i])
		}
		if line == MATCH_WORD {
			count++
		}
	}
	// center to bottom left
	if xAxis >= MATCH_WORD_LENGTH && yAxis+MATCH_WORD_LENGTH <= len(matrix) {
		line := ""
		for i := MATCH_WORD_LENGTH - 1; i >= 0; i-- {
			line += string(matrix[yAxis+i][xAxis-i])
		}
		if line == MATCH_WORD {
			count++
		}
	}

	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	count := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			count += checkWordAtIndex(j, i, lines)
		}
	}

	fmt.Printf("\nPart 1: How many times does XMAS appear?\n%d\n", count)

}
