package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MATCH_WORD = "XMAS"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	matchCount := 0

	for i, line := range lines {
		// check if line contains MATCH_WORD
		if strings.Contains(line, MATCH_WORD) {
			matchCount++
		}
		// check if reverse of line contains MATCH_WORD
		if strings.Contains(reverse(line), MATCH_WORD) {
			matchCount++
		}
		// check if vertical line matches MATCH_WORD
		verticalLine := ""
		for _, line := range lines {
			verticalLine += string(line[i])
		}
		if strings.Contains(verticalLine, MATCH_WORD) {
			matchCount++
		}
		// check if reverse of vertical line matches MATCH_WORD
		if strings.Contains(reverse(verticalLine), MATCH_WORD) {
			matchCount++
		}
	}

	fmt.Println("Match count:", matchCount)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
