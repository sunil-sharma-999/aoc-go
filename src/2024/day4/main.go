package main

import (
	"bufio"
	"fmt"
	"os"
)

const MATCH_WORD = "XMAS"
const MATCH_WORD_LENGTH = len(MATCH_WORD)

const X_PATTERN_WORD = "MAS"
const X_PATTERN_WORD_LENGTH = len(X_PATTERN_WORD)

func reverseStr(s string) string {
	reversed := ""
	for _, char := range s {
		reversed = string(char) + reversed
	}
	return reversed
}

func checkWordAtIndex(xAxis, yAxis int, matrix []string) int {
	count := 0
	// center to left
	if xAxis >= MATCH_WORD_LENGTH-1 {
		line := ""
		for i := MATCH_WORD_LENGTH - 1; i >= 0; i-- {
			line = string(matrix[yAxis][xAxis-i]) + line
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
	if yAxis >= MATCH_WORD_LENGTH-1 {
		line := ""
		for i := MATCH_WORD_LENGTH - 1; i >= 0; i-- {
			line = string(matrix[yAxis-i][xAxis]) + line
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
	if xAxis >= MATCH_WORD_LENGTH-1 && yAxis >= MATCH_WORD_LENGTH-1 {
		line := ""
		for i := 0; i < MATCH_WORD_LENGTH; i++ {
			line += string(matrix[yAxis-i][xAxis-i])
		}
		if line == MATCH_WORD {
			count++
		}
	}
	// center to top right
	if xAxis+MATCH_WORD_LENGTH <= len(matrix[yAxis]) && yAxis >= MATCH_WORD_LENGTH-1 {
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
	if xAxis >= MATCH_WORD_LENGTH-1 && yAxis+MATCH_WORD_LENGTH <= len(matrix) {
		line := ""
		for i := 0; i < MATCH_WORD_LENGTH; i++ {
			line += string(matrix[yAxis+i][xAxis-i])
		}
		if line == MATCH_WORD {
			count++
		}
	}

	return count
}

func checkXPatternAtIndex(xAxis, yAxis int, matrix []string) int {

	line1 := ""
	line2 := ""
	// top left to bottom right
	repositionedY := yAxis - (X_PATTERN_WORD_LENGTH / 2)
	repositionedX := xAxis - (X_PATTERN_WORD_LENGTH / 2)
	xLimit := xAxis + (X_PATTERN_WORD_LENGTH / 2) + 1
	yLimit := yAxis + (X_PATTERN_WORD_LENGTH / 2) + 1
	for x, y := repositionedX, repositionedY; x < xLimit && y < yLimit; x, y = x+1, y+1 {
		line1 += string(matrix[y][x])
	}

	// top right to bottom left
	repositionedY = yAxis - (X_PATTERN_WORD_LENGTH / 2)
	repositionedX = xAxis + (X_PATTERN_WORD_LENGTH / 2)
	xLimit = xAxis - (X_PATTERN_WORD_LENGTH / 2)
	yLimit = yAxis + (X_PATTERN_WORD_LENGTH / 2)

	for x, y := repositionedX, repositionedY; x >= xLimit && y <= yLimit; x, y = x-1, y+1 {
		line2 += string(matrix[y][x])
	}

	reversedLine1 := reverseStr(line1)
	reversedLine2 := reverseStr(line2)
	if line1 == X_PATTERN_WORD || reversedLine1 == X_PATTERN_WORD {
		if line2 == X_PATTERN_WORD || reversedLine2 == X_PATTERN_WORD {
			return 1
		}
	}

	return 0

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

	count1 := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == MATCH_WORD[0] {
				count1 += checkWordAtIndex(j, i, lines)
			}
		}
	}
	fmt.Printf("\nPart 1: How many times does XMAS appear?\n%d\n", count1)

	count2 := 0

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j] == X_PATTERN_WORD[X_PATTERN_WORD_LENGTH/2] {
				count2 += checkXPatternAtIndex(j, i, lines)
			}
		}
	}

	fmt.Printf("\nPart 2: How many times does an X-MAS appear??\n%d\n", count2)

}
